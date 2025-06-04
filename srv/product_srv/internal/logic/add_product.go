package logic

import (
	ab "Data-collection/way/merchant"
	"context"
	"errors"
	"go.uber.org/zap"
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/dao/dao_redis"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// 添加商品
func AddProduct(in *product.AddProductRequest) (*product.AddProductResponse, error) {
	// 判断商家是否有权限添加商品
	common, err := ab.GetMerchantIdInfoCommon(in.MerId)
	if err != nil {
		return nil, err
	}
	if common.ID == 0 {
		zap.L().Error("商家不存在！请重新登陆！")
		return nil, errors.New("商家不存在！请重新登陆！")
	}

	if common.UserState != 1 {
		zap.L().Error("商家未通过审核！请重新登陆！")
		return nil, errors.New("商家未通过审核！请重新登陆！")
	}
	// 判断商品是否存在
	name, err := dao_mysql.FindProductByStoreName(in.StoreName)
	if err != nil {
		return nil, err
	}
	if name.ID != 0 {
		zap.L().Info("商品已存在！请勿重复添加！")
		return nil, errors.New("商品已存在！请勿重复添加！")
	}

	// 添加商品
	createProduct, err := dao_mysql.CreateProduct(in)
	if err != nil {
		return nil, err
	}

	// 同步库存
	err = dao_redis.SyncProductToRedis(context.Background(), int64(createProduct.ID), in.Stock)
	if err != nil {
		return nil, err
	}

	return &product.AddProductResponse{ProductId: int64(createProduct.ID)}, nil
}
