package logic

import (
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 用户查看所有订单
func UserViewOrder(in *user.UserViewOrderRequest) (*user.UserViewOrderResponse, error) {
	userInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	list, err := dao_mysql.FindUserList()
	if err != nil {
		return nil, err
	}
	return &user.UserViewOrderResponse{List: list}, nil
}
