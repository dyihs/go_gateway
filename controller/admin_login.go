package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/dto"

	"go_gateway/middleware"
	"go_gateway/public"
	"time"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLogin(ctx *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	// 1. params.UserName 取得管理员信息 admininfo
	// 2. admininfo.salt + params.Password sha256 => saltPassword
	// 3. saltPassword==admininfo.password
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}

	// 设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LogInTime: time.Now(),
	}
	sessByte, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	sess := sessions.Default(ctx)
	sess.Set(public.AdminSessionInfoKey, string(sessByte))
	sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(ctx, out)
}
