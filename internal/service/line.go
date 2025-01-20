package service

import (
	"go.uber.org/zap"
	"hookfunc/internal/model"
	"hookfunc/internal/okx"
)

type LineService interface {
	UpdateKLine(instId string, bar string)
	GetKLine(instId string, bar string) []model.LineItem
}

func NewLineService(okxConfig *okx.Config, service *Service) LineService {
	api := okx.NewApi(okxConfig)

	return &lineService{
		Service: service,
		Config:  okxConfig,
		api:     api,
	}
}

type lineService struct {
	*Service
	*okx.Config
	// key = InstId + Bar
	klineMap map[string][]model.LineItem
	api      *okx.Api
}

func (l *lineService) UpdateKLine(instId string, bar string) {
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
	l.logger.Logger.Info("【更新K线数据成功】", zap.String("instId", instId), zap.String("bar", bar))
}

func (l *lineService) GetKLine(instId string, bar string) []model.LineItem {
	return l.klineMap[instId+"_"+bar]
}
