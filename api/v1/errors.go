package v1

var (
	ErrSuccess      = newError(0, "ok")
	ErrBadRequest   = newError(400, "Bad Request")
	ErrUnauthorized = newError(401, "Unauthorized")
	ErrNotFound     = newError(404, "Not Found")

	ErrServer = newError(500, "服务繁忙，请稍后再试")

	ErrRequestParamsFail = newError(50000, "请求参数错误")
	ErrPlaceAnOrderFail  = newError(50001, "下单失败")
)
