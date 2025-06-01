package dao_mysql

import (
	"models/model_user/model_mysql"
	"time"
)

func GetLastSignDate(uid int64) (time.Time, error) {
	data, err := model_mysql.GetLastSignDate(uid)
	if err != nil {
		return time.Time{}, err
	}
	return data, err
}
