package service

import (
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type OrderInfoService interface {
	GetOrderInfo(id int64) (*model.OrderInfo, error)
}

func NewOrderInfoService(service *Service, orderInfoRepository repository.OrderInfoRepository) OrderInfoService {
	return &orderInfoService{
		Service:             service,
		orderInfoRepository: orderInfoRepository,
	}
}

type orderInfoService struct {
	*Service
	orderInfoRepository repository.OrderInfoRepository
}

func (s *orderInfoService) GetOrderInfo(id int64) (*model.OrderInfo, error) {
	return s.orderInfoRepository.FirstById(id)
}
