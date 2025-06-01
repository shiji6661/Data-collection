package dao_mysql

import "models/model_user/model_mysql"

func FindUserById(userId int64) (u *model_mysql.User, err error) {
	u = &model_mysql.User{}
	err = u.FindUserById(userId)
	return u, nil
}
