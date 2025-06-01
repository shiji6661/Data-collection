package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
)

// todo 用户手机号登录
func FindUserByPhone(userPhone string) (*model_mysql.User, error) {

	phone, err := U.FindUserByPhone(userPhone)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	return phone, nil
}
