// Code generated by goctl. DO NOT EDIT.
package types

import (
	"go-zero-demo/greet/internal/svc/constant"
)

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//types.Response{Code: -1, Msg: "获取失败"}

//DEFAULT_ERROR_CODE
func NewCodeError(code int, msg string) *Response {
	return &Response{Code: code, Msg: msg}
}

func NewDefaultError(msg string) *Response {
	return NewCodeError(constant.DEFAULT_ERROR_CODE, msg)
}

type PageReq struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页大小
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}