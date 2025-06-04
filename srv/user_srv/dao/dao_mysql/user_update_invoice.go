package dao_mysql

import (
	"errors"
	"go.uber.org/zap"
	"models/model_user/model_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:用户修改发票
func UserUpdateInvoice(in *user.UpdateInvoiceRequest) (i *model_mysql.Invoice, err error) {
	i = &model_mysql.Invoice{
		Type:      int(in.Type),
		TitleType: int(in.TitleType),
		Title:     in.Title,
		TaxId:     in.TaxId,
	}
	err = i.UpdateInvoice(in.InvoiceId, i)
	if err != nil {
		zap.L().Info("修改发票失败！")
		return nil, errors.New("修改发票失败！")
	}
	return i, nil
}
