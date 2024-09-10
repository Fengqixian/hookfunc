package service

import (
	"context"
	"github.com/mitchellh/mapstructure"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type OrderInfoService interface {
	GetOrderInfo(id int64) (*model.OrderInfo, error)
	PlaceAnOrder(ctx context.Context, place v1.PlaceOrderRequest) (*model.OrderInfo, error)
}

func NewOrderInfoService(service *Service, orderInfoRepository repository.OrderInfoRepository, orderGoodsService OrderGoodsService) OrderInfoService {
	return &orderInfoService{
		Service:             service,
		orderInfoRepository: orderInfoRepository,
		orderGoodsService:   orderGoodsService,
	}
}

type orderInfoService struct {
	*Service
	orderInfoRepository repository.OrderInfoRepository
	orderGoodsService   OrderGoodsService
	goodsService        GoodsService
}

func (s *orderInfoService) PlaceAnOrder(ctx context.Context, place v1.PlaceOrderRequest) (*model.OrderInfo, error) {
	var order *model.OrderInfo
	err := mapstructure.Decode(place, &order)
	if err != nil {
		return nil, v1.ErrPlaceAnOrderFail
	}

	order, err = s.orderInfoRepository.CreateOrder(ctx, order)
	if err != nil {
		return nil, v1.ErrPlaceAnOrderFail
	}

	err = s.orderGoodsService.CreateOrderGoods(ctx, place, order.ID)
	if err != nil {
		return nil, v1.ErrPlaceAnOrderFail
	}

	return order, nil
}

func (s *orderInfoService) GetOrderInfo(id int64) (*model.OrderInfo, error) {
	return s.orderInfoRepository.FirstById(id)
}
