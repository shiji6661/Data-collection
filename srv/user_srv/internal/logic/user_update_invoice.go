package logic

import (
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

// todo: 用户修改发票
func UserUpdateInvoice(in *user.UpdateInvoiceRequest) (*user.UpdateInvoiceResponse, error) {
	invoice, err := dao_mysql.UserUpdateInvoice(in)
	if err != nil {
		return nil, err
	}
	return &user.UpdateInvoiceResponse{
		Type:      int64(invoice.Type),
		TitleType: int64(invoice.TitleType),
		Title:     invoice.Title,
		TaxId:     invoice.TaxId,
	}, nil
}
