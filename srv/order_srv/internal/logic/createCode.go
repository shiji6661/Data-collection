package logic

import (
	"common/utils/user"
	"fmt"
	"log"
	"models/model_order/model_mysql"
	"order_srv/proto_order/order"
)

// 二维码生成
func CreateCode(in *order.CreateOrCodeRequest) (*order.CreateOrCodeResponse, error) {
	id, err := model_mysql.FindByOrderId(in.OrderId)
	if err != nil {
		log.Printf("Failed to find order by ID %d:%s", in.OrderId, err)
		return nil, err
	}
	if id.Id == 0 {
		return nil, err
	}
	data := map[string]interface{}{
		"id":     in.OrderId,
		"mer_id": in.MerId,
		"uid":    in.UserId,
	}
	code, err := user.MakeQrCode(int(in.UserId), data)
	if err != nil {
		return nil, err
	}
	log.Printf("or code:%s", code)
	return &order.CreateOrCodeResponse{Success: GetQrCodeImageUrl(code)}, nil
}

func GetQrCodeImageUrl(filename string) string {
	return fmt.Sprintf("http://127.0.0.1:8888/order/qr/code/%s", filename)
}
