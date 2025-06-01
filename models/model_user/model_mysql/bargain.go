package model_mysql

// todo:砍价表
type Bargain struct {
	Id              int64   `gorm:"column:id;type:int UNSIGNED;comment:砍价商品ID;primaryKey;not null;" json:"id"` // 砍价商品ID
	CreatedAt       string  `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt       string  `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt       string  `gorm:"column:deleted_at;type:datetime(3);comment:删除 ;default:NULL;" json:"deleted_at"`                                  // 删除
	ProductId       int64   `gorm:"column:product_id;type:int UNSIGNED;comment:关联商品ID;default:NULL;" json:"product_id"`                              // 关联商品ID
	Title           string  `gorm:"column:title;type:varchar(255);comment:砍价活动名称;default:NULL;" json:"title"`                                        // 砍价活动名称
	Image           string  `gorm:"column:image;type:varchar(150);comment:砍价活动图片;default:NULL;" json:"image"`                                        // 砍价活动图片
	Stock           int64   `gorm:"column:stock;type:int UNSIGNED;comment:库存;default:NULL;" json:"stock"`                                            // 库存
	Sales           int64   `gorm:"column:sales;type:int UNSIGNED;comment:销量;default:NULL;" json:"sales"`                                            // 销量
	Images          string  `gorm:"column:images;type:varchar(2000);comment:砍价商品轮播图;default:NULL;" json:"images"`                                    // 砍价商品轮播图
	StartTime       string  `gorm:"column:start_time;type:datetime;comment:砍价开启时间;default:NULL;" json:"start_time"`                                  // 砍价开启时间
	StopTime        string  `gorm:"column:stop_time;type:datetime;comment:砍价结束时间;default:NULL;" json:"stop_time"`                                    // 砍价结束时间
	StoreName       string  `gorm:"column:store_name;type:varchar(255);comment:砍价商品名称;default:NULL;" json:"store_name"`                              // 砍价商品名称
	Price           float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:砍价金额;default:NULL;" json:"price"`                                // 砍价金额
	MinPrice        float64 `gorm:"column:min_price;type:decimal(8, 2) UNSIGNED;comment:砍价商品最低价;default:NULL;" json:"min_price"`                     // 砍价商品最低价
	Num             int64   `gorm:"column:num;type:int UNSIGNED;comment:每次购买的砍价商品数量;default:NULL;" json:"num"`                                       // 每次购买的砍价商品数量
	BargainMaxPrice float64 `gorm:"column:bargain_max_price;type:decimal(8, 2) UNSIGNED;comment:用户每次砍价的最大金额;default:NULL;" json:"bargain_max_price"` // 用户每次砍价的最大金额
	BargainMinPrice float64 `gorm:"column:bargain_min_price;type:decimal(8, 2) UNSIGNED;comment:用户每次砍价的最小金额;default:NULL;" json:"bargain_min_price"` // 用户每次砍价的最小金额
	BargainNum      int64   `gorm:"column:bargain_num;type:int UNSIGNED;comment:用户每次砍价的次数;default:1;" json:"bargain_num"`                            // 用户每次砍价的次数
	Status          int64   `gorm:"column:status;type:tinyint UNSIGNED;comment:砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间);default:1;" json:"status"`         // 砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间)
	Info            string  `gorm:"column:info;type:varchar(255);comment:砍价活动简介;default:NULL;" json:"info"`                                          // 砍价活动简介
	Cost            float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;default:NULL;" json:"cost"`                                   // 成本价
	IsHot           int64   `gorm:"column:is_hot;type:tinyint UNSIGNED;comment:是否推荐0不推荐1推荐;default:0;" json:"is_hot"`                                // 是否推荐0不推荐1推荐
	AddTime         int64   `gorm:"column:add_time;type:int UNSIGNED;comment:添加时间;default:NULL;" json:"add_time"`                                    // 添加时间
	Look            int64   `gorm:"column:look;type:int UNSIGNED;comment:砍价商品浏览量;default:0;" json:"look"`                                            // 砍价商品浏览量
	Share           int64   `gorm:"column:share;type:int UNSIGNED;comment:砍价商品分享量;default:0;" json:"share"`                                          // 砍价商品分享量
	Weight          float64 `gorm:"column:weight;type:decimal(8, 2);comment:重量;default:0.00;" json:"weight"`                                         // 重量
	Volume          float64 `gorm:"column:volume;type:decimal(8, 2);comment:体积;default:0.00;" json:"volume"`                                         // 体积
	Quota           int64   `gorm:"column:quota;type:int;comment:限购总数;default:0;" json:"quota"`                                                      // 限购总数
	QuotaShow       int64   `gorm:"column:quota_show;type:int;comment:限量总数显示;default:0;" json:"quota_show"`                                          // 限量总数显示
}
