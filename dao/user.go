package dao

import (
	"fmt"
	"ms-go-blog/logger"
	"ms-go-blog/models"
	"sort"
	"time"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		logger.Error(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where user_name=? and passwd = ? limit 1",
		userName,
		passwd,
	)
	if row.Err() != nil {
		logger.Error(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return user
}

func CountUserName(userName string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_user where user_name=?", userName)
	_ = rows.Scan(&count)
	return
}

func InsertNewUser(userName, passwd string) error {
	creat_at := time.Now()
	update_at := time.Now()

	var uidSlice []int
	rows, _ := DB.Query("select uid from blog_user")
	defer rows.Close()

	for rows.Next() {
		var a int
		rows.Scan(&a)
		uidSlice = append(uidSlice, a)
	}

	sort.Ints(uidSlice)

	maxUid := uidSlice[len(uidSlice)-1]
	maxUid += 1

	fmt.Println(userName, passwd)
	_, err := DB.Exec("insert into blog_user "+
		"(uid, user_name,passwd,avatar,create_at,update_at) "+
		"values(?,?,?,?,?,?)",
		maxUid,
		userName,
		passwd,
		"",
		creat_at,
		update_at,
	)
	if err != nil {
		return err
	}
	return nil
}
