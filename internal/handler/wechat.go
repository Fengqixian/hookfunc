package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type WechatHandler struct {
	*Handler
	wechatService service.WechatService
}

func NewWechatHandler(handler *Handler, wechatService service.WechatService) *WechatHandler {
	return &WechatHandler{
		Handler:       handler,
		wechatService: wechatService,
	}
}

func (h *WechatHandler) ProgramQrCodeLogin(ctx *gin.Context) {
	loginQrCodeResponse, err := h.wechatService.GetLoginQrCode(ctx)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, loginQrCodeResponse)
}

func (h *WechatHandler) ProgramLogin(ctx *gin.Context) {
	var req v1.WechatProgramLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	loginQrCodeResponse, err := h.wechatService.GetJsCodeToken(ctx, &req)
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, loginQrCodeResponse)
}
