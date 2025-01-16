package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
)

type StrategyService interface {
	ListStrategyIndex(ctx context.Context, req v1.StrategyRequest) (*[]model.StrategyIndex, error)
	ListStrategy(ctx context.Context, userId int64) (*[]model.Strategy, error)
	CreateStrategy(ctx context.Context, req v1.CreateStrategyRequest) error
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

func (s *strategyService) CreateStrategy(ctx context.Context, req v1.CreateStrategyRequest) error {
	strategy := model.Strategy{
		StrategyName:      req.Name,
		SubscriptionState: req.SubscriptionState,
		UserID:            req.UserId,
	}

	err := s.strategyRepository.CreateStrategy(ctx, &strategy)
	if err != nil {
		s.logger.Logger.Error("【创建策略失败】", zap.Error(err))
		return errors.New("创建策略失败")
	}

	var strategyIndexList []model.StrategyIndex
	for _, index := range req.IndexList {
		strategyIndex := model.StrategyIndex{
			StrategyID:   strategy.ID,
			UserID:       req.UserId,
			IndexID:      index.IndexId,
			Bar:          index.Bar,
			IndexConfig:  index.IndexConfig,
			WarningIndex: index.WarningIndex,
		}
		strategyIndexList = append(strategyIndexList, strategyIndex)
	}

	err = s.strategyRepository.CreateStrategyIndex(ctx, &strategyIndexList)
	if err != nil {
		return err
	}

	return nil
}

func (s *strategyService) ListStrategy(ctx context.Context, userId int64) (*[]model.Strategy, error) {
	return s.strategyRepository.ListStrategyByUserId(ctx, userId)
}

func (s *strategyService) ListStrategyIndex(ctx context.Context, req v1.StrategyRequest) (*[]model.StrategyIndex, error) {
	return s.strategyRepository.ListStrategyIndexByStrategyId(ctx, req.StrategyId, req.UserId)
}
