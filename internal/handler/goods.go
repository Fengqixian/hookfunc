package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type GoodsHandler struct {
	*Handler
	goodsService service.GoodsService
}

func NewGoodsHandler(
	handler *Handler,
	goodsService service.GoodsService,
) *GoodsHandler {
	return &GoodsHandler{
		Handler:      handler,
		goodsService: goodsService,
	}
}

// Info godoc
//
//	@Summary	通过商品ID获取商品信息
//	@Schemes
//	@Description
//	@Tags		商品
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.Goods	true	"params"
//	@Success	200		{object}	v1.Response
//	@Router		/goods/info [POST]
func (h *GoodsHandler) Info(ctx *gin.Context) {
	var req v1.Goods
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	response, err := h.goodsService.GetGoods(ctx, &req)
	if err != nil {
		h.logger.WithContext(ctx).Error("【商品】通过ID获取商品失败", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, response)
}

// List godoc
//
//	@Summary	获取所有商品信息
//	@Schemes
//	@Description
//	@Tags		商品
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	v1.Response
//	@Router		/goods/list [GET]
func (h *GoodsHandler) List(ctx *gin.Context) {
	response, err := h.goodsService.GetAllGoods(ctx)
	if err != nil {
		h.logger.WithContext(ctx).Error("【商品】通过ID获取商品失败", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, response)
}
