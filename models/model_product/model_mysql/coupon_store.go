package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

type CouponStore struct {
	gorm.Model
	MerId         int64   `gorm:"column:mer_id;type:int;comment:商家ID;default:NULL;" json:"mer_id"`                                          // 商家ID
	Cid           int64   `gorm:"column:cid;type:int;comment:优惠券ID;default:NULL;" json:"cid"`                                               // 优惠券ID
	StartTime     string  `gorm:"column:start_time;type:varchar(50);comment:优惠券领取开启时间;" json:"start_time"`                                  // 优惠券领取开启时间
	EndTime       string  `gorm:"column:end_time;type:varchar(50);comment:优惠券领取结束时间;" json:"end_time"`                                      // 优惠券领取结束时间
	TotalCount    int64   `gorm:"column:total_count;type:int;comment:优惠券领取数量;default:NULL;" json:"total_count"`                             // 优惠券领取数量
	RemainCount   int64   `gorm:"column:remain_count;type:int;comment:优惠券剩余领取数量;default:NULL;" json:"remain_count"`                         // 优惠券剩余领取数量
	IsPermanent   int64   `gorm:"column:is_permanent;type:tinyint(1);comment:是否无限张数;not null;default:0;" json:"is_permanent"`               // 是否无限张数
	Status        int64   `gorm:"column:status;type:tinyint(1);comment:1 正常 0 未开启 -1 已无效;not null;default:1;" json:"status"`                // 1 正常 0 未开启 -1 已无效
	FullReduction float64 `gorm:"column:full_reduction;type:decimal(8, 2);comment:消费满多少赠送优惠券;not null;default:0.00;" json:"full_reduction"` // 消费满多少赠送优惠券
}

// TODO：根据cid查询优惠卷是否存在
func (cs *CouponStore) FindCouponStoreByCid(cid int64) error {
	return global.DB.Where("cid = ?", cid).Limit(1).Find(&cs).Error
}

// TODO: 添加优惠卷
func (cs *CouponStore) AddCouponStore() error {
	return global.DB.Create(&cs).Error
}

// TODO:删除优惠卷
func (cs *CouponStore) DelCouponStore(cid int64) error {
	return global.DB.Where("cid =?", cid).Delete(&cs).Error
}
