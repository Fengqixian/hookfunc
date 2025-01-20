package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/repository"
	"hookfunc/internal/strategy"
)

type IndexService interface {
	ListIndex(ctx context.Context) (*[]model.Index, error)
	IndexHitTarget(ctx context.Context, req v1.IndexRequest) (any, error)
}

func NewIndexService(service *Service, lineService LineService, indexRepository repository.IndexRepository) IndexService {
	return &indexService{
		Service:         service,
		indexRepository: indexRepository,
		s:               strategy.NewWarningStrategy(),
		lineService:     lineService,
	}
}

type indexService struct {
	*Service
	indexRepository repository.IndexRepository
	s               *strategy.WarningStrategy
	lineService     LineService
}

func (i indexService) IndexHitTarget(ctx context.Context, req v1.IndexRequest) (any, error) {
	index, err := i.indexRepository.GetIndex(ctx, req.IndexId)
	if err != nil {
		return nil, err
	}

	if index == nil {
		i.logger.Logger.Error("【指标回测失败】index not found", zap.Any("req", req))
		return nil, errors.New("index not found")
	}

	array, err := GetIndexConfigAsInt64Array(req.IndexConfig)
	if err != nil {
		i.logger.Error("【指标回测失败】", zap.Error(err))
		return nil, fmt.Errorf("指标配置错误：%s", req.IndexConfig)
	}

	data := i.lineService.GetKLine(req.InstId, req.Bar)
	if data == nil {
		i.logger.Error("【指标回测失败】获取K线数据失败", zap.Any("req", req))
		return nil, fmt.Errorf("服务繁忙，请稍后再试")
	}

	result, err := i.s.Strategy[index.Name].Execute(data, array, req.WarningConfig)
	if err != nil {
		i.logger.Error("【指标回测失败】", zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (i indexService) ListIndex(ctx context.Context) (*[]model.Index, error) {
	return i.indexRepository.ListIndex(ctx)
}

func GetIndexConfigAsInt64Array(str string) ([]int64, error) {
	var indexConfigArray []int64
	err := json.Unmarshal([]byte(str), &indexConfigArray)
	if err != nil {
		return nil, err
	}
	return indexConfigArray, nil
}
