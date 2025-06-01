package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

type CouponUserUse struct {
	gorm.Model
	Cid         int64   `gorm:"column:cid;type:int UNSIGNED;comment:兑换的项目id;not null;" json:"cid"`                              // 兑换的项目id
	Uid         int64   `gorm:"column:uid;type:int UNSIGNED;comment:优惠券所属用户;not null;" json:"uid"`                              // 优惠券所属用户
	SCId        int64   `gorm:"type:int;not null;comment:'前台优惠券id;not null;'"`                                                  // 前台优惠券id
	CouponTitle string  `gorm:"column:coupon_title;type:varchar(32);comment:优惠券名称;not null;" json:"coupon_title"`               // 优惠券名称
	CouponPrice float64 `gorm:"column:coupon_price;type:decimal(8, 2) UNSIGNED;comment:优惠券的面值;not null;" json:"coupon_price"`   // 优惠券的面值
	UseTime     string  `gorm:"column:use_time;type:varchar(50);comment:使用时间;" json:"use_time"`                                 // 使用时间
	Status      int64   `gorm:"column:status;type:tinyint(1);comment:状态（0：未使用，1：已使用, 2:已过期）;not null;default:0;" json:"status"` // 状态（0：未使用，1：已使用, 2:已过期）
	IsFail      int64   `gorm:"column:is_fail;type:tinyint UNSIGNED;comment:是否有效;not null;default:0;" json:"is_fail"`           // 是否有效
}

// todo：领取优惠券
func (cu *CouponUserUse) ReceiveCoupon() error {
	return global.DB.Create(&cu).Error
}

// todo: 根据用户id 优惠卷id 查询是否领取过优惠券
func (cu *CouponUserUse) FindCouponByUidAndCid(uid, cid int64) error {
	return global.DB.Where("uid = ? and cid = ?", uid, cid).Limit(1).Find(&cu).Error
}
