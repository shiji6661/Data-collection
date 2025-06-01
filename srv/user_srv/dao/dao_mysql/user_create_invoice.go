package dao_mysql

import (
	"common/global"
	"errors"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"go.uber.org/zap"
	"models/model_user/model_mysql"
	"time"
	"user_srv/pkg"
	"user_srv/proto_user/user"
	"way/order"
	"way/product"
)

// TODO:生成发票号码根据数据库id
func GetInvoiceNo(id int64) string {
	now := time.Now()
	Year := now.Year()
	Month := now.Month()
	return fmt.Sprintf("INV%d%02d%06d", Year, Month, id)
}

// TODO:税额计算
func GetTaxAmount(amount float64) float64 {
	return amount * 0.13
}

// TODO:用户申请发票
func UserApplyInvoice(in *user.UserApplyInvoiceRequest) (u *model_mysql.Invoice, pdfPath string, err error) {
	// 查询订单是否存在
	id, err := order.FindOrderById(in.OrderId)
	if err != nil {
		return nil, "", err
	}
	if id.Id == 0 {
		zap.L().Info("不存在该订单")
		return nil, "", errors.New("不存在该订单")
	}

	// 加上事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	if err != nil {
		return nil, "", err
	}

	taxAmount := GetTaxAmount(id.TotalPrice)

	u = &model_mysql.Invoice{
		OrderId:   int(in.OrderId),
		Userid:    int(in.Userid),
		TitleType: int(in.TitleType),
		Title:     in.Title,
		TaxId:     in.TaxId,
		Amount:    id.TotalPrice,
		TaxAmount: taxAmount,
		IssueData: time.Now(),
		Expire:    time.Now().AddDate(0, 6, 0),
	}

	// 创建发票
	err = tx.Create(u).Error
	if err != nil {
		tx.Rollback()
		zap.L().Info("申请发票失败！", zap.Error(err))
		return nil, "", errors.New("申请发票失败！")
	}

	byId, err := product.GetCommonFindProductById(id.ProductId)
	if err != nil {
		tx.Rollback()
		return nil, "", err
	}

	// 发票详情
	it := &model_mysql.InvoiceItem{
		InvoiceId:   int64(u.ID),
		OrderItemId: id.ProductId,
		ProductName: byId.StoreName,
		Spec:        "2kg",
		Unit:        "公斤",
		Quantity:    id.TotalNum,
		Price:       int64(byId.Price),
		Amount:      id.TotalPrice,
		TaxRate:     0.13,
		TaxAmount:   taxAmount, // 修正为计算的税额
	}

	// 创建发票详情
	err = tx.Create(it).Error
	if err != nil {
		tx.Rollback()
		zap.L().Info("创建发票详情失败！", zap.Error(err))
		return nil, "", errors.New("创建发票详情失败！")
	}

	// 获取发票号码
	No := GetInvoiceNo(int64(u.ID))

	// 更新发票号码
	u.InvoiceNo = No
	err = tx.Save(u).Error
	if err != nil {
		tx.Rollback()
		zap.L().Info("更新发票号码失败！", zap.Error(err))
		return nil, "", errors.New("更新发票号码失败！")
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		zap.L().Info("提交事务失败！", zap.Error(err))
		return nil, "", errors.New("提交事务失败！")
	}

	// 生成PDF
	pdfPath, err = GenerateInvoicePDF(u, []*model_mysql.InvoiceItem{it})
	if err != nil {
		zap.L().Info("生成PDF失败！", zap.Error(err))
		// 这里可以选择是否回滚整个操作，根据业务需求决定
		return u, "", errors.New("生成PDF失败！")
	}

	return u, pdfPath, nil
}

// GenerateInvoicePDF 生成发票PDF
func GenerateInvoicePDF(invoice *model_mysql.Invoice, items []*model_mysql.InvoiceItem) (string, error) {
	// 创建PDF对象
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// 设置边距
	pdf.SetMargins(20, 20, 20)
	pdf.SetAutoPageBreak(true, 20)

	// 公司信息
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "公司名称")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 8, "公司地址")
	pdf.Ln(5)
	pdf.Cell(0, 8, "电话：010-12345678")
	pdf.Ln(5)
	pdf.Cell(0, 8, "税号：123456789012345")
	pdf.Ln(15)

	// 发票标题
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 15, "发票", "B", 1, "C", false, 0, "")
	pdf.Ln(10)

	// 发票信息
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 8, "发票号码：")
	pdf.Cell(0, 8, invoice.InvoiceNo)
	pdf.Ln(8)

	pdf.Cell(40, 8, "发票日期：")
	pdf.Cell(0, 8, invoice.IssueData.Format("2006-01-02"))
	pdf.Ln(8)

	pdf.Cell(40, 8, "有效期至：")
	pdf.Cell(0, 8, invoice.Expire.Format("2006-01-02"))
	pdf.Ln(15)

	// 客户信息
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 8, "客户信息")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	if invoice.TitleType == 1 {
		pdf.Cell(40, 8, "个人抬头：")
	} else {
		pdf.Cell(40, 8, "企业抬头：")
	}
	pdf.Cell(0, 8, invoice.Title)
	pdf.Ln(8)

	pdf.Cell(40, 8, "纳税人识别号：")
	pdf.Cell(0, 8, invoice.TaxId)
	pdf.Ln(15)

	// 商品表格
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(10, 8, "序号", "1", 0, "C", false, 0, "")
	pdf.CellFormat(70, 8, "商品名称", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "规格", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "单位", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "数量", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 8, "单价", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "金额", "1", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	for i, item := range items {
		pdf.CellFormat(10, 8, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(70, 8, item.ProductName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 8, item.Spec, "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, item.Unit, "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 8, fmt.Sprintf("%.2f", float64(item.Price)/100), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", item.Amount), "1", 1, "R", false, 0, "")
	}

	// 合计
	pdf.CellFormat(140, 8, "合计金额", "1", 0, "R", false, 0, "")
	pdf.CellFormat(25, 8, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", invoice.Amount), "1", 1, "R", false, 0, "")

	// 税额
	pdf.CellFormat(140, 8, "税额", "1", 0, "R", false, 0, "")
	pdf.CellFormat(25, 8, "13%", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", invoice.TaxAmount), "1", 1, "R", false, 0, "")

	// 价税合计
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(140, 8, "价税合计（大写）", "1", 0, "R", false, 0, "")
	pdf.CellFormat(55, 8, pkg.ConvertToChineseCurrency(invoice.Amount+invoice.TaxAmount), "1", 1, "L", false, 0, "")

	// 底部信息
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 8, "备注：")
	pdf.Ln(8)
	pdf.Cell(0, 8, "感谢您的惠顾！")

	// 保存PDF文件
	filePath := fmt.Sprintf("invoices/invoice_%s.pdf", invoice.InvoiceNo)
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
