package response

import "net/http"

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type Result struct {
	Status     string      `json:"status"`
	Msg        string      `json:"msg"`
	StatusCode int         `json:"status_code"`
	Code       string      `json:"code"`
	Data       interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		StatusCode: http.StatusOK,
		Code:       "0",
		Msg:        StatusSuccess,
		Status:     StatusSuccess,
	}
}
