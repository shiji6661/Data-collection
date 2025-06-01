package logic

import (
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

// todo 用户个人信息修改
func UserModify(in *user.UserModifyRequest) (*user.UserModifyResponse, error) {
	_, err := dao_mysql.FindUserModifyById(in.Id)
	if err != nil {
		return nil, err
	}
	_, err = dao_mysql.ModifyUser(in)
	if err != nil {
		return nil, err
	}
	return &user.UserModifyResponse{Success: true}, nil
}
