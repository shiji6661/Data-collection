package model_mysql

// todo:用户砍价表
type BargainUser struct {
	Id              int64   `gorm:"column:id;type:int UNSIGNED;comment:用户参与砍价表ID;primaryKey;not null;" json:"id"`                               // 用户参与砍价表ID
	Uid             int64   `gorm:"column:uid;type:int UNSIGNED;comment:用户ID;default:NULL;" json:"uid"`                                         // 用户ID
	BargainId       int64   `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价商品id;default:NULL;" json:"bargain_id"`                         // 砍价商品id
	BargainPriceMin float64 `gorm:"column:bargain_price_min;type:decimal(8, 2) UNSIGNED;comment:砍价的最低价;default:NULL;" json:"bargain_price_min"` // 砍价的最低价
	BargainPrice    float64 `gorm:"column:bargain_price;type:decimal(8, 2);comment:砍价金额;default:NULL;" json:"bargain_price"`                    // 砍价金额
	Price           float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:砍掉的价格;default:NULL;" json:"price"`                          // 砍掉的价格
	Status          int64   `gorm:"column:status;type:tinyint UNSIGNED;comment:状态 1参与中 2 活动结束参与失败 3活动结束参与成功;default:0;" json:"status"`          // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	CreatedAt       string  `gorm:"column:created_at;type:datetime(3);comment:参与时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"`            // 参与时间
	DeletedAt       string  `gorm:"column:deleted_at;type:datetime(3);comment:是否取消;default:NULL;" json:"deleted_at"`                            // 是否取消
}
