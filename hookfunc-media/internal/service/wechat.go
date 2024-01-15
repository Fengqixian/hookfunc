package service

import (
	"github.com/gin-gonic/gin"
	v1 "hookfunc-media/api/v1"
)

type WechatService interface {
	ProgramLogin(ctx *gin.Context, params *v1.WechatProgramLoginRequest) (string, error)
}

func NewWechatService(service *Service) WechatService {
	return &wechatService{
		Service: service,
	}
}

type wechatService struct {
	*Service
}

func (w wechatService) ProgramLogin(ctx *gin.Context, params *v1.WechatProgramLoginRequest) (string, error) {
	//TODO implement me
	return "tokenlalalal", nil
}
