package okx

import (
	"hookfunc/internal/model"
)

type Api struct {
	Http *HttpOkx
}

func NewApi(baseURL string, retry int) *Api {
	http := NewHttpOkx(baseURL, retry)
	return &Api{
		Http: http,
	}
}

func (a *Api) GetLine(instId string, bar string, limit string) (*model.LineResult, error) {
	var lineResult model.LineResult
	err := a.Http.Get("/api/v5/market/candles?instId="+instId+"&bar="+bar+"&limit="+limit, &lineResult)
	if err != nil {
		return nil, err
	}

	return &lineResult, nil
}
