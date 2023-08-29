package middlewares

import (
	"chatgpt-web-new-go/common/auth"
	"chatgpt-web-new-go/router/base"
	"github.com/gin-gonic/gin"
)

// Jwt jwt认证
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.EncodeByCtx(c)
		if err != nil {
			base.Fail(c, "用户信息错误: "+err.Error())
			c.Abort()
			return
		}

		if claims.User.ID == 0 {
			base.Fail(c, "用户信息错误，未知的token")
			c.Abort()
			return
		}
		c.Set(auth.GinCtxKey, claims.User)
		c.Next()
	}
}
