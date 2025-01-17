package model

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type LineResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type LineItem struct {
	Timestamp       int64   // 时间戳
	TimeString      string  // 时间字符串
	ClosePrice      float64 // 收盘价格
	OpenPrice       float64 // 开盘价格
	HighPrice       float64 // 最高价格
	LowPrice        float64 // 最低价格
	VolumeContracts float64 // 交易量，以张为单位
	VolumeCoins     float64 // 交易量，以币为单位
	VolumeQuoted    float64 // 交易量，以计价货币为单位
	KLineStatus     int     // K线状态
	Index           LineItemIndex
}

type LineItemIndex struct {
	SortEma float64 // EMA 均线
	LongEma float64 // 长期 EMA 均线
	Dif     float64 // 差值
	Dea     float64 // 讯号线
	Rsi     float64 // RSI 指标
	Atr     float64 // ATR 指标
	KDJ     KDJ
}

type KDJ struct {
	K float64 // K 值
	D float64 // D 值
	J float64 // J 值
}

func (l *LineItem) String() string {
	// 将对象转换为JSON字符串
	jsonData, err := json.Marshal(l)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func (l *LineResult) GetLineItem() ([]LineItem, error) {
	lineItems := make([]LineItem, len(l.Data))
	for i := 0; i < len(l.Data); i++ {
		data := l.Data[len(l.Data)-(i+1)]
		item, err := convertTradingData(data)
		if err != nil {
			return nil, errors.New("k线数据转换失败")
		}

		lineItems[i] = *item
	}

	return lineItems, nil
}

func convertTradingData(dataArray []string) (*LineItem, error) {
	tradingData := LineItem{}

	// 转换各个字段
	var err error
	if tradingData.Timestamp, err = strconv.ParseInt(dataArray[0], 10, 64); err != nil {
		return nil, err
	}

	tradingData.TimeString = DateTimeConvert(tradingData.Timestamp)
	if tradingData.OpenPrice, err = strconv.ParseFloat(dataArray[1], 64); err != nil {
		return nil, err
	}

	if tradingData.HighPrice, err = strconv.ParseFloat(dataArray[2], 64); err != nil {
		return nil, err
	}

	if tradingData.LowPrice, err = strconv.ParseFloat(dataArray[3], 64); err != nil {
		return nil, err
	}

	if tradingData.VolumeContracts, err = strconv.ParseFloat(dataArray[5], 64); err != nil {
		return nil, err
	}

	if tradingData.VolumeCoins, err = strconv.ParseFloat(dataArray[6], 64); err != nil {
		return nil, err
	}

	if tradingData.VolumeQuoted, err = strconv.ParseFloat(dataArray[7], 64); err != nil {
		return nil, err
	}

	if tradingData.KLineStatus, err = strconv.Atoi(dataArray[8]); err != nil {
		return nil, err
	}

	if tradingData.ClosePrice, err = strconv.ParseFloat(dataArray[4], 64); err != nil {
		return nil, err
	}

	return &tradingData, nil
}

func DateTimeConvert(dateTime int64) string {
	shanghaiLocation, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Unix(dateTime/1000, 0).In(shanghaiLocation)
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
