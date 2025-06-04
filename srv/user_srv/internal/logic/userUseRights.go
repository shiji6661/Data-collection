package logic

import (
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 用户使用权益展示
func UserUseRights(in *user.UserUseRightsRequest) (*user.UserUseRightsResponse, error) {
	userInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	list, err := dao_mysql.GetUserUseRightsList(in.Uid)
	if err != nil {
		return nil, err
	}
	return &user.UserUseRightsResponse{List: list}, nil
}
