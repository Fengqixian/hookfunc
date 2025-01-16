package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"hookfunc/internal/model"
)

type StrategyRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Strategy, error)
	ListStrategyByUserId(ctx context.Context, userId int64) (*[]model.Strategy, error)
}

func NewStrategyRepository(repository *Repository) StrategyRepository {
	return &strategyRepository{
		Repository: repository,
	}
}

type strategyRepository struct {
	*Repository
}

func (r *strategyRepository) ListStrategyByUserId(ctx context.Context, userId int64) (*[]model.Strategy, error) {
	var strategies []model.Strategy
	if err := r.DB(ctx).Where("deleted = 0 and user_id = ? order by id desc", userId).Find(&strategies).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &strategies, nil
		}
		return nil, err
	}

	return &strategies, nil
}

func (r *strategyRepository) FirstById(ctx context.Context, id int64) (*model.Strategy, error) {
	var strategy model.Strategy
	// TODO: query db
	return &strategy, nil
}
