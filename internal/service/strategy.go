package service

import (
	"context"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type StrategyService interface {
	GetStrategy(ctx context.Context, id int64) (*model.Strategy, error)
	ListStrategy(ctx context.Context, userId int64) (*[]model.Strategy, error)
}

func NewStrategyService(service *Service, strategyRepository repository.StrategyRepository) StrategyService {
	return &strategyService{
		Service:            service,
		strategyRepository: strategyRepository,
	}
}

type strategyService struct {
	*Service
	strategyRepository repository.StrategyRepository
}

func (s *strategyService) ListStrategy(ctx context.Context, userId int64) (*[]model.Strategy, error) {
	return s.strategyRepository.ListStrategyByUserId(ctx, userId)
}

func (s *strategyService) GetStrategy(ctx context.Context, id int64) (*model.Strategy, error) {
	return nil, nil
}
