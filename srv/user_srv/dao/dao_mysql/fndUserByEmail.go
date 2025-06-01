package dao_mysql

import "models/model_user/model_mysql"

// todo:查询用户邮箱是否存在
func FindUserByEmail(email string) (*model_mysql.User, error) {
	byEmail, err := U.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return byEmail, nil
}
