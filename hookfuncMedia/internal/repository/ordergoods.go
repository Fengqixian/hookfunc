package repository

import (
	"context"
	"hookfunc-media/internal/model"
)

type OrderGoodsRepository interface {
	FirstById(id int64) (*model.OrderGoods, error)
	CreateOrderGoods(ctx context.Context, orderGoods *[]model.OrderGoods) error
}

func NewOrderGoodsRepository(repository *Repository) OrderGoodsRepository {
	return &orderGoodsRepository{
		Repository: repository,
	}
}

type orderGoodsRepository struct {
	*Repository
}

func (r *orderGoodsRepository) CreateOrderGoods(ctx context.Context, orderGoods *[]model.OrderGoods) error {
	if err := r.DB(ctx).Create(orderGoods).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderGoodsRepository) FirstById(id int64) (*model.OrderGoods, error) {
	var orderGoods model.OrderGoods
	// TODO: query db
	return &orderGoods, nil
}
