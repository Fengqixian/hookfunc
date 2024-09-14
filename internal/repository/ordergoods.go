package repository

import (
	"context"
	"hookfunc/internal/model"
)

type OrderGoodsRepository interface {
	ListByOrderIds(ctx context.Context, orderIds []int64) (*[]model.OrderGoods, error)
	ListByOrderId(ctx context.Context, orderId int64) (*[]model.OrderGoods, error)
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

func (r *orderGoodsRepository) ListByOrderIds(ctx context.Context, orderIds []int64) (*[]model.OrderGoods, error) {
	var list []model.OrderGoods
	if err := r.DB(ctx).Where("order_id in ?", orderIds).Find(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *orderGoodsRepository) ListByOrderId(ctx context.Context, orderId int64) (*[]model.OrderGoods, error) {
	var list []model.OrderGoods
	if err := r.DB(ctx).Where("order_id = ?", orderId).Find(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *orderGoodsRepository) CreateOrderGoods(ctx context.Context, orderGoods *[]model.OrderGoods) error {
	if err := r.DB(ctx).Create(orderGoods).Error; err != nil {
		return err
	}

	return nil
}
