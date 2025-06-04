package logic

import (
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 展示会员权益
func UserLevelRights(in *user.UserLevelRightsRequest) (*user.UserLevelRightsResponse, error) {
	userInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	rightsInfo, err := dao_mysql.FindUserLevelRights(in.ULid)
	if err != nil {
		return nil, err
	}
	return &user.UserLevelRightsResponse{List: rightsInfo}, nil
}
