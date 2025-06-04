package logic

import (
	"errors"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 用户查看订单详情
func UserViewOrderInfo(in *user.UserViewOrderInfoRequest) (*user.UserViewOrderInfoResponse, error) {
	userInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	info, err := dao_mysql.FindOrderInfoById(in.OrderId)
	if err != nil {
		return nil, err
	}
	return &user.UserViewOrderInfoResponse{
		CartId:    info.CartId,
		ProductId: info.ProductId,
		CartInfo:  info.CartInfo,
	}, nil
}
