package model_mysql

import (
	"common/global"
	"time"
)

type ArticleComment struct {
	Id        uint32    `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	ArticleId int32     `gorm:"column:article_id;type:int;comment:文章id;not null;" json:"article_id"`             // 文章id
	Username  string    `gorm:"column:username;type:varchar(255);comment:用户名;not null;" json:"username"`         // 用户名
	Content   string    `gorm:"column:content;type:varchar(255);comment:评论内容;not null;" json:"content"`          // 评论内容
	Pid       int32     `gorm:"column:pid;type:int;comment:评论父级id;default:0;" json:"pid"`                        // 评论父级id
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;" json:"created_at"`     // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);comment:修改时间;not null;" json:"updated_at"`     // 修改时间
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"` // 删除时间
}

func (am *ArticleComment) TableName() string {
	return "article_comment"
}
func (am *ArticleComment) CreateComment() error {
	return global.DB.Create(&am).Error
}
