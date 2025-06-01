package request

type CreateOrder struct {
	ProductId int64 `form:"productId" json:"productId" binding:"required"`
	Num       int64 `form:"num" json:"num" binding:"required"`
	Pay       int64 `form:"pay" json:"pay"`
}
type CreateOrderCarts struct {
	ProductId int64 `form:"productId" json:"productId" binding:"required"`
	Pay       int64 `form:"pay" json:"pay"`
}

type OrderPay struct {
	OrderId int64 `form:"orderId" json:"orderId" binding:"required"`
	Pay     int64 `form:"pay" json:"pay"  binding:"required"`
}

type CreateOrderBargain struct {
	BargainId int64 `json:"bargain_id" binding:"required"`
}

type CreateOrCodeReq struct {
	OrderId int64 `json:"order_id" form:"order_id" binding:"required"`
	MerId   int64 `json:"mer_id" form:"mer_id" binding:"required"`
}
type OrderCreateSpike struct {
}
