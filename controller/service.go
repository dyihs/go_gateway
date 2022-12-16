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

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param polygon body dto.ServiceListInput true "body"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
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
