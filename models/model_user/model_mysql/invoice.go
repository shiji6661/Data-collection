package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	InvoiceNo string    `gorm:"type:varchar(50);not null;comment:'发票号码'"`
	OrderId   int       `gorm:"type:int;not null;comment:'订单id'"`
	Userid    int       `gorm:"type:int;not null;comment:'用户id'"`
	Type      int       `gorm:"type:int;not null;default:1;comment:'发票类型：1-增值税普通发票，2-增值税专用发票，3-电子发票，4-纸质发票'"`
	Status    int       `gorm:"type:int;not null;default:1;comment:'发票状态：1-待开具 2-已开具 3-已作废 4-已红冲'"`
	TitleType int       `gorm:"type:int;not null;default:1;comment:'抬头类型： 1-个人 2-企业'"`
	Title     string    `gorm:"type:varchar(255);not null;comment:'抬头名称'"`
	TaxId     string    `gorm:"type:varchar(50);not null;comment:'纳税人识别号'"`
	Amount    float64   `gorm:"type:decimal(10,2);not null;comment:'发票金额'"`
	TaxAmount float64   `gorm:"type:decimal(10,2);not null;comment:'发票税额'"`
	IssueData time.Time `gorm:"type:datetime;not null;comment:'开具日期'"`
	Expire    time.Time `gorm:"type:datetime;not null;comment:'发票有效期'"`
}

// todo:用户申请发票
func (i *Invoice) CreateInvoice() error {
	return global.DB.Create(&i).Error
}

// todo:更新发票号码
func (i *Invoice) UpdateInvoiceNo(id int64, no string) error {
	return global.DB.Model(&i).Where("id = ?", id).Update("invoice_no", no).Error
}

// todo:根据用户id查看发票
func (i *Invoice) InvoiceList(userid int) ([]*Invoice, error) {
	var invoices []*Invoice
	err := global.DB.Where("userid = ?", userid).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

// todo:用户修改发票
func (i *Invoice) UpdateInvoice(id int64, data *Invoice) error {
	return global.DB.Model(&i).Where("id =?", id).Updates(data).Error
}

// todo:根据发票id查询
func (i *Invoice) FindInvoiceById(id int64) error {
	return global.DB.Where("id =?", id).Limit(1).Find(&i).Error
}
