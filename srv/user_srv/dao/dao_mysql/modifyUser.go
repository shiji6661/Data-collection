package dao_mysql

import (
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

var A *model_mysql.Address

func ModifyUser(in *user.UserModifyRequest) (*model_mysql.Address, error) {
	modify, err := A.UpdateUser(in)
	if err != nil {
		return nil, err
	}
	return modify, nil
}
