package v1

import (
	"hookfunc-media/internal/model"
	"time"
)

type PlaceOrderRequest struct {
	UserId       int64
	AddressId    int64               `json:"addressId" binding:"required"`
	DeliveryTime *time.Time          `json:"deliveryTime" binding:"required"`
	Remark       string              `json:"remark"`
	OrderGoods   *[]model.OrderGoods `json:"orderGoods" binding:"required"`
}

type CancelOrderRequest struct {
	OrderId int64 `json:"orderId" binding:"required"`
}

type OrderInfoResponse struct {
	UserId        int64        `json:"userId"`        // 用户ID
	OrderId       int64        `json:"orderId" `      // 订单ID
	Province      string       `json:"province"`      // 省
	City          string       `json:"city"`          // 市
	Region        string       `json:"region"`        // 区
	Longitude     int64        `json:"longitude"`     // 经度
	Latitude      int64        `json:"latitude"`      // 纬度
	UserName      string       `json:"userName"`      // 收获人姓名
	PhoneNumber   string       `json:"phoneNumber"`   // 收获人电话
	AddressDetail string       `json:"addressDetail"` // 详细地址
	Remark        string       `json:"remark"`        // 订单备注
	PlaceTime     time.Time    `json:"placeTime"`     // 下单时间
	Total         int64        `json:"total"`         // 订单金额
	PaymentModel  int32        `json:"paymentModel"`  // 支付方式
	OrderState    int32        `json:"orderState"`    // 订单状态
	Goods         []OrderGoods `json:"goods"`         // 商品列表
}

type OrderGoods struct {
	GoodsId   int64  `json:"goodsId"`   // 商品ID
	GoodsName string `json:"goodsName"` // 商品名称
	Amount    int64  `json:"amount"`    // 商品数量
	Price     int64  `json:"price"`     // 单价
	Total     int64  `json:"total"`     // 商品金额
}
