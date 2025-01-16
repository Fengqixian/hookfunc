package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/model"
	"hookfunc/internal/okx"
	"hookfunc/internal/repository"
)

type IndexService interface {
	ListIndex(ctx context.Context) (*[]model.Index, error)
	IndexHitTarget(ctx context.Context, req v1.IndexRequest) (any, error)
}

func NewIndexService(service *Service, indexRepository repository.IndexRepository) IndexService {
	kline := okx.NewLine("https://www.okx.com", 10)
	return &indexService{
		Service:         service,
		indexRepository: indexRepository,
		kline:           kline,
		s:               okx.NewWarningStrategy(kline),
	}
}

type indexService struct {
	*Service
	indexRepository repository.IndexRepository
	kline           *okx.KLine
	s               *okx.WarningStrategy
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

	line, err := i.kline.GetLine(req.InstId, req.Bar, "300")
	if err != nil {
		i.logger.Error("【指标回测失败】未获取到K线数据", zap.Error(err))
		return nil, err
	}

	if line == nil || len(line) == 0 {
		i.logger.Error("【指标回测失败】kline not found", zap.Any("req", req))
		return nil, errors.New("服务繁忙，请稍后再试")
	}

	array, err := GetIndexConfigAsInt64Array(req.IndexConfig)
	if err != nil {
		i.logger.Error("【指标回测失败】", zap.Error(err))
		return nil, fmt.Errorf("指标配置错误：%s", req.IndexConfig)
	}

	result, err := i.s.Strategy[index.Name].Execute(line, array, req.WarningIndex)
	if err != nil {
		i.logger.Error("【指标回测失败】", zap.Error(err))
		return nil, err
	}

	fmt.Println(result)
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
