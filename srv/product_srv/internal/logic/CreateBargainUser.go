package logic

import (
	"errors"
	"models/model_product/model_mysql"
	"product_srv/dao/dao_mysql"
	"product_srv/proto_product/product"
	"time"
)

// 砍价
func CreateBargainUser(in *product.CreateBargainUserRequest) (*product.CreateBargainUserResponse, error) {
	s := model_mysql.Bargain{}
	bar, err := s.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if now.Before(bar.StartTime) {
		return nil, errors.New("砍价未开始")
	}
	if now.After(bar.StopTime) {
		return nil, errors.New("砍价已结束")
	}
	user, err := dao_mysql.CreateBargainUser(in)
	if err != nil {
		return nil, errors.New("创建砍价失败")
	}
	return &product.CreateBargainUserResponse{
		Result:  int64(user.Id),
		Message: "砍价邀请码:" + user.InviteCode,
	}, nil
}
