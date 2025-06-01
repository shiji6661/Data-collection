package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
	"product_srv/proto_product/product"
)

// todo:商品表
type Product struct {
	gorm.Model
	MerId     int64   `gorm:"column:mer_id;type:int UNSIGNED;comment:商户Id(0为总后台管理员创建,不为0的时候是商户后台创建);default:0;" json:"mer_id"` // 商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image     string  `gorm:"column:image;type:varchar(256);comment:商品图片;default:NULL;" json:"image"`                          // 商品图片
	StoreName string  `gorm:"column:store_name;type:varchar(128);comment:商品名称;default:NULL;" json:"store_name"`                // 商品名称
	StoreInfo string  `gorm:"column:store_info;type:varchar(256);comment:商品简介;default:NULL;" json:"store_info"`                // 商品简介
	Keyword   string  `gorm:"column:keyword;type:varchar(256);comment:关键字;default:NULL;" json:"keyword"`                       // 关键字
	CateId    string  `gorm:"column:cate_id;type:varchar(64);comment:分类id;default:NULL;" json:"cate_id"`                       // 分类id
	Price     float64 `gorm:"column:price;type:decimal(8, 2) UNSIGNED;comment:商品价格;default:0.00;" json:"price"`                // 商品价格
	VipPrice  float64 `gorm:"column:vip_price;type:decimal(8, 2) UNSIGNED;comment:会员价格;default:0.00;" json:"vip_price"`        // 会员价格
	OtPrice   float64 `gorm:"column:ot_price;type:decimal(8, 2) UNSIGNED;comment:市场价;default:0.00;" json:"ot_price"`           // 市场价
	Sales     int64   `gorm:"column:sales;type:mediumint UNSIGNED;comment:销量;default:0;" json:"sales"`                         // 销量
	Stock     int64   `gorm:"column:stock;type:mediumint UNSIGNED;comment:库存;default:0;" json:"stock"`                         // 库存
	IsShow    int64   `gorm:"column:is_show;type:tinyint(1);comment:状态（0：未上架，1：上架）;default:1;" json:"is_show"`                 // 状态（0：未上架，1：上架）
	IsHot     int64   `gorm:"column:is_hot;type:tinyint(1);comment:是否热卖;default:0;" json:"is_hot"`                             // 是否热卖
	IsBenefit int64   `gorm:"column:is_benefit;type:tinyint(1);comment:是否特价 ;0不是 1是 ;default:0;" json:"is_benefit"`            // 是否优惠
	IsBest    int64   `gorm:"column:is_best;type:tinyint(1);comment:是否精品;default:0;" json:"is_best"`                           // 是否精品
	IsNew     int64   `gorm:"column:is_new;type:tinyint(1);comment:是否新品;default:0;" json:"is_new"`                             // 是否新品
	IsPostage int64   `gorm:"column:is_postage;type:tinyint UNSIGNED;comment:是否包邮;default:0;" json:"is_postage"`               // 是否包邮
	Cost      float64 `gorm:"column:cost;type:decimal(8, 2) UNSIGNED;comment:成本价;default:NULL;" json:"cost"`                   // 成本价
	IsSeckill int64   `gorm:"column:is_seckill;type:tinyint UNSIGNED;comment:秒杀状态 0 未开启 1已开启;default:0;" json:"is_seckill"`    // 秒杀状态 0 未开启 1已开启
	IsBargain int64   `gorm:"column:is_bargain;type:tinyint UNSIGNED;comment:砍价状态 0未开启 1开启;default:NULL;" json:"is_bargain"`   // 砍价状态 0未开启 1开启
	IsGood    int64   `gorm:"column:is_good;type:tinyint(1);comment:是否优品推荐;default:0;" json:"is_good"`                         // 是否优品推荐
	Browse    int64   `gorm:"column:browse;type:int;comment:浏览量;default:0;" json:"browse"`                                     // 浏览量
	Activity  string  `gorm:"column:activity;type:varchar(255);comment:活动显示排序1=秒杀，2=砍价，3=拼团;" json:"activity"`                 // 活动显示排序1=秒杀，2=砍价，3=拼团
}

func (p *Product) TableName() string {
	return "product"
}

// TODO: 根据商品id查询商品
func (p *Product) FindProductById(productId int64) (result *Product, err error) {
	err = global.DB.Where("id = ?", productId).Limit(1).Find(&result).Error
	return result, err
}

func UpdateProductStock(productId int64, stock int64) error {
	return global.DB.Model(&Product{}).Where("id = ?", productId).Update("stock", stock).Error
}

// todo:查询商品是否存在
func (p *Product) FindProductByStoreName(storeName string) error {
	return global.DB.Where("store_name = ?", storeName).Limit(1).Find(&p).Error
}

// todo 商品分类
func ProductCategory(str string) (product []*product.ProductList, err error) {
	err = global.DB.Raw("SELECT *  FROM product p  INNER JOIN product_sort s on p.cate_id=s.id where  s.id = ? ", str).Scan(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

// todo 商品关键字查询
func FindProductKeyword(key string) (product []*Product, err error) {
	err = global.DB.Take(&product).Where("keyword like ?", "%"+key+"%").Find(&product).Error
	if err != nil {
		return nil, err
	}
	return
}

// todo 商品推荐
func ProductRecommend() (product []*Product, err error) {
	if err = global.DB.Order("browse desc").Limit(3).Find(&product).Error; err != nil {
		return nil, err
	}
	return
}

// todo:商品详情
func FindProductInfo(name string) (product []*product.ProductList, err error) {

	//p.id,p.mer_id,p.store_name,p.store_info,p.cate_id,p.price,p.sales,p.is_show,p.cost,p.is_good,p.browse,s.id,s.pid,s.cate_name,s.sort
	err = global.DB.Raw("SELECT *  FROM product p  INNER JOIN product_sort s on p.cate_id=s.id where p.store_name = ? limit 1", name).Scan(&product).Error
	if err != nil {
		return nil, err
	}
	return
}

// todo:商品列表查询
func GetProduct() (product []*Product, err error) {
	if err = global.DB.Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// todo: 添加商品
func (p *Product) AddProduct() error {
	return global.DB.Create(&p).Error

}

// todo:删除商品
func (p *Product) DeleteProduct(productId int64) error {
	return global.DB.Where("id = ?", productId).Delete(&p).Error
}

// todo:商品列表展示
func (p *Product) ProductList() ([]*Product, error) {
	var products []*Product
	err := global.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
