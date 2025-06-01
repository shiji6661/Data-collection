package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

// todo:拼团商品表
type GroupBuyingProduct struct {
	gorm.Model
	ProductId   int64   `gorm:"column:product_id;type:int UNSIGNED;comment:商品id;not null;" json:"product_id"`    // 商品id
	MerId       int64   `gorm:"column:mer_id;type:int UNSIGNED;comment:商户id;not null;default:0;" json:"mer_id"`  // 商户id
	Title       string  `gorm:"column:title;type:varchar(255);comment:活动标题;default:NULL;" json:"title"`          // 活动标题
	Attr        string  `gorm:"column:attr;type:varchar(255);comment:活动属性;default:NULL;" json:"attr"`            // 活动属性
	People      int64   `gorm:"column:people;type:int UNSIGNED;comment:参团人数;default:NULL;" json:"people"`        // 参团人数
	Info        string  `gorm:"column:info;type:varchar(255);comment:简介;default:NULL;" json:"info"`              // 简介
	Price       float64 `gorm:"column:price;type:decimal(10, 2) UNSIGNED;comment:价格;default:NULL;" json:"price"` // 价格
	Sort        int64   `gorm:"column:sort;type:int UNSIGNED;comment:排序;default:NULL;" json:"sort"`              // 排序
	Sales       int64   `gorm:"column:sales;type:int UNSIGNED;comment:销量;default:0;" json:"sales"`               // 销量
	Stock       int64   `gorm:"column:stock;type:int UNSIGNED;comment:库存;default:NULL;" json:"stock"`            // 库存
	IsHost      int64   `gorm:"column:is_host;type:tinyint UNSIGNED;comment:推荐;default:0;" json:"is_host"`       // 推荐
	IsShow      int64   `gorm:"column:is_show;type:tinyint UNSIGNED;comment:商品状态;default:NULL;" json:"is_show"`  // 商品状态
	IsDel       int64   `gorm:"column:is_del;type:tinyint UNSIGNED;default:0;" json:"is_del"`
	Combination int64   `gorm:"column:combination;type:tinyint UNSIGNED;default:1;" json:"combination"`
	MerUse      int64   `gorm:"column:mer_use;type:tinyint UNSIGNED;comment:商户是否可用1可用0不可用;default:NULL;" json:"mer_use"`  // 商户是否可用1可用0不可用
	IsPostage   int64   `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮1是0否;default:NULL;" json:"is_postage"` // 是否包邮1是0否
	Postage     float64 `gorm:"column:postage;type:decimal(10, 2) UNSIGNED;comment:邮费;default:NULL;" json:"postage"`      // 邮费
	StartTime   string  `gorm:"column:start_time;type:datetime;comment:拼团开始时间;default:NULL;" json:"start_time"`           // 拼团开始时间
	StopTime    string  `gorm:"column:stop_time;type:datetime;comment:拼团结束时间;default:NULL;" json:"stop_time"`             // 拼团结束时间
	Cost        int64   `gorm:"column:cost;type:int UNSIGNED;comment:拼团商品成本;default:0;" json:"cost"`                      // 拼团商品成本
	Browse      int64   `gorm:"column:browse;type:int;comment:浏览量;default:0;" json:"browse"`                              // 浏览量
	Weight      float64 `gorm:"column:weight;type:decimal(8, 2);comment:重量;default:0.00;" json:"weight"`                  // 重量
	Volume      float64 `gorm:"column:volume;type:decimal(8, 2);comment:体积;default:0.00;" json:"volume"`                  // 体积
	Num         int64   `gorm:"column:num;type:int;comment:单次购买数量;default:NULL;" json:"num"`                              // 单次购买数量
	Quota       int64   `gorm:"column:quota;type:int;comment:限购总数;default:0;" json:"quota"`                               // 限购总数
	QuotaShow   int64   `gorm:"column:quota_show;type:int;comment:限量总数显示;default:0;" json:"quota_show"`                   // 限量总数显示
}

func (gbp *GroupBuyingProduct) TableName() string {
	return "group_buying_product"
}

// 根据商品id查询拼团商品信息
func (gp *GroupBuyingProduct) FindGroupProById(productId int64) error {
	return global.DB.Where("product_id = ?", productId).Limit(1).Find(&gp).Error
}

// 创建拼团商品
func (gp *GroupBuyingProduct) CreateGProduct() error {
	return global.DB.Create(&gp).Error
}

// 删除拼团商品 软删除
func (gp *GroupBuyingProduct) DeleteGroupProduct(productId int64) error {
	return global.DB.Where("product_id =?", productId).Delete(&gp).Error
}

// 扣减库存
func (gp *GroupBuyingProduct) ReduceStock(productId int64, stock int64) error {
	return global.DB.Model(&gp).Where("product_id = ?", productId).Update("stock", gorm.Expr("stock - ?", stock)).Error
}
