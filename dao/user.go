package dao

import (
	"log"
	"ziweiBlog/models"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select username from blog_user where uid = ?", userId)
	if row.Err() != nil {
		log.Println("GetUserNameById [DB.QueryRow] error:", row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where username = ? and passwd = ? limit 1", username, passwd)
	if row.Err() != nil {
		log.Println("GetUser [DB.QueryRow] error:", row.Err())
		return nil
	}

	var user = new(models.User)
	err := row.Scan(&user.Uid, &user.Username, &user.Passwd, &user.Avater, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("GetUser [row.Scan] error:", err)
		return nil
	}

	return user
}
