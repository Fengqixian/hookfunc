package okx

import (
	"encoding/json"
	"hookfunc/internal/model"
	"math"
)

// CalculateATR calculates the ATR (Average True Range) over a given period.
// data: slice of TradeData containing historical trading data
// period: number of periods for the ATR calculation (e.g., 14)
func CalculateATR(data []model.LineItem, period int) []model.LineItem {
	trueRanges := make([]float64, len(data)-1)
	for i := 1; i < len(data); i++ {
		currentHigh := data[i].HighPrice
		currentLow := data[i].LowPrice
		previousClose := data[i-1].ClosePrice

		tr := math.Max(currentHigh-currentLow, math.Max(math.Abs(currentHigh-previousClose), math.Abs(currentLow-previousClose)))
		trueRanges[i-1] = tr
	}

	var initialATR float64
	for i := 0; i < period; i++ {
		initialATR += trueRanges[i]
	}
	initialATR /= float64(period)
	data[period-1].Index.Atr = initialATR

	previousATR := initialATR
	for i := period; i < len(trueRanges); i++ {
		currentATR := (previousATR*float64(period-1) + trueRanges[i]) / float64(period)
		data[i].Index.Atr = currentATR
		previousATR = currentATR
	}

	return data
}

func CalculateMACD(data []model.LineItem, sortEma, longEma, signalPeriod int64) []model.LineItem {
	if len(data) == 0 {
		return data
	}

	data[0].Index.SortEma = data[0].ClosePrice
	data[0].Index.LongEma = data[0].ClosePrice
	for i := 1; i < len(data); i++ {
		data[i].Index.SortEma = CalculateSingletEMA(data[i].ClosePrice, data[i-1].Index.SortEma, sortEma)
		data[i].Index.LongEma = CalculateSingletEMA(data[i].ClosePrice, data[i-1].Index.LongEma, longEma)
		data[i].Index.Dif = data[i].Index.SortEma - data[i].Index.LongEma
		data[i].Index.Dea = CalculateSingletEMA(data[i].Index.Dif, data[i-1].Index.Dea, signalPeriod)
	}

	return data
}

func CalculateSingletEMA(prices, lastEma float64, period int64) float64 {
	k := 2.0 / float64(period+1)
	return prices*k + lastEma*(1-k)
}

// CalculateKDJ 计算 KDJ 指标
func CalculateKDJ(data []model.LineItem, N int) []model.LineItem {
	var kdj model.KDJ
	var k, d float64 = 50, 50 // 初始 K、D 值设置为 50

	// 从第 N 天开始计算 KDJ
	for i := N - 1; i < len(data); i++ {
		// 计算过去 N 天的最高价和最低价
		var highestHigh, lowestLow float64
		highestHigh = math.Inf(-1) // 设置为负无穷
		lowestLow = math.Inf(1)    // 设置为正无穷

		for j := i - N + 1; j <= i; j++ {
			if data[j].HighPrice > highestHigh {
				highestHigh = data[j].HighPrice
			}
			if data[j].LowPrice < lowestLow {
				lowestLow = data[j].LowPrice
			}
		}

		// 计算 RSV
		rsv := (data[i].ClosePrice - lowestLow) / (highestHigh - lowestLow) * 100

		// 计算 K 和 D
		k = (2.0/3.0)*k + (1.0/3.0)*rsv
		d = (2.0/3.0)*d + (1.0/3.0)*k

		// 计算 J
		j := 3.0*k - 2.0*d
		kdj.K = k
		kdj.D = d
		kdj.J = j
		data[i].Index.KDJ = kdj
	}

	return data
}

func CalculateRSI(prices []model.LineItem, period int64) []model.LineItem {
	var gains, losses []float64

	// 计算每日的涨跌幅
	for i := 1; i < len(prices); i++ {
		change := prices[i].ClosePrice - prices[i-1].ClosePrice
		if change > 0 {
			gains = append(gains, change)
			losses = append(losses, 0)
		} else {
			losses = append(losses, -change)
			gains = append(gains, 0)
		}
	}

	// 计算初始的平均涨幅和平均跌幅
	avgGain := average(gains[:period])
	avgLoss := average(losses[:period])

	// 计算 RSI 并更新到 prices 对象中
	for i := period; i < int64(len(prices)); i++ {
		rs := avgGain / avgLoss
		rsiValue := 100 - (100 / (1 + rs))
		prices[i-1].Index.Rsi = rsiValue // 直接将 RSI 值写入 prices[i].RSI

		// 更新平均涨幅和跌幅
		if prices[i].ClosePrice > prices[i-1].ClosePrice {
			avgGain = (avgGain*(float64(period-1)) + (prices[i].ClosePrice - prices[i-1].ClosePrice)) / float64(period)
			avgLoss = (avgLoss * float64(period-1)) / float64(period)
		} else if prices[i].ClosePrice < prices[i-1].ClosePrice {
			avgLoss = (avgLoss*(float64(period-1)) + (prices[i-1].ClosePrice - prices[i].ClosePrice)) / float64(period)
			avgGain = (avgGain * float64(period-1)) / float64(period)
		} else {
			avgGain = avgGain * float64(period-1) / float64(period)
			avgLoss = avgLoss * float64(period-1) / float64(period)
		}
	}

	prices[len(prices)-1].Index.Rsi = 100 - (100 / (1 + avgGain/avgLoss))
	return prices
}

func average(values []float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func ConvertStringToObject(body string, v any) error {
	err := json.Unmarshal([]byte(body), v)
	if err != nil {
		return err
	}

	return nil
}

func RSICrossQK(line []model.LineItem, q bool, value float64) []model.LineItem {
	var res []model.LineItem
	for i := 1; i < len(line); i++ {
		current := line[i]
		if current.Index.Rsi == 0 {
			continue
		}

		if q && current.Index.Rsi > value {
			res = append(res, current)
		} else if !q && current.Index.Rsi < value {
			res = append(res, current)
		}
	}

	return res
}

func RSICrossUpOrDown(line []model.LineItem, isUp bool, value float64) []model.LineItem {
	var res []model.LineItem
	for i := 1; i < len(line); i++ {
		last := line[i-1]
		current := line[i]
		if last.Index.Rsi == 0 || current.Index.Rsi == 0 {
			continue
		}

		if isUp && last.Index.Rsi < value && current.Index.Rsi > value {
			res = append(res, current)
		} else if !isUp && last.Index.Rsi > value && current.Index.Rsi < value {
			res = append(res, current)
		}
	}

	return res
}

func MACDCross(line []model.LineItem, side bool) []model.LineItem {
	var res []model.LineItem
	for i := 1; i < len(line); i++ {
		last := line[i-1]
		current := line[i]
		if side && last.Index.Dif < last.Index.Dea && current.Index.Dif > current.Index.Dea {
			res = append(res, current)
		} else if !side && last.Index.Dif > last.Index.Dea && current.Index.Dif < current.Index.Dea {
			res = append(res, current)
		}
	}

	return res
}

func MACDLongOrSort(line []model.LineItem, side bool) []model.LineItem {
	var res []model.LineItem
	for i := 1; i < len(line); i++ {
		current := line[i]
		if side && current.Index.Dif > current.Index.Dea {
			res = append(res, current)
		} else if !side && current.Index.Dif < current.Index.Dea {
			res = append(res, current)
		}
	}

	return res
}
