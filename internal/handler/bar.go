package handler

import (
	"github.com/gin-gonic/gin"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/okx"
	"hookfunc/internal/service"
	"net/http"
)

type BarHandler struct {
	*Handler
	barService service.BarService
	*okx.Config
}

func NewBarHandler(config *okx.Config, handler *Handler, barService service.BarService) *BarHandler {
	return &BarHandler{
		Handler:    handler,
		barService: barService,
		Config:     config,
	}
}

// ListBar godoc
//
//	@Summary	时间线
//	@Schemes
//	@Description
//	@Tags		公共
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	model.Bar
//	@Router		/index/bar/list [get]
func (h *BarHandler) ListBar(ctx *gin.Context) {
	bars, err := h.barService.ListBar(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, bars)
}

// ListCoin godoc
//
//	@Summary	币币信息
//	@Schemes
//	@Description
//	@Tags		公共
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	model.Bar
//	@Router		/coin/list [get]
func (h *BarHandler) ListCoin(ctx *gin.Context) {
	coins, err := h.barService.ListCoin(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, coins)
}

// ListSubscriptionPrice godoc
//
//	@Summary	获取订阅价格
//	@Schemes
//	@Description
//	@Tags		公共
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	model.Bar
//	@Router		/subscription/price [get]
func (h *BarHandler) ListSubscriptionPrice(ctx *gin.Context) {
	v1.HandleSuccess(ctx, h.Config.SubscriptionPrice)
}

// GetSubscriptionPayWallet godoc
//
//	@Summary	获取订阅支付钱包地址
//	@Schemes
//	@Description
//	@Tags		公共
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	string
//	@Router		/subscription/pay/wallet [get]
func (h *BarHandler) GetSubscriptionPayWallet(ctx *gin.Context) {
	v1.HandleSuccess(ctx, h.Config.WalletAddress)
}
