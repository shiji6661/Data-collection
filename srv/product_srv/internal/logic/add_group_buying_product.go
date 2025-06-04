package logic

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"product_srv/dao/dao_mysql"
	"product_srv/dao/dao_redis"
	"product_srv/proto_product/product"

	"way/merchant"
)

// 创建拼团商品
func CreateGroupBuyingProduct(in *product.AddGroupProductRequest) (*product.AddGroupProductResponse, error) {
	// 查看商家是否有权限添加拼团商品
	common, err := merchant.GetMerchantIdInfoCommon(in.MerId)
	if err != nil {
		return nil, err
	}
	if common.UserState != 1 {
		zap.L().Info("商家没有权限添加拼团商品！请联系管理员！")
		return nil, errors.New("商家没有权限添加拼团商品！请联系管理员！")
	}

	// 查看商品是否存在
	id, err := dao_mysql.FindProductById(in.ProductId)
	if err != nil {
		return nil, err
	}

	if id.ID == 0 {

		zap.L().Info("商品不存在！")
		return nil, errors.New("商品不存在！")
	}

	// 查看拼团商品是否存在
	byId, err := dao_mysql.FindGProductById(in.ProductId)
	if err != nil {
		return nil, err
	}
	if byId.ID != 0 {

		zap.L().Info("拼团商品已存在！请勿重复添加！")
		return nil, errors.New("拼团商品已存在！请勿重复添加！")
	}

	// 判断redis商品库存是否足够
	redis, err := dao_redis.GetFromRedis(context.Background(), in.ProductId)
	if err != nil {
		return nil, err
	}
	if redis < in.Stock {
		message := fmt.Sprintf("%s 库存不足！请重新选择！", id.StoreName)
		zap.L().Info(message)
		return nil, errors.New(message)
	}
	// 创建拼团商品
	//判断拼团时间
	if in.StartTime > in.StopTime {
		zap.L().Info("拼团开始时间不可大于结束时间！")
		return nil, errors.New("拼团开始时间不可大于结束时间！")
	}

	buyingProduct, err := dao_mysql.MerAddGroupBuyingProduct(in)
	if err != nil {
		return nil, err
	}

	// 减少redis库存
	err = dao_redis.ReduceProductNum(context.Background(), in.ProductId, in.Stock)
	if err != nil {
		return nil, err
	}

	return &product.AddGroupProductResponse{GProductId: int64(buyingProduct.ID)}, nil

}
