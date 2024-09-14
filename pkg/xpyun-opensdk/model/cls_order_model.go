package model

// 根据订单号或打印机编号清除打印任务
type ClearPrintOrderRequest struct {
	RestRequest `json:",inline"`

	/**
	 * 打印机编号
	 */
	Sn string `json:"sn"`
	/**
	 * 订单编号
	 */
	OrderId string `json:"orderId,omitempty"`
}
