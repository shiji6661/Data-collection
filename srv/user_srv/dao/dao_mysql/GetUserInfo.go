package dao_mysql

import (
	"errors"
	"user_srv/proto_user/user"
)

func GetUserInfo(id int64) (*user.UserInfoResponse, error) {
	userinfo, err := U.GetUserInfo(id)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	return userinfo, nil
}
