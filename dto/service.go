package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      // 关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        // 页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` // 每页条数
}

func (param *ServiceListInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" from:"id"`                     // id
	ServiceName string `json:"service_name" from:"service_name"` // 服务名称
	ServiceDesc string `json:"service_desc" from:"service_desc"` // 服务描述
	LoadType    int    `json:"load_type" from:"load_type"`       // 类型
	ServiceAddr string `json:"service_addr" from:"service_addr"` // 服务地址
	Qps         int64  `json:"qps" from:"qps"`                   // qps
	Qpd         int64  `json:"qpd" from:"qpd"`                   // qpd
	TotalNode   int    `json:"total_node" from:"total_node"`     // 节点数
}

type ServiceListOutput struct {
	Total int64  `json:"total" form:"total" comment:"总数" example:"" validate:""`     // 总素
	List  string `json:"list" form:"list" comment:"" example:"" validate:"required"` //列表
}
