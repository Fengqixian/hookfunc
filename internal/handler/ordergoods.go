package handler

import (
	"github.com/gin-gonic/gin"
	"hookfunc/internal/service"
)

type OrderGoodsHandler struct {
	*Handler
	orderGoodsService service.OrderGoodsService
}

func NewOrderGoodsHandler(handler *Handler, orderGoodsService service.OrderGoodsService) *OrderGoodsHandler {
	return &OrderGoodsHandler{
		Handler:           handler,
		orderGoodsService: orderGoodsService,
	}
}

func (h *OrderGoodsHandler) GetOrderGoods(ctx *gin.Context) {

}
