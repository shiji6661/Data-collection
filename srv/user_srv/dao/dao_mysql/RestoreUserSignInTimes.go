package dao_mysql

import "models/model_user/model_mysql"

func RestoreUserSignInTimes(context int64, uid int64) error {
	u := model_mysql.User{}
	err := u.RestoreUserSignInTimes(uid, context)
	if err != nil {
		return err
	}
	return nil
}
