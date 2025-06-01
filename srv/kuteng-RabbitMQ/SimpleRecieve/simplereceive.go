package SimpleRecieve

import (
	"common/global"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"kuteng-RabbitMQ/rabbitmq"
	"models/model_order/model_mysql"
	"time"
)

type OrderProduct struct {
	OrderId       string
	Uid           int64
	RealName      string
	UserPhone     string
	UserAddress   string
	CartId        string
	FreightPrice  float64
	TotalNum      int64
	TotalPrice    float64
	PayType       int64
	MerId         int64
	ProductId     int64
	CombinationId int64
	PinkId        int64
}

func SimpleReceive() error {
	fmt.Println("wowoo")
	mq := rabbitmq.NewRabbitMQSimple("" +
		"Order")

	mq.ConsumeSimple(func(b []byte) {
		var S OrderProduct
		err := json.Unmarshal(b, &S)
		if err != nil {
			fmt.Println("错误")
			return
		}
		// 加锁
		lockKey := "order_lock" + fmt.Sprintf(S.OrderId)
		lockValue := "locked"
		lockExpireTime := time.Second * 15
		result, err := global.Rdb.SetNX(context.Background(), lockKey, lockValue, lockExpireTime).Result()
		if err != nil {
			return
		}
		if !result {
			zap.L().Error("订单正在处理...")
			return
		}
		defer func() {
			global.Rdb.Del(context.Background(), lockKey)
		}()

		if S.PinkId == 0 {
			o := &model_mysql.Order{
				OrderId:      S.OrderId,
				Uid:          S.Uid,
				RealName:     S.RealName,
				UserPhone:    S.UserPhone,
				UserAddress:  S.UserAddress,
				CartId:       S.CartId,
				FreightPrice: S.FreightPrice, //运费5元一件
				TotalNum:     S.TotalNum,
				TotalPrice:   S.TotalPrice,
				PayType:      S.PayType,
				MerId:        S.MerId,
				ProductId:    S.ProductId,
			}
			err = o.AddProduct(S.ProductId, S.TotalNum)
			if err != nil {
				zap.L().Info("下单失败")
				return
			}
			err := model_mysql.RewardCommissionForSuperiors(int32(o.Uid), int32(o.Id), o.TotalPrice)
			if err != nil {
				return
			}
		} else {
			o := &model_mysql.Order{
				OrderId:       S.OrderId,
				Uid:           S.Uid,
				RealName:      S.RealName,
				UserPhone:     S.UserPhone,
				UserAddress:   S.UserAddress,
				CartId:        S.CartId,
				FreightPrice:  S.FreightPrice, //运费5元一件
				TotalNum:      S.TotalNum,
				TotalPrice:    S.TotalPrice,
				PayType:       S.PayType,
				MerId:         S.MerId,
				ProductId:     S.ProductId,
				CombinationId: S.CombinationId,
				PinkId:        S.PinkId,
			}
			err = o.AddProduct(S.ProductId, S.TotalNum)
			if err != nil {
				zap.L().Info("拼团下单失败")
				return
			}
			err := model_mysql.RewardCommissionForSuperiors(int32(o.Uid), int32(o.Id), o.TotalPrice)
			if err != nil {
				return
			}
		}

	})
	return nil
}
