package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
)

func FindUserModifyById(id int64) (*model_mysql.User, error) {
	modify, err := U.FindUserModify(id)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if modify.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	return modify, nil
}
