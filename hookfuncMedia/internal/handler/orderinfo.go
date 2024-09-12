package handler

import (
	"github.com/gin-gonic/gin"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/service"
	"net/http"
)

type OrderInfoHandler struct {
	*Handler
	orderInfoService service.OrderInfoService
}

func NewOrderInfoHandler(handler *Handler, orderInfoService service.OrderInfoService) *OrderInfoHandler {
	return &OrderInfoHandler{
		Handler:          handler,
		orderInfoService: orderInfoService,
	}
}

// Place godoc
//
//	@Summary	下单
//	@Schemes
//	@Description
//	@Tags		订单
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.PlaceOrderRequest	true	"params"
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200		{object}	v1.Response
//	@Router		/order/place [POST]
func (h *OrderInfoHandler) Place(ctx *gin.Context) {
	var req v1.PlaceOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	req.UserId = GetUserIdFromCtx(ctx)
	order, err := h.orderInfoService.PlaceAnOrder(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, order)
}

// Cancel godoc
//
//	@Summary	取消订单
//	@Schemes
//	@Description
//	@Tags		订单
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.CancelOrderRequest	true	"params"
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200		{object}	v1.Response
//	@Router		/order/cancel [POST]
func (h *OrderInfoHandler) Cancel(ctx *gin.Context) {
	var req v1.CancelOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	err := h.orderInfoService.CancelOrder(ctx, req.OrderId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// List godoc
//
//	@Summary	历史订单
//	@Schemes
//	@Description
//	@Tags		订单
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	v1.OrderInfoResponse
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Router		/order/list [GET]
func (h *OrderInfoHandler) List(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	v1.HandleSuccess(ctx, h.orderInfoService.ListOrder(ctx, userId))
}
