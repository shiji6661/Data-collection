package model_mysql

import (
	"common/global"
	"time"
)

// todo 会员积分记录表
type MemberPointsRecord struct {
	Id         uint32    `gorm:"column:id;type:int(10) UNSIGNED;primaryKey;not null;" json:"id"`
	Uid        int32     `gorm:"column:uid;type:int(11);comment:用户id;not null;" json:"uid"`                         // 用户id
	Points     int32     `gorm:"column:points;type:int(11);comment:积分变动;not null;" json:"points"`                 // 积分变动
	Type       int8      `gorm:"column:type;type:tinyint(4);comment:类型(1:消费,2:邀请);not null;" json:"type"`       // 类型(1:消费,2:邀请)
	OrderId    string    `gorm:"column:order_id;type:varchar(50);comment:订单号;default:NULL;" json:"order_id"`       // 订单号
	Amount     float64   `gorm:"column:amount;type:decimal(10, 2);comment:消费金额;default:NULL;" json:"amount"`      // 消费金额
	InvitedUid int32     `gorm:"column:invited_uid;type:int(11);comment:被邀请人id;default:NULL;" json:"invited_uid"` // 被邀请人id
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`     // 创建时间
}

func (MemberPointsRecord) TableName() string {
	return "member_points_record"
}

func (r *MemberPointsRecord) GetPoints(uid int64) (mpr []*MemberPointsRecord, err error) {
	err = global.DB.Debug().Raw("SELECT * FROM member_points_record WHERE uid = ?", uid).Scan(&mpr).Error
	if err != nil {
		return nil, err
	}
	return mpr, nil
}
