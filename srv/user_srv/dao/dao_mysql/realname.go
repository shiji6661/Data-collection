package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
)

func FindRealName(card string) error {
	r := &model_mysql.Real{}
	err := r.FindRealName(card)
	if err != nil {
		return errors.New("查询失败")
	}
	return nil
}

func CreateRealName(name string, card string, uid int64) (*model_mysql.Real, error) {
	r := &model_mysql.Real{
		Uid:      uid,
		RealName: name,
		CardNo:   card,
	}
	err := r.CreateRealName()
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return r, nil
}
