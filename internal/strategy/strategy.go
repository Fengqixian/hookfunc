package strategy

import (
	"errors"
	"hookfunc/internal/model"
	"hookfunc/internal/okx"
)

type Strategy interface {
	Execute(line []model.LineItem, params []int64, warningConfig string) (any, error)
}

type WarningStrategy struct {
	Strategy map[string]Strategy
}

type WarningConfig struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

func NewWarningStrategy() *WarningStrategy {
	var strategy = make(map[string]Strategy)
	strategy["MACD"] = &MACDStrategy{}
	strategy["RSI"] = &RSIStrategy{}

	return &WarningStrategy{Strategy: strategy}
}

/**
 * MACD Strategy
 * name: MACD
 * default_config: [{"text":"短周期", "value": 12}, {"text":"长周期", "value": 26}, {"text":"移动平均周期", "value": 9}]
 * warning_config_array: [{"index": 0, "name":"金叉"},{"index": 1, "name":"死叉"},{"index": 2, "name":"多头"},{"index": 3, "name":"空头"}]
 */

type MACDStrategy struct{}

func (s *MACDStrategy) Execute(line []model.LineItem, params []int64, warningConfigString string) (any, error) {
	if len(params) < 3 {
		return nil, errors.New("MACD 策略执行失败: invalid params")
	}

	line = okx.CalculateMACD(line, params[0], params[1], params[2])
	var warningConfig WarningConfig
	err := okx.ConvertStringToObject(warningConfigString, &warningConfig)
	if err != nil {
		return nil, err
	}

	switch warningConfig.Index {
	case 0: // 金叉
		return okx.MACDCross(line, true), nil
	case 1: // 死叉
		return okx.MACDCross(line, false), nil
	case 2: // 多头
		return okx.MACDLongOrSort(line, true), nil
	case 3: // 空头
		return okx.MACDLongOrSort(line, false), nil
	default:
		return nil, errors.New("MACD: 未知的预警指标")
	}
}

/**
 * RSI Strategy
 * name: RSI
 * default_config: [{"text":"RSI1", "value": 14}]
 * warning_config_array: [{"index": 0,"name":"上穿", "value": 30},{"index": 1,"name":"下穿", "value": 70},{"index": 2,"name":"大于", "value": 30},{"index": 3,"name":"小于", "value": 70}]
 */

type RSIStrategy struct{}

func (s *RSIStrategy) Execute(line []model.LineItem, params []int64, warningConfigString string) (any, error) {
	if len(params) < 1 {
		return nil, errors.New("MACD 策略执行失败: invalid params")
	}

	line = okx.CalculateRSI(line, params[0])
	var warningConfig WarningConfig
	err := okx.ConvertStringToObject(warningConfigString, &warningConfig)
	if err != nil {
		return nil, err
	}

	switch warningConfig.Index {
	case 0:
		return okx.RSICrossUpOrDown(line, true, float64(warningConfig.Value)), nil
	case 1:
		return okx.RSICrossUpOrDown(line, false, float64(warningConfig.Value)), nil
	case 2:
		return okx.RSICrossQK(line, true, float64(warningConfig.Value)), nil
	case 3:
		return okx.RSICrossQK(line, false, float64(warningConfig.Value)), nil
	default:
		return nil, errors.New("RSI: 未知的预警指标")
	}
}
