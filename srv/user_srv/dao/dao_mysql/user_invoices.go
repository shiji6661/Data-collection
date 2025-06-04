package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_user/model_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:用户的发票展示
func UserInvoices(userid int64) ([]*user.InvoiceItem, error) {
	i := &model_mysql.Invoice{}
	list, err := i.InvoiceList(int(userid))
	if err != nil {
		zap.L().Info("查询发票失败")
		return nil, errors.New("查询发票失败")
	}
	var invoicesList []*user.InvoiceItem
	for _, i3 := range list {
		layout := "2006-01-02 15:04:05"
		issueData := i3.IssueData.Format(layout)
		expire := i3.Expire.Format(layout)
		invoicesList = append(invoicesList, &user.InvoiceItem{
			InvoiceNo: i3.InvoiceNo,
			OrderId:   int64(i3.OrderId),
			Type:      int64(i3.Type),
			Status:    int64(i3.Status),
			TitleType: int64(i3.TitleType),
			Title:     i3.Title,
			TaxId:     i3.TaxId,
			Amount:    float32(i3.Amount),
			TaxAmount: float32(i3.TaxAmount),
			IssueData: issueData,
			Expire:    expire,
		})
	}
	return invoicesList, nil
}
