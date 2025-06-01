package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
)

var U *model_mysql.User

// todo 注册前根据用户名 查找用户是否存在
func FindUserByUserName(username string) (*model_mysql.User, error) {
	users, err := U.FindUser(username)
	if err != nil {
		return nil, errors.New("用户查询失败")
	}

	return users, nil
}
