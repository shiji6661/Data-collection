package logic

import (
	"errors"
	"go.uber.org/zap"
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
	"way/merchant"
)

// 商家删除商品
func MerDeletePro(in *product.MerchantDeleteProRequest) (*product.MerchantDeleteProResponse, error) {
	// 查看当前商家是否有权限
	common, err := merchant.GetMerchantIdInfoCommon(in.MerId)
	if err != nil {
		return nil, err
	}
	if common.ID == 0 {
		zap.L().Info("用户不存在")
		return nil, errors.New("用户不存在")
	}
	if common.UserState != 1 {
		zap.L().Info("用户已被禁用")
		return nil, errors.New("用户已被禁用")
	}

	// 查看商品是否存在
	id, err := dao_mysql.FindProductById(in.ProductId)
	if err != nil {
		return nil, err
	}
	if id.ID == 0 {
		zap.L().Info("商品不存在")
		return nil, errors.New("商品不存在")
	}
	// 查看商品是否属于当前商家
	if id.MerId != in.MerId {
		zap.L().Info("当前商品不属于当前商家！")
		return nil, errors.New("当前商品不属于当前商家！")
	}

	//删除
	_, err = dao_mysql.MerDeletePro(in.ProductId)
	if err != nil {
		return nil, err
	}
	return &product.MerchantDeleteProResponse{Message: "删除成功！"}, nil
}
