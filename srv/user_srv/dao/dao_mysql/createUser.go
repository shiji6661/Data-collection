package dao_mysql

import (
	"common/utils"
	"errors"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

func CreateUser(in *user.UserRegisterRequest) (*model_mysql.User, error) {
	users := &model_mysql.User{
		UserName:     in.UserName,
		UserPassword: utils.Md5(in.UserPassword),
		UserPhone:    in.UserPhone,
		UserState:    1,
		UserInviteId: in.UserInviteId,
	}
	err := users.CreateUser()

	if err != nil {
		return nil, errors.New("用户注册失败")
	}
	return users, nil
}
