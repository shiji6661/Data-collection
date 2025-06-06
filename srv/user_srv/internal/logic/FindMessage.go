package logic

import (
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
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
