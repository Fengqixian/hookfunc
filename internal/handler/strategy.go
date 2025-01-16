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

// CreateStrategy godoc
//
//	@Summary	创建策略
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header	string						true	"Authorization token"
//	@Param		request			body	v1.CreateStrategyRequest	true	"params"
//	@Success	200
//	@Router		/strategy/create [post]
func (h *StrategyHandler) CreateStrategy(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	var req v1.CreateStrategyRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	req.UserId = userId
	err := h.strategyService.CreateStrategy(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, err)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// ListStrategyIndex godoc
//
//	@Summary	获取策略关联指标
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header	string				true	"Authorization token"
//	@Param		request			body	v1.StrategyRequest	true	"params"
//	@Success	200
//	@Router		/strategy/index/list [post]
func (h *StrategyHandler) ListStrategyIndex(ctx *gin.Context) {
	var req v1.StrategyRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	userId := GetUserIdFromCtx(ctx)
	req.UserId = userId
	strategies, err := h.strategyService.ListStrategyIndex(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, strategies)
}

// DeleteStrategyIndex godoc
//
//	@Summary	删除策略关联指标
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header	string				true	"Authorization token"
//	@Param		request			body	v1.StrategyIndexRequest	true	"params"
//	@Success	200
//	@Router		/strategy/index/delete [post]
func (h *StrategyHandler) DeleteStrategyIndex(ctx *gin.Context) {
	var req v1.StrategyIndexRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	userId := GetUserIdFromCtx(ctx)
	req.UserId = userId
	err := h.strategyService.DeleteStrategyIndex(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// DeleteStrategy godoc
//
//	@Summary	删除策略
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header	string				true	"Authorization token"
//	@Param		request			body	v1.StrategyRequest	true	"params"
//	@Success	200
//	@Router		/strategy/delete [post]
func (h *StrategyHandler) DeleteStrategy(ctx *gin.Context) {
	var req v1.StrategyRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	userId := GetUserIdFromCtx(ctx)
	req.UserId = userId
	err := h.strategyService.DeleteStrategy(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// UpdateStrategySubscriptionState godoc
//
//	@Summary	更新策略订阅状态
//	@Schemes
//	@Description
//	@Tags		策略
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header	string				true	"Authorization token"
//	@Param		request			body	v1.StrategyRequest	true	"params"
//	@Success	200
//	@Router		/strategy/subscription [post]
func (h *StrategyHandler) UpdateStrategySubscriptionState(ctx *gin.Context) {
	var req v1.StrategyRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	userId := GetUserIdFromCtx(ctx)
	req.UserId = userId
	err := h.strategyService.UpdateStrategySubscriptionState(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
