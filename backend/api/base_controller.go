package api

import (
	"eve-plugin-vue3-template/backend/response"
)

// 父控制器结构体
type BaseController struct {
	*response.Response
}

func NewBaseController(response *response.Response) *BaseController {
	return &BaseController{Response: response}
}
