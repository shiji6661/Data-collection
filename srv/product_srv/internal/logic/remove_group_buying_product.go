package logic

import (

	"context"
	"errors"
	"go.uber.org/zap"
	"product_srv/dao/dao_mysql"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"
	"way/merchant"
)

// 删除拼团商品
func MerRemoveGProduct(in *product.RemoveGroupProductRequest) (*product.RemoveGroupProductResponse, error) {
	// 查看商家是否有权限删除拼团商品
	common, err := merchant.GetMerchantIdInfoCommon(in.MerId)
	if err != nil {
		return nil, err
	}
	if common.UserState != 1 {
		zap.L().Info("商家没有权限删除拼团商品！请联系管理员！")
		return nil, errors.New("商家没有权限删除拼团商品！请联系管理员！")
	}

	// 查看拼团商品是否存在
	byId, err := dao_mysql.FindGProductById(in.ProductId)
	if err != nil {
		return nil, err
	}

	if byId.ID == 0 {
		zap.L().Info("拼团商品不存在！！！")
		return nil, errors.New("拼团商品不存在！！！")
	}

	// 删除拼团商品
	_, err = dao_mysql.MerRemoveGProduct(in.ProductId)
	if err != nil {
		return nil, err
	}

	// 返回库存
	err = dao_redis.ReturnProductNum(context.Background(), in.ProductId, byId.Stock)
	if err != nil {
		return nil, err
	}

	return &product.RemoveGroupProductResponse{Message: "删除拼团商品成功！！！"}, nil
}
