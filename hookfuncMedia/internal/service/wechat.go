package service

import (
	"context"
	"encoding/base64"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"go.uber.org/zap"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/repository"
	"math/rand"
	"time"
)

const defaultLoginCode = "defaultLoginCode"
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const loginCodeCacheKeyPrefix = "service:login:code:"

type wechatService struct {
	*Service
	*repository.Repository
}

func (w wechatService) GetLoginQrCode(context context.Context) (*v1.LoginQrCodeResponse, error) {
	loginCode := w.GenRandomCode(context)
	qrcodeParams := qrcode.QRCoder{
		Scene: loginCode,
	}
	codeBytes, err := w.GetWechatMiniProgram().GetQRCode().GetWXACodeUnlimit(qrcodeParams)
	if err != nil {
		return nil, err
	}

	return &v1.LoginQrCodeResponse{
		QrCode:    base64.StdEncoding.EncodeToString(codeBytes),
		LoginCode: loginCode,
	}, nil
}

func (w wechatService) GenRandomCode(context context.Context) string {
	var randStr string
	for {
		randStr = randStringBytes(10)
		exists, err := w.Repository.GetRedisClient().Exists(context, randStr).Result()
		if err != nil {
			w.Service.logger.WithContext(context).Error("【获取登录二维码】检查LoginCode是否存在失败: {}", zap.Error(err))
			break
		}

		if exists == 0 {
			break
		}
	}
	if len(randStr) == 0 {
		return defaultLoginCode
	}

	w.Repository.GetRedisClient().Set(context, loginCodeCacheKeyPrefix+randStr, "1", time.Minute*3)
	return randStr
}

type WechatService interface {
	GetLoginQrCode(context context.Context) (*v1.LoginQrCodeResponse, error)
	GenRandomCode(context context.Context) string
}

func NewWechatService(service *Service, repository *repository.Repository) WechatService {
	return &wechatService{
		Service:    service,
		Repository: repository,
	}
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
