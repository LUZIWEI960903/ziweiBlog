package service

import (
	"errors"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func Login(username, passwd string) (*models.LoginRes, error) {
	user := dao.GetUser(username, passwd)
	if user == nil {
		return nil, errors.New("username or passwd error")
	}

	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.Username = user.Username
	userInfo.Avater = user.Avater

	var ls = &models.LoginRes{
		Token:    "",
		UserInfo: userInfo,
	}

	return ls, nil
}
