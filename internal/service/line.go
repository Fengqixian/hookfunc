package service

import (
	"context"
	"go.uber.org/zap"
	"hookfunc/internal/model"
	"hookfunc/internal/okx"
	"time"
)

type LineService interface {
	AllCoinChainKlineDataSync(ctx context.Context)
	UpdateKLine(instId string, bar string)
	GetKLine(instId string, bar string) []model.LineItem
}

func NewLineService(bar BarService, okxConfig *okx.Config, service *Service) LineService {
	api := okx.NewApi(okxConfig)
	return &lineService{
		Service:  service,
		Config:   okxConfig,
		api:      api,
		bar:      bar,
		klineMap: make(map[string][]model.LineItem),
	}
}

type lineService struct {
	*Service
	*okx.Config
	// key = InstId + Bar
	klineMap map[string][]model.LineItem
	api      *okx.Api
	bar      BarService
}

func (l *lineService) AllCoinChainKlineDataSync(ctx context.Context) {
	coins, err := l.bar.ListCoin(ctx)
	if err != nil {
		return
	}

	bars, err := l.bar.ListBar(ctx)
	if err != nil {
		return
	}

	for _, coin := range *coins {
		for _, bar := range *bars {
			go func() {
				if okx.CheckMinuteModulo(bar.Interval) {
					l.UpdateKLine(coin.Name, bar.Value)
				}
			}()
		}
	}
}

func (l *lineService) UpdateKLine(instId string, bar string) {
	start := time.Now()
	data, err := l.api.GetLine(instId, bar, l.Config.Limit)
	if err != nil {
		l.logger.Logger.Error("【获取K线数据失败】", zap.String("instId", instId), zap.String("bar", bar), zap.Error(err))
		return
	}

	result, err := data.GetLineItem()
	if err != nil {
		l.logger.Logger.Error("【转换K线数据失败】", zap.String("instId", instId), zap.String("bar", bar), zap.Error(err))
		return
	}

	l.klineMap[instId+"_"+bar] = result

	elapsed := time.Since(start)
	l.logger.Logger.Info("【OKX数据同步成功】", zap.String("instId", instId), zap.String("bar", bar), zap.Any("elapsed", elapsed))
}

func (l *lineService) GetKLine(instId string, bar string) []model.LineItem {
	key := instId + "_" + bar
	line := l.klineMap[key]
	if line == nil {
		l.UpdateKLine(instId, bar)
	}

	return l.klineMap[instId+"_"+bar]
}
