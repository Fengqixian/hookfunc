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
	case 0: // 金叉
		return MACDCross(line, true), nil
	case 1: // 死叉
		return MACDCross(line, false), nil
	case 2: // 多头
		return MACDLongOrSort(line, true), nil
	case 3: // 空头
		return MACDLongOrSort(line, false), nil
	default:
		return nil, errors.New("未知的预警指标")
	}
}
