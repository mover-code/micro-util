package xhttp

type ResponseSuccessBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Error(errCode int, errMsg string) *ResponseErrorBean {
	empty := make([]interface{}, 0)
	return &ResponseErrorBean{errCode, errMsg, empty}
}
