package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
	"models/model_user/model_mysql"
)

type Commission struct {
	Id      uint32  `gorm:"column:id;type:int UNSIGNED;comment:分销id;primaryKey;not null;" json:"id"` // 分销id
	Uid     int32   `gorm:"column:uid;type:int;comment:用户id;not null;" json:"uid"`                   // 用户id
	Orderid int32   `gorm:"column:orderId;type:int;comment:订单id;not null;" json:"orderId"`           // 订单id
	Amount  float64 `gorm:"column:amount;type:decimal(10, 2);comment:佣金金额;not null;" json:"amount"`  // 佣金金额
	Level   int32   `gorm:"column:level;type:int;comment:佣金级别1一级2二级;not null;" json:"level"`         // 佣金级别1一级2二级
}

func (c *Commission) TableName() string {
	return "commission"
}

// 查询用户上级
func FindUserInviteId(uid int32) (int32, error) {
	var user model_mysql.User
	err := global.DB.Where("id = ?", uid).Find(&user).Error
	if err != nil {
		return 0, err
	}
	return int32(user.UserPid), nil
}

// 查询用户上上级
func FindUserInviteId2(uid int32) (int32, error) {
	var user model_mysql.User
	id, err := FindUserInviteId(uid)
	if err != nil {
		return 0, err
	}
	err = global.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return 0, err
	}
	return int32(user.UserPid), nil
}

// 给用户分配佣金
func AllocateCommissionToUser(uid, orderId, level int32, amount float64) error {
	err := global.DB.Model(&model_mysql.User{}).Where("id =?", uid).Update("sub_user_price", gorm.Expr("sub_user_price +?", amount)).Error
	if err != nil {
		return err
	}
	c := &Commission{
		Uid:     uid,
		Orderid: orderId,
		Amount:  amount,
		Level:   level,
	}
	err = global.DB.Create(c).Error
	if err != nil {
		return err
	}
	return nil
}

// 为上级和上上级发放佣金
func RewardCommissionForSuperiors(uid, orderId int32, orderAmount float64) error {
	// 定义上级和上上级的佣金比例
	firstLevelCommissionRate := 0.1
	secondLevelCommissionRate := 0.05

	// 查询上级
	firstLevelID, err := FindUserInviteId(uid)
	if err != nil {
		return err
	}
	if firstLevelID != 0 {
		// 计算上级的佣金
		firstLevelCommission := orderAmount * firstLevelCommissionRate
		err = AllocateCommissionToUser(firstLevelID, orderId, 1, firstLevelCommission)
		if err != nil {
			return err
		}
	}

	// 查询上上级
	secondLevelID, err := FindUserInviteId2(uid)
	if err != nil {
		return err
	}
	if secondLevelID != 0 {
		// 计算上上级的佣金
		secondLevelCommission := orderAmount * secondLevelCommissionRate
		err = AllocateCommissionToUser(secondLevelID, orderId, 2, secondLevelCommission)
		if err != nil {
			return err
		}
	}
	return nil
}
