package logic

import (
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
)

// TODO: 商家删除优惠卷
func MerDelCouponStore(in *product.MerDeleteCouponStoreRequest) (*product.MerDeleteCouponStoreResponse, error) {
	_, err := dao_mysql.MerDeleteCoupon(in.Cid)
	if err != nil {
		return nil, err
	}
	return &product.MerDeleteCouponStoreResponse{Message: "删除成功！"}, nil
}
