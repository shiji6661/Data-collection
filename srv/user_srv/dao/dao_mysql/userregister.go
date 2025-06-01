package dao_mysql

import "models/model_user/model_mysql"

func Register() (*model_mysql.User, error) {
	err := U.Register()
	if err != nil {
		return nil, err
	}
	return nil, err
}
