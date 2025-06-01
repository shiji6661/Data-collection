package dao_mysql

import "models/model_user/model_mysql"

func Login(account string) (*model_mysql.User, error) {
	err := U.Login(account)
	if err != nil {
		return nil, err
	}
	return nil, err
}
