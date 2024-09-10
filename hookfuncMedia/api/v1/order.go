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
