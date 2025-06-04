package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"

	"time"
)

const (
	CStatus = 0 // 状态（0：关闭，1：开启）
	CIsDel  = 0 // 是否删除 0 未删除 1已删除
	Type    = 0 // 类型（0：满减，1：折扣，2：立减）
)

// TODO:商家添加优惠卷
func FindCouponByTitle(title string) (c *model_mysql.Coupon, err error) {
	c = &model_mysql.Coupon{}
	err = c.FindCouponByTitle(title)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// TODO: 查询优惠卷
func FindCouponById(id int64) (c *model_mysql.Coupon, err error) {
	c = &model_mysql.Coupon{}
	err = c.FindCouponById(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// TODO:商家添加优惠卷
func AddCoupon(in *product.AddCouponRequest) (c *model_mysql.Coupon, err error) {
	title, err := FindCouponByTitle(in.Title)
	if err != nil {
		return nil, err
	}
	if title.ID != 0 {
		zap.L().Info("优惠卷已存在！")
		return nil, err
	}
	c = &model_mysql.Coupon{
		Title:       in.Title,
		Integral:    in.Integral,
		CouponPrice: float64(in.CouponPrice),
		UseMinPrice: float64(in.UseMinPrice),
		CouponTime:  time.Now().AddDate(0, 12, 0),
		Status:      CStatus,
		IsDel:       CIsDel,
		Type:        Type,
	}
	err = c.AddCoupon()
	if err != nil {
		zap.L().Info("优惠卷添加失败！")
		return nil, errors.New("优惠卷添加失败！")
	}
	return c, nil
}
