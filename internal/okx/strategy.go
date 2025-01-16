package okx

import (
	"errors"
	"hookfunc/internal/model"
)

type Strategy interface {
	Execute(line []model.LineItem, params []int64, warningIndex int32) (any, error)
}

type WarningStrategy struct {
	Strategy map[string]Strategy
}

func NewWarningStrategy(kline *KLine) *WarningStrategy {
	var strategy = make(map[string]Strategy)
	strategy["MACD"] = &MACDStrategy{Kline: kline}

	return &WarningStrategy{Strategy: strategy}
}

/**
 * MACD Strategy
 * name: MACD
 * default_config: [12, 26, 9]
 * warning_config: [{"name":"金叉"},{"name":"死叉"},{"name":"多头"},{"name":"空头"}]
 */

type MACDStrategy struct {
	Kline *KLine
}

func (s *MACDStrategy) Execute(line []model.LineItem, params []int64, warningIndex int32) (any, error) {
	if len(params) != 3 {
		return nil, errors.New("MACD 策略执行失败: invalid params")
	}
	line = s.Kline.CalculateMACD(line, params[0], params[1], params[2])
	switch warningIndex {
	case 0:
		return "金叉", nil
	case 1:
		return "死叉", nil
	case 2:
		return "多头", nil
	case 3:
		return "空头", nil
	default:
		return nil, errors.New("未知的预警指标")
	}
}
