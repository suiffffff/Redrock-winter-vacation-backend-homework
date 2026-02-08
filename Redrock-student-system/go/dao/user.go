package dao

import "system/models"

// 写到这里我开始想，user的model里的数据相当多，那么对于一些要不了那么多字段的功能函数，会出现什么问题？
// 于是有了一个中转的dto层
func FindUserName(user *models.User) (bool, error) {
	var count int64
	username := user.Username
	err := DB.Model(&models.User{}).Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func AddUser(user *models.User) error {
	return DB.Create(user).Error
}
func Login(user *models.User) (*models.User, error) {
	name := user.Username
	password := user.Password
	err := DB.Where("name=? AND password=?", name, password).First(user).Error
	return user, err
}
func StoreRefreshToken(token *models.UserToken) error {
	return DB.Create(token).Error
}
func RefreshToken() {
	return
}
