package logic

import (
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
)

// TODO:商家添加优惠卷
func AddCoupon(in *product.AddCouponRequest) (*product.AddCouponResponse, error) {
	coupon, err := dao_mysql.AddCoupon(in)
	if err != nil {
		return nil, err
	}
	return &product.AddCouponResponse{CouponId: int64(coupon.ID)}, nil
}
