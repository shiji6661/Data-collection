package logic

import (
	"errors"
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"time"
)

// 商家添加砍价商品
func CreateBargainProduct(in *product.CreateBargainProductRequest) (*product.CreateBargainProductResponse, error) {
	nowTime := time.Now().String()
	if in.StartTime > in.StopTime {
		return nil, errors.New("开始时间不能大于结束时间")
	}
	if in.StartTime < nowTime {
		return nil, errors.New("开始时间不能晚于当前时间")
	}
	bargainProduct, err := dao_mysql.CreateBargainProduct(in)
	if err != nil {
		return nil, errors.New("添加到数据库中失败")
	}
	return &product.CreateBargainProductResponse{Result: int64(bargainProduct.ID)}, err
}
