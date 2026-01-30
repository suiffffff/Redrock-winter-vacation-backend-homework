package service

import (
	"system/dao"
	"system/dto"
	"system/models"
	"system/pkg"
)

func AddUser(req *dto.AddUserReq) error {
	usermodel := models.User{
		Username:   req.Username,
		Password:   pkg.Jiami(req.Password),
		Nickname:   req.Nickname,
		Department: req.Department,
	}
	return dao.AddUser(&usermodel)
}
func FindUserName(username string) (bool, error) {
	return dao.FindUserName(username)
}
func Login(req *dto.LoginReq) error {
	usermodel := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	return dao.Login(&usermodel)
}
