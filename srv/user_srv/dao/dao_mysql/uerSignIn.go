package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
	"time"
)

func UserSignIn(uid int64, now time.Time, makeup int) error {
	u := model_mysql.UserSign{
		Uid:      int32(uid),
		Title:    "日签到",
		Number:   5,
		AddTime:  now,
		IsMakeup: int32(makeup),
	}
	var points int64
	point := points + 5
	var frequencys int64
	frequency := frequencys + 1
	err := u.UserSigns(int(point), frequency)
	if err != nil {
		return errors.New("用户签到失败")
	}
	return nil
}
