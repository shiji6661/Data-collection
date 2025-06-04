package dao_mysql

import (
	"models/model_user/model_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

func SendMessage(in *user.SendMessageRequest) (s *model_mysql.Message, err error) {
	s = &model_mysql.Message{
		SendId:     int32(in.SendId),
		ReceiverId: int32(in.ReceiverId),
		Context:    in.Context,
	}
	err = s.UserSendToReceiver()
	if err != nil {
		return nil, err
	}
	return s, err
}
