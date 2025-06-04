package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

const (
	IsHost      = 0  // 是否推荐 0 否 1 是
	ISShow      = 0  // 是否上架 0 否 1 是
	IsDel       = 0  // 是否删除 0 否 1 是
	Combination = 0  // 是否组合 0 否 1 是
	MerUse      = 0  // 是否商户使用 0 否 1 是
	ISPostage   = 0  // 是否包邮 0 否 1 是
	Postage     = 0  // 邮费
	Cost        = 0  // 拼团商品成本
	Num         = 2  // 单次购买数量
	Quota       = 10 // 限购总数
	QuotaShow   = 10 // 限购总数显示
)

// 查询拼团商品是否存在
func FindGProductById(productId int64) (gp *model_mysql.GroupBuyingProduct, err error) {
	gp = &model_mysql.GroupBuyingProduct{}
	err = gp.FindGroupProById(productId)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

// 创建拼团商品
func MerAddGroupBuyingProduct(in *product.AddGroupProductRequest) (gp *model_mysql.GroupBuyingProduct, err error) {
	gp = &model_mysql.GroupBuyingProduct{
		ProductId:   in.ProductId,
		MerId:       in.MerId,
		Title:       in.Title,
		Attr:        in.Attr,
		People:      in.People,
		Info:        in.Info,
		Price:       float64(in.Price),
		Sort:        in.Sort,
		Sales:       in.Sales,
		Stock:       in.Stock,
		IsHost:      IsHost,
		IsShow:      ISShow,
		IsDel:       IsDel,
		Combination: Combination,
		MerUse:      MerUse,
		IsPostage:   ISPostage,
		Postage:     Postage,
		StartTime:   in.StartTime,
		StopTime:    in.StopTime,
		Cost:        Cost,
		Browse:      Browse,
		Weight:      10.5,
		Volume:      13.3,
		Num:         Num,
		Quota:       Quota,
		QuotaShow:   QuotaShow,
	}
	err = gp.CreateGProduct()
	if err != nil {
		zap.L().Error("创建拼团商品失败", zap.Error(err))
		return nil, errors.New("创建拼团商品失败")
	}
	return gp, nil
}

// 删除拼团商品
func MerRemoveGProduct(productId int64) (gp *model_mysql.GroupBuyingProduct, err error) {
	gp = &model_mysql.GroupBuyingProduct{}
	err = gp.DeleteGroupProduct(productId)
	if err != nil {
		return nil, err
	}
	return gp, nil
}
