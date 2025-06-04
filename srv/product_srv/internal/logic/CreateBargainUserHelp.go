package logic

import (
	"errors"
	"models/model_product/model_mysql"
	"time"
	"weikang/Data-collection/srv/product_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// 用户帮砍
func CreateBargainUserHelp(in *product.CreateBargainUserHelpRequest) (*product.CreateBargainUserHelpResponse, error) {
	result := model_mysql.BargainUserHelp{}
	_, err := result.BargainUserCount(int(in.Uid), int(in.BargainUserId))
	if err != nil {
		return nil, err
	}

	s := model_mysql.Bargain{}
	bar, err := s.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, err
	}

	bu := model_mysql.BargainUser{}
	code, err := bu.FindBargainUserCode(in.InviteCode)
	if in.InviteCode != code.InviteCode {
		return nil, errors.New("邀请码错误")
	}

	now := time.Now()
	if now.Before(bar.StartTime) {
		return nil, errors.New("砍价未开始")
	}
	if now.After(bar.StopTime) {
		return nil, errors.New("砍价已结束")
	}
	help, err := dao_mysql.CreateBargainUserHelp(in)
	if err != nil {
		return nil, err
	}
	return &product.CreateBargainUserHelpResponse{Result: int64(help.Id)}, err
}
