package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
)

type StrategyRepository interface {
	FirstById(ctx context.Context, id int64) (*model.Strategy, error)
	ListStrategyByUserId(ctx context.Context, userId int64) (*[]model.Strategy, error)
	CreateStrategy(ctx context.Context, strategy *model.Strategy) error
	CreateStrategyIndex(ctx context.Context, strategyIndex *[]model.StrategyIndex) error
	ListStrategyIndexByStrategyId(ctx context.Context, strategyId int64, userId int64) (*[]model.StrategyIndex, error)
	DeleteStrategyIndex(ctx context.Context, req v1.StrategyIndexRequest) error
	DeleteStrategy(ctx context.Context, req v1.StrategyRequest) error
	UpdateStrategySubscriptionState(ctx context.Context, req v1.StrategyRequest) error
}

func NewStrategyRepository(repository *Repository) StrategyRepository {
	return &strategyRepository{
		Repository: repository,
	}
}

type strategyRepository struct {
	*Repository
}

func (r *strategyRepository) UpdateStrategySubscriptionState(ctx context.Context, req v1.StrategyRequest) error {
	sql := "UPDATE strategy SET subscription_state = IF(subscription_state = 0, 1, 0) WHERE id = ? AND user_id = ? AND deleted = 0"
	if err := r.DB(ctx).Exec(sql, req.StrategyId, req.UserId).Error; err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) DeleteStrategy(ctx context.Context, req v1.StrategyRequest) error {
	if err := r.DB(ctx).Model(&model.Strategy{}).Where("id = ? and user_id = ? and deleted = 0", req.StrategyId, req.UserId).Update("deleted", 1).Error; err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) DeleteStrategyIndex(ctx context.Context, req v1.StrategyIndexRequest) error {
	if err := r.DB(ctx).Model(&model.StrategyIndex{}).Where("id = ? and user_id = ? and deleted = 0", req.StrategyIndexId, req.UserId).Update("deleted", 1).Error; err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) ListStrategyIndexByStrategyId(ctx context.Context, strategyId int64, userId int64) (*[]model.StrategyIndex, error) {
	var strategies []model.StrategyIndex
	if err := r.DB(ctx).Where("deleted = 0 and strategy_id = ? and user_id = ? order by id asc", strategyId, userId).Find(&strategies).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &strategies, nil
		}
		return nil, err
	}

	return &strategies, nil
}

func (r *strategyRepository) CreateStrategyIndex(ctx context.Context, strategyIndex *[]model.StrategyIndex) error {
	if err := r.DB(ctx).Create(strategyIndex).Error; err != nil {
		return err
	}

	return nil
}

func (r *strategyRepository) CreateStrategy(ctx context.Context, strategy *model.Strategy) error {
	if err := r.DB(ctx).Create(strategy).Error; err != nil {
		return err
	}

	return nil
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
