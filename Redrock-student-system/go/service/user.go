package service

import (
	"system/dao"
	"system/dto"
	"system/models"
	"system/pkg"
	"time"
)

func FindUserName(req *dto.FindUserNameReq) (bool, error) {
	usermodel := models.User{
		Username: req.Username,
	}
	return dao.FindUserName(&usermodel)
}
func AddUser(req *dto.AddUserReq) error {
	usermodel := models.User{
		Username:   req.Username,
		Password:   pkg.Jiami(req.Password),
		Nickname:   req.Nickname,
		Department: req.Department,
	}
	return dao.AddUser(&usermodel)
}

func Login(req *dto.LoginReq) (string, string, error) {
	usermodel := models.User{
		Username: req.Username,
		Password: pkg.Jiami(req.Password),
	}
	user, err := dao.Login(&usermodel)
	if err != nil {
		return "", "", err
	}
	accessToken, refreshToken, err := pkg.GenerateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	exp := time.Now().Add(7 * 24 * time.Hour).Unix()
	tokenmodel := models.UserToken{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    exp,
		Revoked:      false,
	}
	if err := dao.StoreRefreshToken(&tokenmodel); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}
