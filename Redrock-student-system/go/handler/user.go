package handler

import (
	"system/dto"
	"system/pkg"
	"system/service"

	"github.com/gin-gonic/gin"
)

func FindUserName(c *gin.Context) {
	var req dto.FindUserNameReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "用户名不能为空")
		return
	}
	exists, err := service.FindUserName(&req)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "系统查询错误")
		return
	}
	if exists {
		pkg.Success(c, gin.H{"exists": true, "msg": "用户名已存在"})
	} else {
		pkg.Success(c, gin.H{"exists": false, "msg": "用户名可用"})
	}
}
func AddUser(c *gin.Context) {
	var req dto.AddUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "参数错误: "+err.Error())
		return
	}
	if err := service.AddUser(&req); err != nil {
		pkg.Error(c, pkg.CodeSystemError, "注册失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"msg": "注册成功"})
}
func Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "参数错误")
		return
	}
	at, rt, err := service.Login(&req)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "系统查询错误")
		return
	}
	pkg.Success(c, dto.LoginRes{
		AccessToken:  at,
		RefreshToken: rt,
	})
}
func RefreshToken(c *gin.Context) {

}
