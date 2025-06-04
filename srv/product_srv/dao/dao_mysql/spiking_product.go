package dao_mysql

import (
	"common/global"
	"errors"
	"github.com/google/uuid"
	"kuteng-RabbitMQ/SimlpePublish"
	"models/model_product/model_mysql"
	"models/model_product/model_redis"
	o "order_srv/pkg"
	"product_srv/pkg"
	"product_srv/proto_product/product"
	"strconv"
	"time"
)

func FlashSale(in *product.FlashSaleRequest) (response *product.FlashSaleResponse, err error) {
	result := model_mysql.SpikeProducts{}
	re, err := result.FindSpikeProductsById(in.SpikeProductsId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if re.ID == 0 {
		return nil, errors.New("秒杀商品不存在")
	}

	p := model_mysql.Product{}
	pro, err := p.FindProductById(re.ProductId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if pro.ID == 0 {
		return nil, errors.New("商品不存在")
	}

	now := time.Now()
	StartTime, err := pkg.ParseTime(re.StartTime)
	if err != nil {
		return nil, errors.New("开始时间解析失败")
	}
	EndTime, err := pkg.ParseTime(re.EndTime)
	if err != nil {
		return nil, errors.New("结束时间解析失败")
	}
	if now.After(StartTime) && now.Before(EndTime) {
		tx := global.DB.Begin()
		err = tx.Model(&model_mysql.SpikeProducts{}).Where("id=?", re.ID).Update("spike_status", 2).Error
		if err != nil {
			tx.Rollback()
			return nil, errors.New("秒杀开启失败")
		}
		spikeLen := model_redis.SpikeLen(int(re.ID))
		if spikeLen <= 0 {
			tx.Rollback()
			tx.Model(&model_mysql.SpikeProducts{}).Where("id=?", re.ID).Update("spike_status", 3)
			return nil, errors.New("商品已售罄")
		}

		err = model_redis.SpikeRPop(int(re.ID))
		if err != nil {
			return nil, errors.New("抢购失败")
		}

		s2 := now.Format("2006-01-02 15:04:05")
		orderId := uuid.New().String() + s2
		SecondKillProduct := map[string]interface{}{
			"orderId":    orderId,
			"productNum": re.SpikeNumber,
			"orderPaid":  re.SpikePrice,
			"title":      re.ProductName,
		}
		err = SimlpePublish.SimplePublishb(SecondKillProduct)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("存入消息队列失败")
		}
		tx.Commit()
		pay := o.NewAliPay()
		Price := strconv.FormatFloat(re.SpikePrice, 'f', 2, 64)
		s := pay.Pay(re.ProductName, orderId, Price)
		return &product.FlashSaleResponse{Success: "抢购成功" + s}, err
	} else if now.Before(StartTime) {
		return nil, errors.New("秒杀活动尚未开始")
	} else {
		return nil, errors.New("秒杀活动已结束")
	}

}
