package logic

import (
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 展示会员
func UserLevelInfo(in *user.UserLevelInfoRequest) (*user.UserLevelInfoResponse, error) {
	userInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	info, err := dao_mysql.FindLevelInfo()
	if err != nil {
		return nil, err
	}
	return &user.UserLevelInfoResponse{
		Info: info,
	}, nil
}
