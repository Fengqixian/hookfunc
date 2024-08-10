package service

import (
	"context"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/model"
	"hookfunc-media/internal/repository"
)

type GoodsService interface {
	GetGoods(ctx context.Context, params *v1.Goods) (*model.GoodsInfo, error)
	GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error)
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

func (s *goodsService) GetAllGoods(ctx context.Context) (*[]model.GoodsInfo, error) {
	return s.goodsRepository.GetAllGoods(ctx)
}

func (s *goodsService) GetGoods(ctx context.Context, params *v1.Goods) (*model.GoodsInfo, error) {
	return s.goodsRepository.GetGoods(ctx, params.Id)
}
