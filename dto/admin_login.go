package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	LogInTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,is_valid_username"`// 管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`// 密码
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`// token
}

func (param *AdminLoginInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}
