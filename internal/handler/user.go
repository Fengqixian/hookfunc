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
