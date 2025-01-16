package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hookfunc/internal/model"
)

type BarRepository interface {
	ListBar(ctx context.Context) (*[]model.Bar, error)
	ListCoin(ctx context.Context) (*[]model.Coin, error)
}

func NewBarRepository(repository *Repository) BarRepository {
	return &barRepository{
		Repository: repository,
	}
}

type barRepository struct {
	*Repository
}

func (b barRepository) ListCoin(ctx context.Context) (*[]model.Coin, error) {
	var coins []model.Coin
	if err := b.DB(ctx).Where("deleted = 0 order by id asc").Find(&coins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &coins, nil
		}
		return nil, err
	}

	return &coins, nil
}

func (b barRepository) ListBar(ctx context.Context) (*[]model.Bar, error) {
	var bars []model.Bar
	if err := b.DB(ctx).Where("deleted = 0 order by id asc").Find(&bars).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &bars, nil
		}
		return nil, err
	}

	return &bars, nil
}
