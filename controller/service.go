package controller

import (
	"github.com/gin-gonic/gin"
	"go_gateway/dto"
	"go_gateway/middleware"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
}

func (service ServiceController) ServiceList(ctx *gin.Context) {
	params := &dto.ServiceListInput{}
	err := params.BindValidParam(ctx)
	if err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	out := &dto.ServiceListOutput{}
	middleware.ResponseSuccess(ctx, out)

}
