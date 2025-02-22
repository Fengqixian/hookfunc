package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type IndexHandler struct {
	*Handler
	indexService service.IndexService
}

func NewIndexHandler(handler *Handler, indexService service.IndexService) *IndexHandler {
	return &IndexHandler{
		Handler:      handler,
		indexService: indexService,
	}
}

// ListIndex godoc
//
//	@Summary	指标列表
//	@Schemes
//	@Description
//	@Tags		指标
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{object}	model.Index
//	@Router		/index/list [get]
func (h *IndexHandler) ListIndex(ctx *gin.Context) {
	index, err := h.indexService.ListIndex(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, index)
}

// IndexTest godoc
//
//	@Summary	指标回测
//	@Schemes
//	@Description
//	@Tags		指标
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		request			body	v1.IndexRequest	true	"params"
//	@Success	200	{object}	model.Index
//	@Router		/index/test [post]
func (h *IndexHandler) IndexTest(ctx *gin.Context) {
	var req v1.IndexRequest
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrRequestParamsFail, err)
		return
	}

	result, err := h.indexService.IndexHitTarget(ctx, req)
	if err != nil {
		h.logger.Error("【指标回测失败】", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, result)
}
