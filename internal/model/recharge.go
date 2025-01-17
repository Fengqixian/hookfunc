package model

type OkxRechargeResponse struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []OkxRechargeRecord `json:"data"`
}

type OkxRechargeRecord struct {
	ActualDepBlkConfirm string `json:"actualDepBlkConfirm"`
	Amt                 string `json:"amt"`
	AreaCodeFrom        string `json:"areaCodeFrom"`
	Ccy                 string `json:"ccy"`
	Chain               string `json:"chain"`
	DepId               string `json:"depId"`
	From                string `json:"from"`
	FromWdId            string `json:"fromWdId"`
	State               string `json:"state"`
	To                  string `json:"to"`
	Ts                  string `json:"ts"`
	TxId                string `json:"txId"`
}
