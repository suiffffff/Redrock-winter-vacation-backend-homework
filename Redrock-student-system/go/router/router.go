package router

import (
	"system/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	public := r.Group("/user")
	{
		public.GET("/check", handler.FindUserName)
		public.POST("/register", handler.AddUser)
		public.POST("/login", handler.Login)
		public.POST("/refresh", handler.RefreshToken)
	}
	return r
}
