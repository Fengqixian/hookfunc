package handler

import (
	"github.com/gin-gonic/gin"
	"hookfunc/api/v1"
	"hookfunc/internal/service"
	"net/http"
)

type UserHandler struct {
	*Handler
	userInfoService service.UserInfoService
}

func NewUserHandler(handler *Handler, userInfoService service.UserInfoService) *UserHandler {
	return &UserHandler{
		Handler:         handler,
		userInfoService: userInfoService,
	}
}

// GetProfile godoc
//
//	@Summary	获取用户信息
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200				{object}	v1.GetProfileResponse
//	@Router		/user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	userinfo, err := h.userInfoService.GetUserInfoById(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, nil)
		return
	}

	user, err := h.userInfoService.GetUserInfo(ctx, userinfo.Openid)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, user)
}

// SendSmsCode godoc
//
//	@Summary	发送短信验证码
//	@Schemes
//	@Description
//	@Tags		短信
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		request	body	v1.SendSMSCodeRequest	true	"params"
//	@Success	200
//	@Router		/sms/code [post]
func (h *UserHandler) SendSmsCode(ctx *gin.Context) {
	var req v1.SendSMSCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	err := h.userInfoService.SendNoteVerificationCode(ctx, req.PhoneNumber)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}

// VerificationSmsCode godoc
//
//	@Summary	核对验证码
//	@Schemes
//	@Description
//	@Tags		短信
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		request	body	v1.SendSMSCodeRequest	true	"params"
//	@Success	200
//	@Router		/verification/sms/code [post]
func (h *UserHandler) VerificationSmsCode(ctx *gin.Context) {
	var req v1.SendSMSCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	token, err := h.userInfoService.VerificationCode(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, token)
}

// ConfirmRechargeRecord godoc
//
//	@Summary	充值确认
//	@Schemes
//	@Description
//	@Tags		交易
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200
//	@Router		/transaction/recharge/confirm [get]
func (h *UserHandler) ConfirmRechargeRecord(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	result, err := h.userInfoService.ConfirmRechargeRecord(ctx, userId)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, result)
}
