package dao_mysql

import (
	"common/global"
	"errors"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"go.uber.org/zap"
	"models/model_user/model_mysql"
	"os"
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
	// 增加对输入参数的nil检查
	if invoice == nil {
		zap.L().Error("GenerateInvoicePDF: invoice 参数为 nil")
		return "", errors.New("invoice 参数为 nil")
	}
	if items == nil {
		zap.L().Error("GenerateInvoicePDF: items 参数为 nil")
		return "", errors.New("items 参数为 nil")
	}

	// 创建PDF对象
	pdf := gofpdf.New("P", "mm", "A4", "")
	// 检查PDF对象是否成功创建
	if pdf == nil {
		zap.L().Error("GenerateInvoicePDF: 创建 gofpdf 对象失败")
		return "", errors.New("创建 gofpdf 对象失败")
	}

	pdf.AddPage()
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: AddPage 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("AddPage 失败: %w", pdf.Error())
	}

	// 设置字体
	fontFileName := "simhei.ttf"
	// Corrected fontPath construction
	fontPath := "D:\\gocode\\src\\Data-collection\\public\\" + fontFileName
	fontAlias := "simhei"
	// 检查字体文件是否存在
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		zap.L().Error("GenerateInvoicePDF: 字体文件不存在", zap.String("fontPath", fontPath), zap.Error(err))
		return "", fmt.Errorf("字体文件不存在: %s", fontPath)
	}

	pdf.AddUTF8Font(fontAlias, "", fontPath)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: AddUTF8Font 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("AddUTF8Font 失败: %w", pdf.Error())
	}

	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont 失败: %w", pdf.Error())
	}

	// 设置边距
	pdf.SetMargins(20, 20, 20)
	pdf.SetAutoPageBreak(true, 20)

	// 公司信息
	pdf.SetFont(fontAlias, "", 16)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (公司信息标题) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (公司信息标题) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 10, "维康科技有限公司")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (公司名称) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (公司名称) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)

	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (公司地址) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (公司地址) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, "上海市浦东新区惠南镇盐大路1188号")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (公司地址) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (公司地址) 失败: %w", pdf.Error())
	}
	pdf.Ln(5)
	pdf.Cell(0, 8, "电话：010-12345678")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (电话) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (电话) 失败: %w", pdf.Error())
	}
	pdf.Ln(5)
	pdf.Cell(0, 8, "税号：123456789012345")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (税号) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (税号) 失败: %w", pdf.Error())
	}
	pdf.Ln(15)

	// 发票标题
	pdf.SetFont(fontAlias, "", 20)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (发票标题) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (发票标题) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(0, 15, "发票", "B", 1, "C", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (发票) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (发票) 失败: %w", pdf.Error())
	}
	pdf.Ln(10)

	// 发票信息
	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (发票信息) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (发票信息) 失败: %w", pdf.Error())
	}
	pdf.Cell(40, 8, "发票号码：")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (发票号码标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (发票号码标签) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, invoice.InvoiceNo)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (发票号码值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (发票号码值) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)

	pdf.Cell(40, 8, "发票日期：")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (发票日期标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (发票日期标签) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, invoice.IssueData.Format("2006-01-02"))
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (发票日期值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (发票日期值) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)

	pdf.Cell(40, 8, "有效期至：")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (有效期至标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (有效期至标签) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, invoice.Expire.Format("2006-01-02"))
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (有效期至值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (有效期至值) 失败: %w", pdf.Error())
	}
	pdf.Ln(15)

	// 客户信息
	pdf.SetFont(fontAlias, "", 14)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (客户信息标题) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (客户信息标题) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, "客户信息")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (客户信息) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (客户信息) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)

	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (客户抬头) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (客户抬头) 失败: %w", pdf.Error())
	}
	if invoice.TitleType == 1 {
		pdf.Cell(40, 8, "个人抬头：")
	} else {
		pdf.Cell(40, 8, "企业抬头：")
	}
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (抬头类型标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (抬头类型标签) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, invoice.Title)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (抬头值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (抬头值) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)

	pdf.Cell(40, 8, "纳税人识别号：")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (纳税人识别号标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (纳税人识别号标签) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, invoice.TaxId)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (纳税人识别号值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (纳税人识别号值) 失败: %w", pdf.Error())
	}
	pdf.Ln(15)

	// 商品表格表头
	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (商品表格表头) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (商品表格表头) 失败: %w", pdf.Error())
	}
	headers := []struct {
		label string
		width float64
	}{
		{"序号", 10},
		{"商品名称", 50},
		{"规格", 20},
		{"单位", 20},
		{"数量", 20},
		{"单价", 25},
		{"金额", 30},
	}
	for _, h := range headers {
		pdf.CellFormat(h.width, 8, h.label, "1", 0, "C", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品表头: "+h.label+") 失败", zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品表头: %s) 失败: %w", h.label, pdf.Error())
		}
	}
	pdf.Ln(-1) // 确保下一行在正确位置

	// 商品表格内容
	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (商品表格内容) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (商品表格内容) 失败: %w", pdf.Error())
	}
	for i, item := range items {
		pdf.CellFormat(10, 8, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品序号) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品序号 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(50, 8, item.ProductName, "1", 0, "L", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品名称) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品名称 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(20, 8, item.Spec, "1", 0, "C", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品规格) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品规格 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(20, 8, item.Unit, "1", 0, "C", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品单位) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品单位 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(20, 8, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品数量) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品数量 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(25, 8, fmt.Sprintf("%.2f", float64(item.Price)/100), "1", 0, "R", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品单价) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品单价 %d) 失败: %w", i+1, pdf.Error())
		}
		pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", item.Amount), "1", 1, "R", false, 0, "")
		if pdf.Err() {
			zap.L().Error("GenerateInvoicePDF: CellFormat (商品金额) 失败", zap.Int("itemIndex", i), zap.Error(pdf.Error()))
			return "", fmt.Errorf("CellFormat (商品金额 %d) 失败: %w", i+1, pdf.Error())
		}
	}

	// 合计
	pdf.CellFormat(60, 8, "合计金额", "1", 0, "R", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (合计金额标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (合计金额标签) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(60, 8, "", "1", 0, "C", false, 0, "") // 空白单元格，对应单价列
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (合计金额空单元格) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (合计金额空单元格) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(55, 8, fmt.Sprintf("%.2f", invoice.Amount), "1", 1, "R", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (合计金额值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (合计金额值) 失败: %w", pdf.Error())
	}

	// 税额
	pdf.CellFormat(60, 8, "税额", "1", 0, "R", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (税额标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (税额标签) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(60, 8, "13%", "1", 0, "C", false, 0, "") // 税率单元格
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (税率) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (税率) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(55, 8, fmt.Sprintf("%.2f", invoice.TaxAmount), "1", 1, "R", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (税额值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (税额值) 失败: %w", pdf.Error())
	}

	// 价税合计
	pdf.SetFont(fontAlias, "", 12)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (价税合计) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (价税合计) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(60, 8, "价税合计（大写）", "1", 0, "R", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (价税合计标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (价税合计标签) 失败: %w", pdf.Error())
	}
	pdf.CellFormat(115, 8, pkg.ConvertToChineseCurrency(invoice.Amount+invoice.TaxAmount), "1", 1, "L", false, 0, "")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: CellFormat (价税合计值) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("CellFormat (价税合计值) 失败: %w", pdf.Error())
	}

	// 底部信息
	pdf.Ln(20)
	pdf.SetFont(fontAlias, "", 10)
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: SetFont (备注) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("SetFont (备注) 失败: %w", pdf.Error())
	}
	pdf.Cell(0, 8, "备注：")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (备注标签) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (备注标签) 失败: %w", pdf.Error())
	}
	pdf.Ln(8)
	pdf.Cell(0, 8, "感谢您的惠顾！")
	if pdf.Err() {
		zap.L().Error("GenerateInvoicePDF: Cell (感谢惠顾) 失败", zap.Error(pdf.Error()))
		return "", fmt.Errorf("Cell (感谢惠顾) 失败: %w", pdf.Error())
	}

	// 保存PDF文件
	outputDir := "public/invoices"
	// 确保输出目录存在
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		zap.L().Info("PDF输出目录不存在，正在创建", zap.String("outputDir", outputDir))
		err = os.MkdirAll(outputDir, 0755)
		if err != nil {
			zap.L().Error("GenerateInvoicePDF: 创建PDF输出目录失败", zap.String("outputDir", outputDir), zap.Error(err))
			return "", fmt.Errorf("创建PDF输出目录失败: %w", err)
		}
		zap.L().Info("PDF输出目录创建成功", zap.String("outputDir", outputDir))
	}

	filePath := fmt.Sprintf("%s/invoice_%s.pdf", outputDir, invoice.InvoiceNo)
	zap.L().Info("尝试保存PDF文件", zap.String("filePath", filePath))
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		zap.L().Error("GenerateInvoicePDF: 生成PDF文件失败", zap.String("filePath", filePath), zap.Error(err))
		return "", fmt.Errorf("生成PDF文件失败: %w", err) // 使用 fmt.Errorf 包装错误以便追踪
	}

	zap.L().Info("PDF文件生成成功", zap.String("filePath", filePath))
	return filePath, nil
}
