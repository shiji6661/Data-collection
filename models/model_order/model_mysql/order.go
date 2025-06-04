package model_mysql

import (
	"common/global"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:订单
type Order struct {
	Id                     int64   `gorm:"column:id;type:int UNSIGNED;comment:订单ID;primaryKey;not null;" json:"id"`                                  // 订单ID
	CreatedAt              string  `gorm:"column:created_at;type:datetime(3);comment:创建时间;not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 创建时间
	UpdatedAt              string  `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt              string  `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	OrderId                string  `gorm:"column:order_id;type:varchar(80);comment:订单号;default:NULL;" json:"order_id"`                                              // 订单号
	Uid                    int64   `gorm:"column:uid;type:int UNSIGNED;comment:用户id;default:NULL;" json:"uid"`                                                      // 用户id
	RealName               string  `gorm:"column:real_name;type:varchar(32);comment:用户姓名;default:NULL;" json:"real_name"`                                           // 用户姓名
	UserPhone              string  `gorm:"column:user_phone;type:varchar(18);comment:用户电话;default:NULL;" json:"user_phone"`                                         // 用户电话
	UserAddress            string  `gorm:"column:user_address;type:varchar(100);comment:详细地址;default:NULL;" json:"user_address"`                                    // 详细地址
	CartId                 string  `gorm:"column:cart_id;type:varchar(256);comment:购物车id;default:[];" json:"cart_id"`                                               // 购物车id
	FreightPrice           float64 `gorm:"column:freight_price;type:decimal(8, 2);comment:运费金额;default:0.00;" json:"freight_price"`                                 // 运费金额
	TotalNum               int64   `gorm:"column:total_num;type:int UNSIGNED;comment:订单商品总数;default:0;" json:"total_num"`                                           // 订单商品总数
	TotalPrice             float64 `gorm:"column:total_price;type:decimal(8, 2) UNSIGNED;comment:订单总价;default:0.00;" json:"total_price"`                            // 订单总价
	Paid                   int64   `gorm:"column:paid;type:tinyint UNSIGNED;comment:支付状态 0待支付 1已支付 2取消支付;not null;default:0;" json:"paid"`                          // 支付状态 0待支付 1已支付 2取消支付
	PayTime                string  `gorm:"column:pay_time;type:datetime(3);comment:支付时间;default:NULL;" json:"pay_time"`                                             // 支付时间
	PayType                int64   `gorm:"column:pay_type;type:tinyint(1);comment:支付方式 1支付宝 2微信 3银行卡;default:NULL;" json:"pay_type"`                                // 支付方式 1支付宝 2微信 3银行卡
	Status                 int64   `gorm:"column:status;type:tinyint(1);comment:订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）;default:0;" json:"status"` // 订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）
	RefundStatus           int64   `gorm:"column:refund_status;type:tinyint UNSIGNED;comment:0 未退款 1 申请中 2 已退款;default:0;" json:"refund_status"`                    // 0 未退款 1 申请中 2 已退款
	RefundReasonWapExplain string  `gorm:"column:refund_reason_wap_explain;type:varchar(255);comment:退款用户说明;default:NULL;" json:"refund_reason_wap_explain"`        // 退款用户说明
	RefundReasonTime       string  `gorm:"column:refund_reason_time;type:datetime(3);comment:退款时间;default:NULL;" json:"refund_reason_time"`                         // 退款时间
	RefundReasonWap        string  `gorm:"column:refund_reason_wap;type:varchar(255);comment:前台退款原因;default:NULL;" json:"refund_reason_wap"`                        // 前台退款原因
	RefundReason           string  `gorm:"column:refund_reason;type:varchar(255);comment:不退款的理由;default:NULL;" json:"refund_reason"`                                // 不退款的理由
	RefundPrice            float64 `gorm:"column:refund_price;type:decimal(8, 2) UNSIGNED;comment:退款金额;default:0.00;" json:"refund_price"`                          // 退款金额
	MerId                  int64   `gorm:"column:mer_id;type:int UNSIGNED;comment:商户ID;default:0;" json:"mer_id"`                                                   // 商户ID
	CombinationId          int64   `gorm:"column:combination_id;type:int UNSIGNED;comment:拼团商品id0一般商品;default:0;" json:"combination_id"`                            // 拼团商品id0一般商品
	PinkId                 int64   `gorm:"column:pink_id;type:int UNSIGNED;comment:拼团id 0没有拼团;default:0;" json:"pink_id"`                                           // 拼团id 0没有拼团
	Cost                   float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;default:NULL;" json:"cost"`                                           // 成本价
	SeckillId              int64   `gorm:"column:seckill_id;type:int UNSIGNED;comment:秒杀商品ID;default:0;" json:"seckill_id"`                                         // 秒杀商品ID
	BargainId              int64   `gorm:"column:bargain_id;type:int UNSIGNED;comment:砍价id;default:0;" json:"bargain_id"`                                           // 砍价id
	StoreId                int64   `gorm:"column:store_id;type:int;comment:门店id;default:0;" json:"store_id"`                                                        // 门店id
	ProductId              int64   `gorm:"column:product_id;type:int;comment:商品id;default:0;" json:"product_id"`                                                    // 商品id
}

func (o *Order) TableName() string {
	return "order"
}

func (o *Order) FindOrderById(oid int64) error {
	err := global.DB.Debug().Where("id = ?", oid).Limit(1).Find(&o).Error
	if err != nil {
		return err
	}
	return nil
}

// todo 用户查看所有订单
func (o *Order) GetOrder() (ord []*user.OrderList, err error) {
	err = global.DB.Debug().Raw("SELECT `order`.order_id,`order`.real_name,product.image,product.store_name,`order`.total_num,`order`.`status`,`order`.total_price,`order`.created_at,`order`.pay_time FROM `order` JOIN product ON `order`.product_id = product.id").Scan(&ord).Error
	if err != nil {
		return nil, err
	}
	return ord, nil
}

// 通过商品id查询
func (o *Order) FindOrderByProductId(productId int64) error {
	return global.DB.Where("product_id = ?", productId).Limit(1).Find(&o).Error
}

// todo 下单同时减少库存
func (o *Order) AddProduct(id int64, stock int64) error {

	tx := global.DB.Begin()
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄

	if err := tx.Create(&o).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model_mysql.Product{}).Where("id=?", id).Update("stock", stock).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (o *Order) UpdateProductPay(pay int64, id int64) error {
	err := global.DB.Model(&Order{}).Where("id = ?", id).Update("pay_type", pay).Error
	if err != nil {
		return err

	}
	return nil
}

func FindByOrderId(id int64) (order *Order, err error) {
	err = global.DB.Debug().Where("id = ?", id).Limit(1).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return
}

func (o *Order) FindOrderOrderSn(id string, status int64) error {
	err := global.DB.Model(&Order{}).Debug().Where("order_id=?", id).Update("paid", status).Error
	if err != nil {
		return err
	}
	return nil
}

// 商家统计销售额
func (o *Order) CountAmount(merId int) (amount float64, err error) {
	salesAmount := 0.0
	var orderList []*Order
	err = global.DB.Where("mer_id=?", merId).Find(&orderList).Error
	if err != nil {
		return 0, err
	}
	for _, v := range orderList {
		salesAmount += v.TotalPrice
	}
	return salesAmount, nil
}
