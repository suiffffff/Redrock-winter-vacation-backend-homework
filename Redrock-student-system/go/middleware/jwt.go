package middleware

import (
	"strings"
	"system/pkg"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer") {
			pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, "未登录或Token格式错误")
			c.Abort()
			return
		}
		claims, err := pkg.VerifyAccessToken(tokenHeader)
		if err != nil {
			pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, "Token无效: "+err.Error())
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(403, gin.H{
				"code": 403,
				"msg":  "不该碰的东西不要碰，你把握不住",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
