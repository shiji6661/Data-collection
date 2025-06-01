package model_mysql

import (
	"common/global"
	"time"
)

type ArticleContent struct {
	Id          uint32    `gorm:"column:id;type:int UNSIGNED;comment:文章内容id;primaryKey;not null;" json:"id"`                                    // 文章内容id
	Nid         uint32    `gorm:"column:nid;type:int UNSIGNED;comment:文章id;not null;" json:"nid"`                                               // 文章id
	Content     string    `gorm:"column:content;type:text;comment:文章内容;not null;" json:"content"`                                               // 文章内容
	CreatedTime time.Time `gorm:"column:created_time;type:datetime(3);comment:添加时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_time"` // 添加时间
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
}

func (an *ArticleContent) TableName() string {
	return "article_content"
}

// 文章内容添加
func (an *ArticleContent) CreateArticleContent() error {
	return global.DB.Create(&an).Error
}
