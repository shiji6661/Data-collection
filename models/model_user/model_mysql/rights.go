package model_mysql

import (
	"time"
)

// todo 权益表
type Rights struct {
	Id          int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	RightsName  string    `gorm:"column:rights_name;type:varchar(255);comment:权益名称;default:NULL;" json:"rights_name"` // 权益名称
	UserLevelId int64     `gorm:"column:user_level_id;type:int;comment:会员等级;default:NULL;" json:"user_level_id"`      // 会员等级
	Img         string    `gorm:"column:img;type:varchar(255);comment:权益图标;default:NULL;" json:"img"`                 // 权益图标
	Explain     string    `gorm:"column:explain;type:varchar(255);comment:说明;default:NULL;" json:"explain"`           // 说明
	IsShow      int64     `gorm:"column:is_show;type:tinyint(1);comment:是否启用 1启用 0不启用;default:1;" json:"is_show"`     // 是否启用 1启用 0不启用
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:NULL;" json:"created_at"`    // 添加时间
}

func (Rights) TableName() string {
	return "rights"
}
