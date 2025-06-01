package dao_mysql

import "models/model_user/model_mysql"

func ResetSignInDays(uid int64) error {
	u := model_mysql.UserSign{}
	err := u.ResetSignInDays(uid)
	if err != nil {
		return err
	}
	return nil
}
