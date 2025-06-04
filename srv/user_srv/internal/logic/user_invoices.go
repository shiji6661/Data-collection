package logic

import (
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:用户发票展示
func UserInvoices(in *user.UserInvoicesListRequest) (*user.UserInvoicesListResponse, error) {
	invoices, err := dao_mysql.UserInvoices(in.Userid)
	if err != nil {
		return nil, err
	}
	return &user.UserInvoicesListResponse{List: invoices}, nil
}
