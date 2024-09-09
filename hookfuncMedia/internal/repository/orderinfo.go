package repository

import (
	"hookfunc-media/internal/model"
)

type OrderInfoRepository interface {
	FirstById(id int64) (*model.OrderInfo, error)
}

func NewOrderInfoRepository(repository *Repository) OrderInfoRepository {
	return &orderInfoRepository{
		Repository: repository,
	}
}

type orderInfoRepository struct {
	*Repository
}

func (r *orderInfoRepository) FirstById(id int64) (*model.OrderInfo, error) {
	var orderInfo model.OrderInfo
	// TODO: query db
	return &orderInfo, nil
}
