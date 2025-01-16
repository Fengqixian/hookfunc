package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hookfunc/internal/model"
)

type BarRepository interface {
	ListBar(ctx context.Context) (*[]model.Bar, error)
}

func NewBarRepository(repository *Repository) BarRepository {
	return &barRepository{
		Repository: repository,
	}
}

type barRepository struct {
	*Repository
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
