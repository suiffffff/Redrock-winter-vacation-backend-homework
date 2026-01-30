package router

import (
	"system/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	public := r.Group("/user")
	{
		public.POST("/register", handler.AddUser)
		public.GET("/check", handler.FindUserName)
	}
	return r
}
