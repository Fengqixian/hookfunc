package service

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/thoas/go-funk"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
	"sort"
)

type OrderInfoService interface {
	GetOrderInfo(ctx context.Context, id int64) (*model.OrderInfo, error)
	PlaceAnOrder(ctx context.Context, place *v1.PlaceOrderRequest) (*model.OrderInfo, error)
	ListOrder(ctx context.Context, userId int64) []v1.OrderInfoResponse
	GetOrderGoodsDetail(ctx context.Context, orderId int64) *[]v1.OrderGoods
	MapOrderGoodsDetail(ctx context.Context, orders *[]model.OrderInfo) map[int64][]v1.OrderGoods
	CancelOrder(ctx context.Context, orderId int64) error
}

func NewOrderInfoService(service *Service, orderInfoRepository repository.OrderInfoRepository, orderGoodsService OrderGoodsService, goodsService GoodsService, userAddressRepository repository.UserAddressRepository) OrderInfoService {
	return &orderInfoService{
		Service:               service,
		orderInfoRepository:   orderInfoRepository,
		orderGoodsService:     orderGoodsService,
		goodsService:          goodsService,
		userAddressRepository: userAddressRepository,
	}
}

type orderInfoService struct {
	*Service
	orderInfoRepository   repository.OrderInfoRepository
	orderGoodsService     OrderGoodsService
	goodsService          GoodsService
	userAddressRepository repository.UserAddressRepository
}

func (s *orderInfoService) CancelOrder(ctx context.Context, orderId int64) error {
	order := model.OrderInfo{
		ID:         orderId,
		Orderstate: 1,
	}

	err := s.orderInfoRepository.UpdateOrder(ctx, &order)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderInfoService) MapOrderGoodsDetail(ctx context.Context, orders *[]model.OrderInfo) map[int64][]v1.OrderGoods {
	mapping := make(map[int64][]v1.OrderGoods)
	orderGoods := s.orderGoodsService.ListOrderGoods(ctx, funk.Map(*orders, func(item model.OrderInfo) int64 { return item.ID }).([]int64))
	if len(*orderGoods) == 0 {
		return mapping
	}

	goodsMapping := s.goodsService.MapGoodsInfo(ctx)
	for _, item := range *orderGoods {
		var e v1.OrderGoods
		err := mapstructure.Decode(item, &e)
		if err != nil {
			continue
		}

		e.GoodsName = goodsMapping[e.GoodsId].GoodsName
		mapping[item.ID] = append(mapping[item.OrderID], e)
	}
	return mapping
}

func (s *orderInfoService) ListOrder(ctx context.Context, userId int64) []v1.OrderInfoResponse {
	list := make([]v1.OrderInfoResponse, 0)
	orderList, err := s.orderInfoRepository.GetAllOrderByUserId(ctx, userId)
	if err != nil || len(*orderList) == 0 {
		return list
	}

	addressIds := funk.Map(*orderList, func(item model.OrderInfo) int64 { return item.AddressID }).([]int64)
	addressList, err := s.userAddressRepository.GetAddressByIds(ctx, funk.Uniq(addressIds).([]int64))
	if err != nil {
		return list
	}

	addressMapping := funk.ToMap(addressList, "ID").(map[int64]model.UserAddressInfo)
	orderGoodsMap := s.MapOrderGoodsDetail(ctx, orderList)
	for _, info := range *orderList {
		var res v1.OrderInfoResponse
		err := mapstructure.Decode(addressMapping[info.AddressID], &res)
		if err != nil {
			return list
		}

		res.UserId = info.UserID
		res.OrderId = info.ID
		res.PlaceTime = info.CreateTime

		res.Goods = orderGoodsMap[res.OrderId]
		list = append(list, res)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].PlaceTime.Before(list[j].PlaceTime)
	})

	return list
}

func (s *orderInfoService) GetOrderGoodsDetail(ctx context.Context, orderId int64) *[]v1.OrderGoods {
	var list []v1.OrderGoods
	orderGoods, err := s.orderGoodsService.GetOrderGoods(ctx, orderId)
	if err != nil {
		return &list
	}

	goodsMapping := s.goodsService.MapGoodsInfo(ctx)
	for _, goods := range *orderGoods {
		var res v1.OrderGoods
		err := mapstructure.Decode(goods, &res)
		if err != nil {
			return &list
		}

		res.GoodsName = goodsMapping[goods.GoodsID].GoodsName
		res.Total = res.Price * res.Amount
		list = append(list, res)
	}

	return &list
}

func (s *orderInfoService) PlaceAnOrder(ctx context.Context, place *v1.PlaceOrderRequest) (*model.OrderInfo, error) {
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

func (s *orderInfoService) GetOrderInfo(ctx context.Context, id int64) (*model.OrderInfo, error) {
	return s.orderInfoRepository.FirstById(ctx, id)
}
