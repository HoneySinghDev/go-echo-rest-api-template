package types

type ApiRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiRes(code int, message string) *ApiRes {
	return &ApiRes{
		Code:    code,
		Message: message,
	}
}
