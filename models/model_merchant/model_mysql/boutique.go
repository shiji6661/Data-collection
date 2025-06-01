package model_mysql

// todo:店铺表
type Boutique struct {
	Id           int64  `gorm:"column:id;type:int UNSIGNED;comment:商户申请ID;primaryKey;not null;" json:"id"`                       // 商户申请ID
	CreatedAt    string `gorm:"column:created_at;type:datetime(3);comment:添加时间;default:CURRENT_TIMESTAMP(3);" json:"created_at"` // 添加时间
	UpdatedAt    string `gorm:"column:updated_at;type:datetime(3);default:NULL;" json:"updated_at"`
	DeletedAt    string `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	MerId        int64  `gorm:"column:mer_id;type:int UNSIGNED;comment:商户Id;default:0;uniqueIndex" json:"mer_id"`    // 商户ID
	Province     string `gorm:"column:province;type:varchar(32);comment:商户所在省;" json:"province"`                     // 商户所在省
	City         string `gorm:"column:city;type:varchar(32);comment:商户所在市;" json:"city"`                             // 商户所在市
	District     string `gorm:"column:district;type:varchar(32);comment:商户所在区;" json:"district"`                     // 商户所在区
	Address      string `gorm:"column:address;type:varchar(256);comment:商户详细地址;" json:"address"`                     // 商户详细地址
	MerchantName string `gorm:"column:merchant_name;type:varchar(256);comment:商户名称;" json:"merchant_name"`           // 商户名称
	LinkTel      string `gorm:"column:link_tel;type:varchar(16);comment:商户电话;" json:"link_tel"`                      // 商户电话
	Charter      string `gorm:"column:charter;type:varchar(512);comment:商户证书;" json:"charter"`                       // 商户证书
	ApplyTime    int64  `gorm:"column:apply_time;type:int UNSIGNED;comment:审核时间;default:0;" json:"apply_time"`       // 审核时间
	SuccessTime  int64  `gorm:"column:success_time;type:int;comment:通过时间;default:NULL;" json:"success_time"`         // 通过时间
	FailMessage  string `gorm:"column:fail_message;type:varchar(256);comment:未通过原因;" json:"fail_message"`            // 未通过原因
	FailTime     int64  `gorm:"column:fail_time;type:int UNSIGNED;comment:未通过时间;default:0;" json:"fail_time"`        // 未通过时间
	Status       int64  `gorm:"column:status;type:tinyint(1);comment:-1 审核未通过 0未审核 1审核通过;default:0;" json:"status"`  // -1 审核未通过 0未审核 1审核通过
	IsLock       int64  `gorm:"column:is_lock;type:tinyint UNSIGNED;comment:0 = 开启 1= 关闭;default:0;" json:"is_lock"` // 0 = 开启 1= 关闭
	IsDel        int64  `gorm:"column:is_del;type:tinyint UNSIGNED;comment:是否删除;default:0;" json:"is_del"`           // 是否删除
}

func (b *Boutique) TableName() string {
	return "boutique"
}
