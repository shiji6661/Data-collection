package dao_mysql

import (
	"models/model_user/model_mysql"
	"time"
)

func IsUserSignIn(uid int64, today time.Time, tomorrow time.Time) bool {
	var u model_mysql.UserSign
	err := u.IsUserSignIn(uid, today, tomorrow)
	if err != nil {
		return false
	}
	return true
}
