package model_mysql

import "time"

// todo 用户使用权益表
type UserRights struct {
	Id        uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid       int32     `gorm:"column:uid;type:int;comment:用户id;default:NULL;" json:"uid"`                       // 用户id
	RightsId  int32     `gorm:"column:rights_id;type:int;comment:权益id;default:NULL;" json:"rights_id"`           // 权益id
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:NULL;" json:"created_at"` // 添加时间
}

func (UserRights) TableName() string {
	return "user_rights"
}
