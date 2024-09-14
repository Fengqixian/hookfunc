package service

import (
	"context"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type GoodsService interface {
	GetGoods(ctx context.Context, params *v1.Goods) (*model.GoodsInfo, error)
	GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error)
	GetGoodsByIds(ctx context.Context, orderGoods []model.OrderGoods) (*[]model.GoodsInfo, error)
	MapGoodsInfo(ctx context.Context) map[int64]model.GoodsInfo
}

func NewGoodsService(
	service *Service,
	goodsRepository repository.GoodsRepository,
) GoodsService {
	return &goodsService{
		Service:         service,
		goodsRepository: goodsRepository,
	}
}

type goodsService struct {
	*Service
	goodsRepository repository.GoodsRepository
}

func (s *goodsService) MapGoodsInfo(ctx context.Context) map[int64]model.GoodsInfo {
	mapping := make(map[int64]model.GoodsInfo)
	goods, err := s.GetAllGoods(ctx)
	if err != nil {
		return mapping
	}

	for _, info := range *goods {
		mapping[info.ID] = info
	}

	return mapping
}

func (s *goodsService) GetGoodsByIds(ctx context.Context, orderGoods []model.OrderGoods) (*[]model.GoodsInfo, error) {
	if len(orderGoods) == 0 {
		return nil, v1.ErrRequestParamsFail
	}

	var ids []int64
	for _, good := range orderGoods {
		ids = append(ids, good.GoodsID)
	}

	return s.goodsRepository.ListGoodsById(ctx, ids)
}

func (s *goodsService) GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error) {
	return s.goodsRepository.GetAllGoods(ctx)
}

func (s *goodsService) GetGoods(ctx context.Context, params *v1.Goods) (*model.GoodsInfo, error) {
	return s.goodsRepository.GetGoods(ctx, params.Id)
}
