package repository

import (
	"hookfunc-media/internal/model"
)

type OrderGoodsRepository interface {
	FirstById(id int64) (*model.OrderGoods, error)
}

func NewOrderGoodsRepository(repository *Repository) OrderGoodsRepository {
	return &orderGoodsRepository{
		Repository: repository,
	}
}

type orderGoodsRepository struct {
	*Repository
}

func (r *orderGoodsRepository) FirstById(id int64) (*model.OrderGoods, error) {
	var orderGoods model.OrderGoods
	// TODO: query db
	return &orderGoods, nil
}
