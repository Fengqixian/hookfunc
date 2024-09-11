package service

import (
	"context"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type OrderGoodsService interface {
	GetOrderGoods(ctx context.Context, orderId int64) (*[]model.OrderGoods, error)
	ListOrderGoods(ctx context.Context, orderIds []int64) *[]model.OrderGoods
	CreateOrderGoods(ctx context.Context, place *v1.PlaceOrderRequest, orderId int64) error
}

func NewOrderGoodsService(service *Service, orderGoodsRepository repository.OrderGoodsRepository, goodsService GoodsService) OrderGoodsService {
	return &orderGoodsService{
		Service:              service,
		orderGoodsRepository: orderGoodsRepository,
		goodsService:         goodsService,
	}
}

type orderGoodsService struct {
	*Service
	orderGoodsRepository repository.OrderGoodsRepository
	goodsService         GoodsService
}

func (s *orderGoodsService) ListOrderGoods(ctx context.Context, orderIds []int64) *[]model.OrderGoods {
	list, err := s.orderGoodsRepository.ListByOrderIds(ctx, orderIds)
	if err != nil {
		res := make([]model.OrderGoods, 0)
		return &res
	}

	return list
}

func (s *orderGoodsService) CreateOrderGoods(ctx context.Context, place *v1.PlaceOrderRequest, orderId int64) error {
	if len(*place.OrderGoods) == 0 {
		return v1.ErrRequestParamsFail
	}

	goodsInfoList, err := s.goodsService.GetGoodsByIds(ctx, *place.OrderGoods)
	if err != nil || len(*goodsInfoList) == 0 {
		return v1.ErrPlaceAnOrderFail
	}

	var goodsPriceMap = make(map[int64]int32)
	for _, goods := range *goodsInfoList {
		goodsPriceMap[goods.ID] = goods.Price
	}

	orderGoodsList := *place.OrderGoods
	for i, info := range orderGoodsList {
		orderGoodsList[i].OrderID = orderId
		orderGoodsList[i].Price = goodsPriceMap[info.GoodsID]
	}

	err = s.orderGoodsRepository.CreateOrderGoods(ctx, &orderGoodsList)
	if err != nil {
		return err
	}

	//TODO 打印订单
	return nil
}

func (s *orderGoodsService) GetOrderGoods(ctx context.Context, orderId int64) (*[]model.OrderGoods, error) {
	return s.orderGoodsRepository.ListByOrderId(ctx, orderId)
}
