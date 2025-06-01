package dao_mysql

import (
	"common/global"
	"errors"
	"github.com/google/uuid"
	"models/model_order/model_mysql"
	mysql2 "models/model_product/model_mysql"
	mysql3 "models/model_user/model_mysql"
	"order_srv/proto_order/order"
	"time"
)

func CreateOrderBargain(in *order.CreateOrderBargainRequest) (o *model_mysql.Order, err error) {
	b := mysql2.Bargain{}
	id, err := b.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, err
	}
	if id.ID == 0 {
		return nil, errors.New("砍价商品不存在")
	}
	c := mysql2.BargainUser{}
	userId, err := c.FindBargainUserId(int(in.UserId))
	if err != nil {
		return nil, err
	}
	if userId.Id == 0 {
		return nil, errors.New("未发起砍价")
	}
	u := mysql3.User{}
	bargain, err := u.FindUserByIdBargain(in.UserId)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}
	sn := uuid.New().String()
	s := time.Now().Format("2006-01-02 15:04:05")
	o = &model_mysql.Order{
		OrderId:       sn,
		Uid:           int64(userId.Uid),
		RealName:      bargain.UserName,
		UserPhone:     bargain.UserPhone,
		TotalNum:      1,
		TotalPrice:    userId.Price,
		Paid:          0,
		PayTime:       s,
		PayType:       1,
		Status:        0,
		MerId:         1,
		CombinationId: 0,
		PinkId:        0,
		SeckillId:     0,
		BargainId:     1,
		StoreId:       1,
	}
	err = global.DB.Create(&o).Error
	if err != nil {
		return
	}
	return o, nil
}
