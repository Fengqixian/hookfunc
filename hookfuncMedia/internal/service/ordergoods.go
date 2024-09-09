package service

import (
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type OrderGoodsService interface {
	GetOrderGoods(id int64) (*model.OrderGoods, error)
}

func NewOrderGoodsService(service *Service, orderGoodsRepository repository.OrderGoodsRepository) OrderGoodsService {
	return &orderGoodsService{
		Service:              service,
		orderGoodsRepository: orderGoodsRepository,
	}
}

type orderGoodsService struct {
	*Service
	orderGoodsRepository repository.OrderGoodsRepository
}

func (s *orderGoodsService) GetOrderGoods(id int64) (*model.OrderGoods, error) {
	return s.orderGoodsRepository.FirstById(id)
}
