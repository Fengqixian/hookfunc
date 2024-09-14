package model

//接口请求响应数据结构

type XPYunResp struct {
	HttpStatusCode int `json:"httpStatusCode"`

	Content *XPYunRespContent `json:"content"`
}

type XPYunRespContent struct {
	Code               int         `json:"code"`
	Msg                string      `json:"msg"`
	Data               interface{} `json:"data,omitempty"`
	ServerExecutedTime int         `json:"serverExecutedTime"`
}
