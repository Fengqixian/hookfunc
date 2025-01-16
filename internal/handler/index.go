package handler

import (
	"github.com/gin-gonic/gin"
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
//	@Success	200				{object}	model.Index
//	@Router		/index/list [get]
func (h *IndexHandler) ListIndex(ctx *gin.Context) {
	index, err := h.indexService.ListIndex(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, v1.ErrServer, nil)
		return
	}

	v1.HandleSuccess(ctx, index)
}
