package okx

import (
	"fmt"
	"hookfunc/internal/model"
)

type Api struct {
	Http   *Http
	Config *Config
}

func NewApi(c *Config) *Api {
	http := NewHttp(c)
	return &Api{
		Http:   http,
		Config: c,
	}
}

func (a *Api) GetLine(instId string, bar string, limit int) (*model.LineResult, error) {
	var lineResult model.LineResult
	err := a.Http.Get(a.Config.Server+"/api/v5/market/candles?instId="+instId+"&bar="+bar+"&limit="+fmt.Sprintf("%v", limit), &lineResult)
	if err != nil {
		return nil, err
	}

	if lineResult.Code != "0" {
		return nil, fmt.Errorf("【OKX获取K线失败】error code: %v, message: %v", lineResult.Code, lineResult.Msg)
	}

	return &lineResult, nil
}

func (a *Api) GetRechargeRecord(instId string) (*[]model.OkxRechargeRecord, error) {
	var lineResult model.OkxRechargeResponse
	err := a.Http.Get(a.Config.Server+"/api/v5/asset/deposit-history?ccy="+instId, &lineResult)
	if err != nil {
		return nil, err
	}

	if lineResult.Code != "0" {
		return nil, fmt.Errorf("【OKX获取充值记录失败】error code: %v, message: %v", lineResult.Code, lineResult.Msg)
	}

	return &lineResult.Data, nil
}
