package logic

import (
	"common/pkg/pkg_user"
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 实名认证
func RealName(in *user.RealNameRequest) (*user.RealNameResponse, error) {

	realInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if realInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	err = dao_mysql.FindRealName(in.Card)
	if err != nil {
		return nil, err
	}
	name := pkg_user.RealName(in.Name, in.Card)
	if !name {
		return nil, errors.New("实名认证失败")
	}
	reals, err := dao_mysql.CreateRealName(in.Name, in.Card, in.Uid)
	if err != nil {
		return nil, err
	}
	return &user.RealNameResponse{Id: reals.Id}, nil
}
