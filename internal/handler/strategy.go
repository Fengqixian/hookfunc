package handler

import (
	"github.com/gin-gonic/gin"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type StrategyHandler struct {
	*Handler
	strategyService service.StrategyService
}

func NewStrategyHandler(handler *Handler, strategyService service.StrategyService) *StrategyHandler {
	return &StrategyHandler{
		Handler:         handler,
		strategyService: strategyService,
	}
}

// ListStrategy godoc
//
//	@Summary	获取当前用户的策略列表
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200				{object}	model.Strategy
//	@Router		/strategy/list [get]
func (h *StrategyHandler) ListStrategy(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	strategies, err := h.strategyService.ListStrategy(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, strategies)
}
