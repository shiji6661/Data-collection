package model_mysql

type LimitedTimeProduct struct {
	Id           int64   `gorm:"column:id;type:int UNSIGNED;comment:限时商品id;primaryKey;not null;" json:"id"`                       // 限时商品id
	CreatedAt    string  `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
	UpdateAt     string  `gorm:"column:update_at;type:datetime(3);comment:修改时间;default:CURRENT_TIMESTAMP(3);" json:"update_at"`   // 修改时间
	DeletedAt    string  `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`                 // 删除时间
	StartTime    string  `gorm:"column:start_time;type:datetime(3);comment:开始时间;default:NULL;" json:"start_time"`                 // 开始时间
	EndTime      string  `gorm:"column:end_time;type:datetime(3);comment:结束时间;default:NULL;" json:"end_time"`                     // 结束时间
	ProductId    int64   `gorm:"column:product_id;type:int;comment:商品id;default:NULL;" json:"product_id"`                         // 商品id
	MerId        int64   `gorm:"column:mer_id;type:int;comment:商户id;default:NULL;" json:"mer_id"`                                 // 商户id
	OldPrice     float64 `gorm:"column:old_price;type:decimal(10, 2);comment:商品原价;not null;" json:"old_price"`                    // 商品原价
	LimitedPrice float64 `gorm:"column:limited_price;type:decimal(10, 2);comment:限时抢购价格;not null;" json:"limited_price"`          // 限时抢购价格
	Stock        int64   `gorm:"column:stock;type:int;comment:限时抢购的商品库存数量;not null;" json:"stock"`                                // 限时抢购的商品库存数量
	SoulCount    int64   `gorm:"column:soul_count;type:int;comment:已经售出的商品数量 初始为0;not null;default:0;" json:"soul_count"`         // 已经售出的商品数量 初始为0
	Status       int64   `gorm:"column:status;type:tinyint(1);comment:''0 未开始'',''1 进行中'',''2 已结束'';not null;" json:"status"`     // ''0 未开始'',''1 进行中'',''2 已结束''
	Image        string  `gorm:"column:image;type:varchar(255);comment:商品图片;default:NULL;" json:"image"`                          // 商品图片
	Images       string  `gorm:"column:images;type:varchar(255);comment:轮播图;default:NULL;" json:"images"`                         // 轮播图
	Title        string  `gorm:"column:title;type:varchar(255);comment:活动标题;default:NULL;" json:"title"`                          // 活动标题
	Info         string  `gorm:"column:info;type:varchar(255);comment:简介;default:NULL;" json:"info"`                              // 简介
	Weight       float64 `gorm:"column:weight;type:decimal(10, 2);comment:商品重量;not null;" json:"weight"`                          // 商品重量
	Volume       float64 `gorm:"column:volume;type:decimal(10, 2);comment:商品体积;not null;" json:"volume"`                          // 商品体积
}

func (l *LimitedTimeProduct) TableName() string {
	return "limited_time_product"
}
