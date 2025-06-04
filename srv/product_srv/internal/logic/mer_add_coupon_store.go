package logic

import (
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// TODO: 商家添加优惠卷
func MerAddCouponStore(in *product.MerAddCouponStoreRequest) (*product.MerAddCouponStoreResponse, error) {
	store, err := dao_mysql.MerAddCouponStore(in)
	if err != nil {
		return nil, err
	}
	return &product.MerAddCouponStoreResponse{Result: int64(store.ID)}, nil
}
