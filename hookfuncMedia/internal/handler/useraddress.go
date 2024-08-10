package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/service"
	"net/http"
)

type UserAddressHandler struct {
	*Handler
	userAddressService service.UserAddressService
}

func NewUserAddressHandler(
	handler *Handler,
	userAddressService service.UserAddressService,
) *UserAddressHandler {
	return &UserAddressHandler{
		Handler:            handler,
		userAddressService: userAddressService,
	}
}

// Create godoc
//
//	@Summary	保存收货地址信息
//	@Schemes
//	@Description
//	@Tags		用户模块
//	@Accept		json
//	@Produce	json
//	@Param		request	body		v1.UserAddressInfoRequest	true	"params"
//	@Param		Authorization	header		string	true	"Authorization token"
//	@Success	200		{object}	v1.Response
//	@Router		/user/address/create [post]
func (h *UserAddressHandler) Create(ctx *gin.Context) {
	var req v1.UserAddressInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	req.UserId = GetUserIdFromCtx(ctx)
	userAddress, err := h.userAddressService.CreateUserAddress(ctx, &req)
	if err != nil {
		h.logger.WithContext(ctx).Error("【用户模块】保存用户收货地址失败", zap.Error(err))
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, userAddress)
}
