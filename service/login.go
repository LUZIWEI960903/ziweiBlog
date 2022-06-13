package service

import (
	"errors"
	"ziweiBlog/dao"
	"ziweiBlog/models"
	"ziweiBlog/utils"
)

func Login(username, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "ziwei")
	user := dao.GetUser(username, passwd)
	if user == nil {
		return nil, errors.New("username or passwd error")
	}

	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.Username = user.Username
	userInfo.Avater = user.Avater

	// create token JWT
	token, err := utils.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("create Token error")
	}

	var ls = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}

	return ls, nil
}
