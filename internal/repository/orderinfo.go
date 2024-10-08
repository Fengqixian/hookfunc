package repository

import (
	"context"
	"hookfunc/internal/model"
)

type OrderInfoRepository interface {
	FirstById(ctx context.Context, id int64) (*model.OrderInfo, error)
	CreateOrder(ctx context.Context, order *model.OrderInfo) (*model.OrderInfo, error)
	GetAllOrderByUserId(ctx context.Context, userId int64) (*[]model.OrderInfo, error)
	UpdateOrder(ctx context.Context, order *model.OrderInfo) error
}

func NewOrderInfoRepository(repository *Repository) OrderInfoRepository {
	return &orderInfoRepository{
		Repository: repository,
	}
}

type orderInfoRepository struct {
	*Repository
}

func (r *orderInfoRepository) GetAllOrderByUserId(ctx context.Context, userId int64) (*[]model.OrderInfo, error) {
	var list []model.OrderInfo
	if err := r.DB(ctx).Where("deleted = 0 and user_id = ?", userId).Order("id desc").Find(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *orderInfoRepository) UpdateOrder(ctx context.Context, order *model.OrderInfo) error {
	if err := r.DB(ctx).Updates(order).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderInfoRepository) CreateOrder(ctx context.Context, order *model.OrderInfo) (*model.OrderInfo, error) {
	if err := r.DB(ctx).Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *orderInfoRepository) FirstById(ctx context.Context, id int64) (*model.OrderInfo, error) {
	var orderInfo model.OrderInfo
	if err := r.DB(ctx).Where("id = ?", id).First(&orderInfo).Error; err != nil {
		return nil, err
	}

	return &orderInfo, nil
}
