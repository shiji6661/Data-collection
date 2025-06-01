package model_mysql

import (
	"common/global"
)

// todo:拼团表
type GroupBuying struct {
	Id             int64   `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt      string  `gorm:"column:created_at;type:datetime(3);comment:开始时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 开始时间
	UpdatedAt      string  `gorm:"column:updated_at;type:datetime(3);comment:修改时间;default:NULL;" json:"updated_at"`                 // 修改时间
	DeletedAt      string  `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                 // 删除时间
	Uid            int64   `gorm:"column:uid;type:int UNSIGNED;comment:用户id;default:NULL;" json:"uid"`                              // 用户id
	TotalNum       int64   `gorm:"column:total_num;type:int UNSIGNED;comment:购买商品个数;default:NULL;" json:"total_num"`                // 购买商品个数
	TotalPrice     float64 `gorm:"column:total_price;type:decimal(10, 2) UNSIGNED;comment:购买总金额;default:NULL;" json:"total_price"`  // 购买总金额
	Cid            int64   `gorm:"column:cid;type:int UNSIGNED;comment:拼团商品id;default:NULL;" json:"cid"`                            // 拼团商品id
	Pid            int64   `gorm:"column:pid;type:int UNSIGNED;comment:商品id;default:NULL;" json:"pid"`                              // 商品id
	People         int64   `gorm:"column:people;type:int UNSIGNED;comment:拼图总人数;default:NULL;" json:"people"`                       // 拼图总人数
	Price          float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:拼团商品单价;default:NULL;" json:"price"`             // 拼团商品单价
	StopTime       string  `gorm:"column:stop_time;type:datetime(3);comment:拼团结束时间;default:NULL;" json:"stop_time"`                 // 拼团结束时间
	KId            int64   `gorm:"column:k_id;type:int UNSIGNED;comment:团长id 0为团长;default:0;" json:"k_id"`                          // 团长id 0为团长
	IsTpl          int64   `gorm:"column:is_tpl;type:tinyint UNSIGNED;comment:是否发送模板消息0未发送1已发送;default:0;" json:"is_tpl"`           // 是否发送模板消息0未发送1已发送
	IsRefund       int64   `gorm:"column:is_refund;type:tinyint UNSIGNED;comment:是否退款 0未退款 1已退款;default:0;" json:"is_refund"`       // 是否退款 0未退款 1已退款
	Status         int64   `gorm:"column:status;type:tinyint UNSIGNED;comment:状态1进行中2已完成3未完成;default:1;" json:"status"`             // 状态1进行中2已完成3未完成
	InvitationCode string  `gorm:"type:varchar(255);comment:'拼团邀请码'"`
}

func (gb *GroupBuying) TableName() string {
	return "group_buying"
}

// 拼团
func (gb *GroupBuying) CreateGroupBuying() error {
	return global.DB.Create(&gb).Error
}

// 根据拼团商品id查找拼团信息
func (gb *GroupBuying) FindGroupBuyingById(cid int64) error {
	return global.DB.Where("cid = ?", cid).Limit(1).Find(&gb).Error
}

// 查询用户是否已经拼过团
func (gb *GroupBuying) FindGroupBuyingByUidAndCid(uid, cid int64) error {
	return global.DB.Where("uid =? and cid =?", uid, cid).Limit(1).Find(&gb).Error
}
