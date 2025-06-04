package dao_mysql

import (
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

const (
	Sales     = 100 // 销量
	IsShow    = 1   // 状态（0：未上架，1：上架）
	IsHot     = 0   // 是否热卖
	IsBenefit = 0   // 是否特价
	IsBest    = 0   // 是否精品
	IsNew     = 0   // 是否新品
	IsPostage = 1   // 是否包邮
	IsSeckill = 0   // 秒杀状态 0 未开启 1已开启
	IsBargain = 0   // 砍价状态 0未开启 1开启
	IsGood    = 0   // 是否优品推荐
	Browse    = 200 // 浏览量
	Activity  = ""  // 活动显示排序1=秒杀，2=砍价，3=拼团
)

// 通过商品名称查询信息
func FindProductByStoreName(storeName string) (p *model_mysql.Product, err error) {
	p = &model_mysql.Product{}
	err = p.FindProductByStoreName(storeName)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 创建商品
func CreateProduct(in *product.AddProductRequest) (p *model_mysql.Product, err error) {
	p = &model_mysql.Product{
		MerId:     in.MerId,
		Image:     in.Image,
		StoreName: in.StoreName,
		StoreInfo: in.StoreInfo,
		Keyword:   in.Keyword,
		CateId:    in.CateId,
		Price:     float64(in.Price),
		VipPrice:  float64(in.VipPrice),
		OtPrice:   float64(in.OtPrice),
		Sales:     Sales,
		Stock:     in.Stock,
		IsShow:    IsShow,
		IsHot:     IsHot,
		IsBenefit: IsBenefit,
		IsBest:    IsBest,
		IsNew:     IsNew,
		IsPostage: IsPostage,
		Cost:      float64(in.Cost),
		IsSeckill: IsSeckill,
		IsBargain: IsBargain,
		IsGood:    IsGood,
		Browse:    Browse,
		Activity:  Activity,
	}
	err = p.AddProduct()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 更具商品id查询
func FindProductById(productId int64) (p *model_mysql.Product, err error) {
	id, err := p.FindProductById(productId)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// 删除商品
func MerDeletePro(productId int64) (p *model_mysql.Product, err error) {
	p = &model_mysql.Product{}
	err = p.DeleteProduct(productId)
	if err != nil {
		return nil, err
	}
	return p, nil
}
