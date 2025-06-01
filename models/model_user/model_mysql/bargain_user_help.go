package model_mysql

// todo:用户帮砍表
type BargainUserHelp struct {
	Id            int64   `gorm:"column:id;type:int UNSIGNED;comment:砍价用户帮助表ID;primaryKey;not null;" json:"id"`                    // 砍价用户帮助表ID
	Uid           int64   `gorm:"column:uid;type:int UNSIGNED;comment:帮助的用户id;default:NULL;" json:"uid"`                           // 帮助的用户id
	BargainId     int64   `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品ID;default:NULL;" json:"bargain_id"`              // 砍价商品ID
	BargainUserId int64   `gorm:"column:bargain_user_id;type:int UNSIGNED;comment:用户参与砍价表id;default:NULL;" json:"bargain_user_id"` // 用户参与砍价表id
	Price         float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:帮助砍价多少金额;default:NULL;" json:"price"`            // 帮助砍价多少金额
	CreatedAt     string  `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
}
