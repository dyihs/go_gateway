package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gateway/public"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		if adminInfo, ok := sess.Get(public.AdminSessionInfoKey).(string); !ok || adminInfo == "" {
			ResponseError(ctx, InternalErrorCode, errors.New("user not login"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
