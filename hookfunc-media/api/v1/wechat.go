package v1

type WechatProgramLoginRequest struct {
	JsCode string `json:"jsCode" binding:"required"`
}

// LoginQrCodeResponse 登录二维码
type LoginQrCodeResponse struct {
	QrCode    string `json:"qrCode"`
	LoginCode string `json:"loginCode"`
}
