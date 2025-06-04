package logic

import (
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// TODO:用户申请发票
func UserApplyInvoice(in *user.UserApplyInvoiceRequest) (*user.UserApplyInvoiceResponse, error) {
	invoice, s, err := dao_mysql.UserApplyInvoice(in)
	if err != nil {
		return nil, err
	}
	return &user.UserApplyInvoiceResponse{
		InvoiceId: int64(invoice.ID),
		PdfUrl:    s,
	}, nil
}
