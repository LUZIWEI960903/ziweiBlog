package dao

import "log"

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select username from blog_user where uid = ?", userId)
	if row.Err() != nil {
		log.Println("GetUserNameById [DB.QueryRow] error:", row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}
