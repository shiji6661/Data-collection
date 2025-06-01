package model_mysql

import "common/global"

// todo:订单详情
type OrderInfo struct {
	Id        int64  `gorm:"column:id;type:int UNSIGNED;comment:唯一id;primaryKey;not null;" json:"id"` // 唯一id
	CreatedAt string `gorm:"column:created_at;type:datetime(3);default:NULL;" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Oid       int64  `gorm:"column:oid;type:int;comment:订单id;default:NULL;" json:"oid"`                           // 订单id
	CartId    int64  `gorm:"column:cart_id;type:int;comment:购物车id (id为0时商品没有经过购物车);default:NULL;" json:"cart_id"` // 购物车id (id为0时商品没有经过购物车)
	ProductId int64  `gorm:"column:product_id;type:int;comment:商品id;default:NULL;" json:"product_id"`             // 商品id
	CartInfo  string `gorm:"column:cart_info;type:text;comment:购买东西的详细信息;" json:"cart_info"`                      // 购买东西的详细信息
}

func (oi *OrderInfo) TableName() string {
	return "order_info"
}

func (oi *OrderInfo) FindOrderInfoById(id int64) (ori *OrderInfo, err error) {
	err = global.DB.Debug().Where("id = ?", id).Limit(1).Find(&ori).Error
	if err != nil {
		return nil, err
	}
	return ori, nil
}
