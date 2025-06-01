package dao_mysql

import (
	"models/model_order/model_mysql"
	"user_srv/proto_user/user"
)

var O *model_mysql.Order

// todo 用户查看所有订单
func FindUserList() (newList []*user.OrderList, err error) {
	getOrder, err := O.GetOrder()
	if err != nil {
		return nil, err
	}

	var news []*user.OrderList
	for _, i := range getOrder {
		news = append(news, &user.OrderList{
			OrderId:    i.OrderId,
			Image:      i.Image,
			StoreName:  i.StoreName,
			TotalNum:   i.TotalNum,
			Status:     i.Status,
			TotalPrice: i.TotalPrice,
			CreatedAt:  i.CreatedAt,
			PayTime:    i.PayTime,
			RealName:   i.RealName,
		})
	}
	return news, nil
}
