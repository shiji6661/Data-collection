package model_mysql

// todo:商品秒杀表
type SecondKillProduct struct {
	Id         int64   `gorm:"column:id;type:int UNSIGNED;comment:商品秒杀商品表id;primaryKey;not null;" json:"id"` // 商品秒杀商品表id
	CreatedAt  string  `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt  string  `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt  string  `gorm:"column:deleted_at;type:datetime(3);comment:删除;default:NULL;" json:"deleted_at"`         // 删除
	ProductId  int64   `gorm:"column:product_id;type:int UNSIGNED;comment:商品id;default:NULL;" json:"product_id"`      // 商品id
	MerchantId int64   `gorm:"column:merchant_id;type:int;comment:商户id;default:NULL;" json:"merchant_id"`             // 商户id
	Image      string  `gorm:"column:image;type:varchar(255);comment:推荐图;default:NULL;" json:"image"`                 // 推荐图
	Images     string  `gorm:"column:images;type:varchar(2000);comment:轮播图;default:NULL;" json:"images"`              // 轮播图
	Title      string  `gorm:"column:title;type:varchar(255);comment:活动标题;default:NULL;" json:"title"`                // 活动标题
	Info       string  `gorm:"column:info;type:varchar(255);comment:简介;default:NULL;" json:"info"`                    // 简介
	Price      float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:价格;default:NULL;" json:"price"`       // 价格
	Cost       float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本;default:0.00;" json:"cost"`          // 成本
	OtPrice    float64 `gorm:"column:ot_price;type:decimal(10, 2) UNSIGNED;comment:原价;default:NULL;" json:"ot_price"` // 原价
	Stock      int64   `gorm:"column:stock;type:int UNSIGNED;comment:库存;default:NULL;" json:"stock"`                  // 库存
	StartTime  string  `gorm:"column:start_time;type:datetime(3);comment:开始时间;default:NULL;" json:"start_time"`       // 开始时间
	StopTime   string  `gorm:"column:stop_time;type:datetime(3);comment:结束时间;default:NULL;" json:"stop_time"`         // 结束时间
	Status     int64   `gorm:"column:status;type:tinyint UNSIGNED;comment:商品状态;default:NULL;" json:"status"`          // 商品状态
	IsHot      int64   `gorm:"column:is_hot;type:tinyint UNSIGNED;comment:热门推荐;not null;default:0;" json:"is_hot"`    // 热门推荐
	Weight     float64 `gorm:"column:weight;type:decimal(8, 2);comment:商品重量;default:0.00;" json:"weight"`             // 商品重量
	Volume     float64 `gorm:"column:volume;type:decimal(8, 2);comment:商品体积;default:0.00;" json:"volume"`             // 商品体积
	Quota      int64   `gorm:"column:quota;type:int;comment:限购总数;default:0;" json:"quota"`                            // 限购总数
	QuotaShow  int64   `gorm:"column:quota_show;type:int;comment:限购总数显示;default:0;" json:"quota_show"`                // 限购总数显示
}

func (sp *SecondKillProduct) TableName() string {
	return "second_kill_product"
}
