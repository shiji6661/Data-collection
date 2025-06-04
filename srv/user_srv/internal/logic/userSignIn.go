package logic

import (
	"common/global"
	"context"
	"errors"
	"fmt"
	"product_srv/dao/dao_redis"
	"time"
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

// 用户签到
func UserSignIn(in *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	now := time.Now()
	realInfo, err := dao_mysql.FindUserById(in.Uid)
	if err != nil {
		return nil, err
	}
	if realInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	// 将时间截断到日期
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	// IsUserSignIn 检查用户当天是否已签到
	data := dao_mysql.IsUserSignIn(in.Uid, today, tomorrow)
	if data == true {
		return nil, errors.New("你今天已经签到过了")
	}
	// GetLastSignDate 获取用户上一次的签到日期
	lastSignDate, err := dao_mysql.GetLastSignDate(in.Uid)
	if err != nil {
		return nil, errors.New("获取用户上一次的签到日期失败")
	}

	// todo 断签功能实现
	// todo !lastSignDate.IsZero() 用于检查 lastSignDate 是否为零值（即是否有上一次签到记录）。
	// todo lastSignDate.AddDate(0, 0, 1) 会将 lastSignDate 往后推一天。
	// todo Truncate(24*time.Hour) 会将时间截断到当天的零点。
	// todo 最后将截断后的时间与 today（当前日期）进行比较，如果不相等，说明用户断签了。
	if !lastSignDate.IsZero() && lastSignDate.AddDate(0, 0, 1).Truncate(24*time.Hour) != today {
		// 断签，重置连续签到天数
		err = dao_mysql.ResetSignInDays(in.Uid)
		if err != nil {
			return nil, err
		}
	}
	//todo 存入redis位图
	signDate := time.Now()
	todays := signDate.Format("2006-01-02")
	dayOfYear := today.YearDay()
	key := fmt.Sprintf("sign:user:%d:%s", in.Uid, todays)
	ctx := context.Background()
	err = global.Rdb.SetBit(ctx, key, int64(dayOfYear), 1).Err()
	if err != nil {
		fmt.Printf("Failed to sign in: %v\n", err)
	}
	//todo 查询连续签到天数数量
	count, err := dao_redis.GetSignCount(in.Uid)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	//todo 默认为签到0
	makeup := 0
	//添加用户签到记录
	err = dao_mysql.UserSignIn(in.Uid, now, makeup)
	if err != nil {
		return nil, errors.New("用户签到失败")
	}

	return &user.UserSignInResponse{
		Success:    true,
		Message:    fmt.Sprintf("用户%d签到成功", in.Uid),
		Continuous: fmt.Sprintf("连续签到%d", count),
	}, nil
}
