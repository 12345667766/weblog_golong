package service

import (
	"errors"
	"ms-go-blog/dao"
	"ms-go-blog/logger"
	"ms-go-blog/utils"
)

func Register(userName, passwd string) error {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	count := dao.CountUserName(userName)
	if count != 0 {
		logger.Error("当前用户%s已存在, 无需再次注册\n", count)
		return errors.New("用户已存在，无需注册")
	}
	err := dao.InsertNewUser(userName, passwd)
	if err != nil {
		return err
	}
	return nil
}
