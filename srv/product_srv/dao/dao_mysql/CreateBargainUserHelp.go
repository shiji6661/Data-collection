package dao_mysql

import (
	"errors"
	"math/rand"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:用户帮助砍价
func CreateBargainUserHelp(in *product.CreateBargainUserHelpRequest) (result *model_mysql.BargainUserHelp, err error) {
	s := model_mysql.Bargain{}

	bar, err := s.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if bar.ID == 0 {
		return nil, errors.New("砍价商品不存在")
	}

	b := model_mysql.BargainUser{}

	uid, err := b.FindBargainUserId(int(in.BargainUserId)) //todo:查询用户发起砍价表
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if uid.Id == 0 {
		return nil, errors.New("用户未发起砍价")
	}
	if uid.Uid == in.Uid {
		return nil, errors.New("不能帮自己砍价")
	}
	if uid.Price == uid.BargainPriceMin {
		return nil, errors.New("已是最低价！！！")
	}

	price := rand.Intn(int(bar.BargainMaxPrice-bar.BargainMinPrice)+1) + int(bar.BargainMinPrice)

	result = &model_mysql.BargainUserHelp{
		Uid:           in.Uid,
		BargainId:     in.BargainId,
		BargainUserId: in.BargainUserId,
		Price:         float64(price),
	}
	err = result.CreateBargainUserHelp()
	if err != nil {
		return nil, errors.New("帮助砍价失败")
	}
	if uid.Price-result.Price < uid.BargainPriceMin {
		err = uid.UpdateBargainPrice(int64(in.BargainUserId), uid.BargainPriceMin)
		if err != nil {
			return nil, errors.New("修改价格失败")
		}
		result.Price = uid.Price - uid.BargainPriceMin
	} else {
		err = uid.UpdateBargainPrice(int64(in.BargainUserId), uid.Price-result.Price)
		if err != nil {
			return nil, errors.New("修改价格失败")
		}
	}
	return result, err
}
