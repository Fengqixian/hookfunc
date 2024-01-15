package v1

type WechatProgramLoginRequest struct {
	JsCode string `json:"jsCode" binding:"required"`
}
