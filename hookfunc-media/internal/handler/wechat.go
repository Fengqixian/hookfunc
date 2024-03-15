package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/service"
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

// ProgramLogin godoc
//
//	@Summary	微信小程序用户登录
//	@Schemes
//	@Description
//	@Tags		微信小程序
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.WechatProgramLoginRequest	true	"params"
//	@Success	200		{object}	v1.Response
//	@Router		/wechat/program/login [post]
func (h *WechatHandler) ProgramLogin(ctx *gin.Context) {
	req := &v1.WechatProgramLoginRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	loginQrCodeResponse, err := h.wechatService.GetLoginQrCode()
	if err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, loginQrCodeResponse)
}
