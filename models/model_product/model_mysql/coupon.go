package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(30);not null;comment:'优惠卷名称'" json:"title"`                   // 优惠券名称
	Integral    int64     `gorm:"type:int;not null;comment:'兑换消耗积分值'" json:"integral"`                      // 兑换消耗积分值
	CouponPrice float64   `gorm:"type:decimal(10,2);not null;comment:'兑换的优惠券面值'" json:"coupon_price"`       // 兑换的优惠券面值
	UseMinPrice float64   `gorm:"type:decimal(10,2);not null;comment:'最低消费多少金额可用优惠券'" json:"use_min_price"` // 最低消费多少金额可用优惠券
	CouponTime  time.Time `gorm:"type:datetime;not null;comment:'优惠券有效期限（单位：天）'" json:"coupon_time"`        // 优惠券有效期限（单位：天）
	Status      int64     `gorm:"type:int;not null;comment:'状态（0：关闭，1：开启）'" json:"status"`                  // 状态（0：关闭，1：开启）
	IsDel       int64     `gorm:"type:int;not null;comment:'是否删除'" json:"is_del"`                           // 是否删除
	Type        int64     `gorm:"type:int;not null;comment:'优惠券类型 0-通用 1-品类券 2-商品券'" json:"type"`           // 优惠券类型 0-通用 1-品类券 2-商品券
}

// TODO:添加优惠卷
func (c *Coupon) AddCoupon() error {
	return global.DB.Create(&c).Error
}

// TODO:根据优惠卷名称查询优惠卷是否存在
func (c *Coupon) FindCouponByTitle(title string) error {
	return global.DB.Where("title = ?", title).Limit(1).Find(&c).Error
}

// TODO:根据优惠卷id查询优惠卷
func (c *Coupon) FindCouponById(cid int64) error {
	return global.DB.Where("id = ?", cid).Limit(1).Find(&c).Error
}
