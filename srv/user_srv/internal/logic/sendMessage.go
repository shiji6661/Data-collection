package logic

import (
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

func SendMessage(in *user.SendMessageRequest) (*user.SendMessageResponse, error) {
	_, err := dao_mysql.SendMessage(in)
	if err != nil {
		return nil, err
	}
	return &user.SendMessageResponse{Result: "发送成功"}, err
}
