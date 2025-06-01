package logic

import (
	merchant2 "common/utils/merchant"
	"errors"
	"fmt"
	"merchant_srv/proto_merchant/merchant"
	"models/model_merchant/model_mysql"
)

type Order struct {
	OrderID int `json:"id"`
	UserId  int `json:"uid"`
	MerId   int `json:"mer_id"`
}

func ParseCode(in *merchant.ParseCodeRequest) (*merchant.ParseCodeResponse, error) {
	result, err := merchant2.ParseQrCode[Order](in.Filepath)
	if err != nil {
		fmt.Println("解析二维码时出错:", err)
		return nil, err
	}
	m := model_mysql.MerchantParse{
		UserId:  int32(result.UserId),
		OrderId: int32(result.OrderID),
		MerId:   int32(result.MerId),
	}
	err = m.CreateParse()
	if err != nil {
		return nil, errors.New("商家核销添加失败")
	}
	return &merchant.ParseCodeResponse{
		UserId:  int64(result.UserId),
		OrderId: int64(result.OrderID),
		MerId:   int64(result.MerId),
	}, nil
}
