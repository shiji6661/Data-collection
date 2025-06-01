package model_mysql

import (
	"common/global"
	"time"
	"user_srv/proto_user/user"
)

// todo 用户等级表
type UserLevel struct {
	Id        int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);comment:会员名称;default:NULL;" json:"name"`            // 会员名称
	Score     int64     `gorm:"column:score;type:int;comment:会员分;default:NULL;" json:"score"`                    // 会员分
	Img       string    `gorm:"column:img;type:varchar(255);comment:会员图标;default:NULL;" json:"img"`              // 会员图标
	Explain   string    `gorm:"column:explain;type:varchar(255);comment:说明;default:NULL;" json:"explain"`        // 说明
	IsShow    int64     `gorm:"column:is_show;type:tinyint(1);comment:是否启用 1启用 0不启用;default:1;" json:"is_show"`  // 是否启用 1启用 0不启用
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:NULL;" json:"created_at"` // 添加时间
}

func (UserLevel) TableName() string {
	return "user_level"
}

func (l *UserLevel) FindLevelInfo() (ui []*user.LevelInfo, err error) {
	err = global.DB.Debug().Raw("SELECT user_level.`name`,user_level.score,user_level.img,user_level.`explain` FROM user_level").Scan(&ui).Error
	if err != nil {
		return nil, err
	}
	return ui, nil
}
