package logic

import (
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// TODO:用户领取优惠卷
func UserReceiveCoupon(in *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
	coupon, err := dao_mysql.UserReceiveCoupon(in)
	if err != nil {
		return nil, err
	}
	return &user.UserReceiveCouponResponse{CUId: int64(coupon.ID)}, nil
}
