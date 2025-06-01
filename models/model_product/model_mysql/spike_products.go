package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

type SpikeProducts struct {
	gorm.Model
	ProductId    int64   `gorm:"column:product_id;type:bigint;comment:''商品id'';not null;" json:"product_id"`                                // ''商品id''
	ProductName  string  `gorm:"column:product_name;type:varchar(20);comment:''商品名称'';not null;" json:"product_name"`                       // ''商品名称''
	ProductPrice float64 `gorm:"column:product_price;type:decimal(10, 2);comment:''商品原价格'';not null;" json:"product_price"`                 // ''商品原价格''
	SpikePrice   float64 `gorm:"column:spike_price;type:decimal(10, 2);comment:''商品秒杀价格'';not null;" json:"spike_price"`                    // ''商品秒杀价格''
	SpikeNumber  int64   `gorm:"column:spike_number;type:bigint;comment:''秒杀库存'';not null;" json:"spike_number"`                            // ''秒杀库存''
	StartTime    string  `gorm:"column:start_time;type:varchar(20);comment:''秒杀开始时间'';not null;" json:"start_time"`                         // ''秒杀开始时间''
	EndTime      string  `gorm:"column:end_time;type:varchar(20);comment:''秒杀结束时间'';not null;" json:"end_time"`                             // ''秒杀结束时间''
	SpikeStatus  int64   `gorm:"column:spike_status;type:bigint;comment:''秒杀状态1准备阶段2秒杀阶段3秒杀结算阶段'';not null;default:1;" json:"spike_status"` // ''秒杀状态1准备阶段2秒杀阶段3秒杀结算阶段''
}

func (s *SpikeProducts) TableName() string {
	return "spike_products"
}

func (s *SpikeProducts) AddSpikeProduct() error {
	return global.DB.Create(&s).Error
}

func (s *SpikeProducts) FindSpikeProductsById(id int64) (spike *SpikeProducts, err error) {
	err = global.DB.Where("id = ?", id).Limit(1).Find(&spike).Error
	return
}
func (s *SpikeProducts) UpdateSpikeStatus(id int, i int) error {
	return global.DB.Model(&s).Where("id=?", id).Update("spike_status", i).Error
}
