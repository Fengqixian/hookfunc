package service

import (
	"encoding/base64"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	v1 "hookfunc-media/api/v1"
	"hookfunc-media/internal/repository"
)

type wechatService struct {
	*Service
	*repository.Repository
}

type WechatService interface {
	GetLoginQrCode() (*v1.LoginQrCodeResponse, error)
}

func NewWechatService(service *Service, repository *repository.Repository) WechatService {
	return &wechatService{
		Service:    service,
		Repository: repository,
	}
}

func (w wechatService) GetLoginQrCode() (*v1.LoginQrCodeResponse, error) {
	qrcodeParams := qrcode.QRCoder{
		Scene: "100001",
	}
	codeBytes, err := w.GetWechatMiniProgram().GetQRCode().GetWXACodeUnlimit(qrcodeParams)
	if err != nil {
		return nil, err
	}

	return &v1.LoginQrCodeResponse{
		QrCode:    base64.StdEncoding.EncodeToString(codeBytes),
		LoginCode: "1231",
	}, nil
}
