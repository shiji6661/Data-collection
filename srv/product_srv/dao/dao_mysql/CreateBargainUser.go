package dao_mysql

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"models/model_product/model_mysql"
	"weikang/Data-collection/srv/product_srv/proto_product/product"
)

// todo:用户创建砍价表
func CreateBargainUser(in *product.CreateBargainUserRequest) (result *model_mysql.BargainUser, err error) {
	s := model_mysql.Bargain{}
	bar, err := s.FindBargainProductById(int(in.BargainId))
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if bar.ID == 0 {
		return nil, errors.New("砍价商品不存在")
	}
	fmt.Println(1)
	result = &model_mysql.BargainUser{
		Uid:             uint32(in.Uid),
		BargainId:       uint32(in.BargainId),
		BargainPriceMin: bar.MinPrice,
		BargainPrice:    bar.Price,
		Price:           bar.Price,
	}

	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("生成邀请码失败")
	}
	result.InviteCode = uuidObj.String()

	err = result.CreateBargainUser()
	if err != nil {
		return nil, errors.New("添加砍价失败")
	}
	fmt.Println(result.Uid)
	return result, nil
}
