package controller

import (
	"encoding/json"
	"fmt"
	"go_gateway/dto"
	"go_gateway/middleware"
	"go_gateway/public"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/admin_info", adminLogin.AdminInfo)
}

// AdminInfo godoc
// @Summary 管理员信息获取
// @Description 管理员信息获取
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminInfo *AdminController) AdminInfo(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	// 1.读取sessionKey对应json，转化为结构体
	// 2. 取出数据然后封装输出结构体
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LogInTime:    adminSessionInfo.LogInTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(ctx, out)
}
