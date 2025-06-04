package dao_mysql

import (
	"common/global"
	"errors"
	"models/model_product/model_mysql"
	"models/model_product/model_redis"
	"product_srv/proto_product/product"
)

func AddSpikeProduct(in *product.AddSpikeProductRequest) (spike *model_mysql.SpikeProducts, err error) {
	s := model_mysql.Product{}
	id, err := s.FindProductById(in.ProductId)
	if err != nil {
		return nil, err
	}
	if id.ID == 0 {
		return nil, errors.New("商品不存在")
	}
	if id.Stock-in.SpikeStock < 0 {
		return nil, errors.New("商品库存不足")
	}
	ss := id.Stock - in.SpikeStock
	tx := global.DB.Begin()
	err = tx.Model(&model_mysql.Product{}).Where("id = ?", in.ProductId).Update("stock", ss).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("库存修改失败")
	}
	spike = &model_mysql.SpikeProducts{
		ProductId:    in.ProductId,
		ProductName:  id.StoreName,
		ProductPrice: id.Price,
		SpikePrice:   float64(in.SpikePrice),
		SpikeNumber:  in.SpikeStock,
		StartTime:    in.StartTime,
		EndTime:      in.EndTime,
	}
	err = tx.Create(&spike).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("添加失败")
	}
	err = model_redis.CreateSpike(uint(spike.ID), int(in.SpikeStock))
	if err != nil {
		return nil, errors.New("入redis队列失败")
	}
	tx.Commit()
	return spike, err
}
