package handler

import (
	"github.com/gin-gonic/gin"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type BarHandler struct {
	*Handler
	barService service.BarService
}

func NewBarHandler(handler *Handler, barService service.BarService) *BarHandler {
	return &BarHandler{
		Handler:    handler,
		barService: barService,
	}
}

// ListBar godoc
//
//	@Summary	时间线
//	@Schemes
//	@Description
//	@Tags		指标
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
