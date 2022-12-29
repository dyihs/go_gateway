package controller

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/middleware"
)

type DashboardController struct{}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/panelGroupData", service.PanelGroupData)
}

// PanelGroupData godoc
// @Summary 指标统计
// @Description 指标统计
// @Tags 首页大盘
// @ID /dashboard/panelGroupData
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.PanelGroupDataOutput} "success"
// @Router /dashboard/panelGroupData [get]
func (service *DashboardController) PanelGroupData(c *gin.Context) {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err := serviceInfo.PageList(c, tx, &dto.ServiceListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	app := &dao.App{}
	_, appNum, err := app.APPList(c, tx, &dto.APPListInput{PageSize: 1, PageNo: 1})
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	out := &dto.PanelGroupDataOutput{
		ServiceNum:      serviceNum,
		AppNum:          appNum,
		CurrentQPS:      0,
		TodayRequestNum: 0,
	}
	middleware.ResponseSuccess(c, out)
}
