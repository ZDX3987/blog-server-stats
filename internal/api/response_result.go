package api

import "net/http"

type ResponseResult struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Success() *ResponseResult {
	return &ResponseResult{"success", http.StatusOK, nil}
}

func SuccessResult(data interface{}) *ResponseResult {
	return &ResponseResult{"success", http.StatusOK, data}
}

func Fail(errorMsg string) *ResponseResult {
	return &ResponseResult{errorMsg, http.StatusInternalServerError, nil}
}
