package handler

import (
	"system/dto"
	"system/service"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var req dto.AddUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"err": "参数错误:" + err.Error()})
		return

	}
	if err := service.AddUser(req); err != nil {
		c.JSON(500, gin.H{"err": "注册失败" + err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "注册成功"})
}
func FindUserName(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(400, gin.H{"error": "用户名不能为空"})
		return
	}

	exists, err := service.FindUserName(username)
	if err != nil {
		c.JSON(500, gin.H{"error": "系统错误"})
		return
	}
	if exists {
		c.JSON(200, gin.H{"exists": true, "msg": "用户名已存在"})
	} else {
		c.JSON(200, gin.H{"exists": false, "msg": "用户名可用"})
	}
}
