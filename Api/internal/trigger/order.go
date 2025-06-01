package trigger

import (
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"order_srv/proto_order/order"
)

func CreateOrder(c *gin.Context) {
	var data request.CreateOrder
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	id := c.GetUint("userId")

	register, err := handler.CreateOrder(c, &order.CreateOrderRequest{
		UserId:    int64(id),
		ProductId: data.ProductId,
		Num:       data.Num,
		Pay:       data.Pay,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

func CreateOrderCart(c *gin.Context) {
	var data request.CreateOrderCarts
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}

	id := c.GetUint("userId")

	register, err := handler.CreateOrderCart(c, &order.CreateOrderCartRequest{
		UserId:    int64(id),
		ProductId: data.ProductId,
		Pay:       data.Pay,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// todo 订单支付
func OrderPay(c *gin.Context) {
	var data request.OrderPay
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	register, err := handler.OrderPay(c, &order.OrderPayRequest{
		UserId:  int64(id),
		OrderId: data.OrderId,
		Pay:     data.Pay,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

func CreateOrderBargain(c *gin.Context) {
	var data request.CreateOrderBargain
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	bargain, err := handler.CreateOrderBargain(c, &order.CreateOrderBargainRequest{
		BargainId: data.BargainId,
		UserId:    int64(id),
	})
	response.ResponseSuccess(c, bargain)
}

// todo 支付回调
func OrderCallBack(c *gin.Context) {
	OrderSn := c.Request.FormValue("out_trade_no")
	status := c.Request.FormValue("trade_status")

	var paId int64

	switch status {
	case "TRADE_SUCCESS":
		paId = 1
	case "TRADE_FINISHED":
		paId = 2
	case "TRADE_CLOSED":
		paId = 2
	case "WAIT_BUYER_PAY":
		paId = 0
	}

	fmt.Println(paId)
	fmt.Println(status)

	register, err := handler.OrderCallBack(c, &order.OrderPayBakeCallRequest{
		OrderId: OrderSn,
		Status:  paId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// todo 消费队列进行下单
func RabbitMqToCreate(c *gin.Context) {
	register, err := handler.RabbitMqToCreate(c, &order.CreateOrderMessageRequest{})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// 生成二维码
func CreateOrCode(c *gin.Context) {
	var data request.CreateOrCodeReq
	if err := c.ShouldBind(&data); err != nil {
	}
	id := c.GetUint("userId")
	register, err := handler.CreateOrderQrCode(c, &order.CreateOrCodeRequest{
		UserId:  int64(id),
		OrderId: data.OrderId,
		MerId:   data.MerId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
	}
	response.ResponseSuccess(c, register)
}
func OrderCreateSpike(c *gin.Context) {
	var data request.OrderCreateSpike
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	spike, err := handler.OrderCreateSpike(c, &order.OrderCreateRequest{UserId: int64(id)})

	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, spike)

}

func MerchantCountAmount(c *gin.Context) {
	id := c.GetUint("userId")
	sales, err := handler.MerchantSales(c, &order.MerchantCountAmountRequest{
		MerId: int64(id),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, sales)

}
