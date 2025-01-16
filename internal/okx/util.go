package okx

import (
	"hookfunc/internal/model"
)

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
