package dao_mysql

import (
	"common/global"
	"errors"
	"models/model_product/model_mysql"
	"product_srv/pkg"
	"product_srv/proto_product/product"
)

// todo:添加砍价商品
func CreateBargainProduct(in *product.CreateBargainProductRequest) (result *model_mysql.Bargain, err error) {
	s := model_mysql.Product{}
	pro, err := s.FindProductById(in.ProductId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if pro.ID == 0 {
		return nil, errors.New("商品不存在")
	}
	if pro.Stock-in.Stock <= 0 {
		return nil, errors.New("库存不足")
	}
	num := pro.Stock - in.Stock
	tx := global.DB.Begin()
	err = tx.Model(&model_mysql.Product{}).Where("id = ?", in.ProductId).Update("stock", num).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("修改库存失败")
	}
	StartTime, err := pkg.ParseTime(in.StartTime)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("StartTime转为时间格式失败")
	}
	StopTime, err := pkg.ParseTime(in.StopTime)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("StopTime转为时间格式失败")
	}
	result = &model_mysql.Bargain{
		ProductId:       uint32(pro.ID),
		Title:           in.Title,
		Stock:           uint32(in.Stock),
		StartTime:       StartTime,
		StopTime:        StopTime,
		StoreName:       pro.StoreName,
		Price:           pro.Price,
		MinPrice:        float64(in.MinPrice),
		BargainMaxPrice: float64(in.BargainMaxPrice),
		BargainMinPrice: float64(in.BargainMinPrice),
	}

	err = tx.Create(&result).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("添加失败")
	}
	tx.Commit()
	return result, err
}
