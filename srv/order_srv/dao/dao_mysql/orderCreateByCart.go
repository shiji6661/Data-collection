package dao_mysql

import (

	"common/global"
	"context"
	"errors"
	"github.com/google/uuid"
	"models/model_order/model_mysql"
	"order_srv/dao/dao_redis"
	"strconv"
	"time"
	"way/user"
)

const (
	FREIGHT_PRICE = 2 // 运费
	Paid          = 0 //支付状态
	PayType       = 1 //支付类型
	Status        = 0 //订单状态
	MerId         = 1 //商户id
	CombinationId = 0 //组合id
	PinkId        = 0 //拼团id
	Cost          = 1 //成本
	SeckillId     = 1 //秒杀id
	BargainId     = 1 //砍价id
	StoreId       = 1 //店铺id
)

func OrderCreateByCart(userId, productId int64) (o *model_mysql.Order, err error) {
	cartId, err := dao_redis.GetCartIdByUserId(context.Background(), strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, err
	}

	// 检查指定商品是否存在于购物车
	if !dao_redis.IsCartProductExists(context.Background(), cartId, userId, productId) {
		return nil, errors.New("商品不存在于购物车！请重新选择商品！")
	}

	// 判断库存是否足够
	redisStock := dao_redis.GetProductsFromRedis(context.Background(), productId)
	if redisStock <= 0 {
		return nil, errors.New("库存不足")
	}

	// 根据商品id和购物车id查询商品信息
	productInfo := dao_redis.GetCartProductInfo(context.Background(), cartId, userId, productId)

	// 查询用户信息
	users, err := user.CommonGetUserIdInfo(userId)
	if err != nil {
		return nil, err
	}
	// 查询用户地址信息
	info, err := user.CommonGetUserInfoById(userId)
	if err != nil {
		return nil, err
	}
	// 生成 uuid
	orderId := uuid.New().String()
	o = &model_mysql.Order{
		OrderId:       orderId,
		Uid:           userId,
		RealName:      users.UserName,
		UserPhone:     users.UserPhone,
		UserAddress:   info.City,
		CartId:        cartId,
		FreightPrice:  FREIGHT_PRICE,
		TotalNum:      productInfo.Num,
		TotalPrice:    productInfo.ProductPrice,
		Paid:          Paid,
		PayTime:       time.Now().Format("2006-01-02 15:04:05"),
		PayType:       PayType,
		Status:        Status,
		MerId:         MerId,
		CombinationId: CombinationId,
		PinkId:        PinkId,
		Cost:          Cost,
		SeckillId:     SeckillId,
		BargainId:     BargainId,
		StoreId:       StoreId,
		ProductId:     productInfo.ProductId,
	}
	err = global.DB.Create(&o).Error
	if err != nil {
		return nil, err
	}

	return o, nil
}
