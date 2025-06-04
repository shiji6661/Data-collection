package logic

import (
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

func FindMessage(in *user.FindMessageRequest) (*user.FindMessageResponse, error) {
	message, err := dao_mysql.FindMessage(in)
	if err != nil {
		return nil, err
	}
	return &user.FindMessageResponse{
		List: message,
	}, err
}
