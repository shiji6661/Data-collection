package dao_mysql

import (
	"models/model_user/model_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

func FindMessage(in *user.FindMessageRequest) (s []*user.FindMessageInfo, err error) {
	message, err := model_mysql.FindMessage(int(in.ReceiverId))
	if err != nil {
		return nil, err
	}
	var so []*user.FindMessageInfo
	for _, m := range message {
		so = append(so, &user.FindMessageInfo{
			SendId:  int64(m.SendId),
			Context: m.Context,
		})
	}
	return so, err
}
