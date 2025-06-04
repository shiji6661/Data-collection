package logic

import (
	"common/global"
	"errors"
	"models/model_user/model_mysql"
	"time"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo 用户补签
func UserReSigning(in *user.UserReSigningRequest) (*user.UserReSigningResponse, error) {
	// 1. 解析补签日期
	makeupDate, err := time.Parse("2006-01-02", in.SignDate)
	if err != nil {
		return nil, errors.New("无效的补签日期格式")
	}
	// 检查当前日期是否已签到
	var count int64
	result := global.DB.Model(&model_mysql.UserSign{}).Where("uid = ? AND add_time = ?", in.Uid, makeupDate).Count(&count)
	if result.Error != nil {
		return nil, errors.New("查询日期失败")
	}
	if count > 0 {
		return nil, errors.New("该日期已签到，不能补签")
	}

	// 2. 检验补签日期 是不是在一周之内的
	now := time.Now()
	if makeupDate.After(now) { // 补签日期不能是未来时间
		return nil, errors.New("补签日期不能是未来时间")
	}
	if now.Sub(makeupDate) > 7*24*time.Hour { // 改用 time.Time.Sub() 更清晰
		return nil, errors.New("只能补签过去7天内的签到")
	}

	// 4. 检查用户是否有补签卡
	makeupCard := &model_mysql.UserMakeup{}
	err = makeupCard.GetUserMakeupCard(in.Uid)
	if err != nil {
		return nil, errors.New("没有可用的补签卡")
	}
	if makeupCard.Cardcount <= 0 {
		return nil, errors.New("没有可用的补签卡")
	}

	// 5. 计算积分（补签固定得1分）
	points := 1

	// 6. 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//  扣除补签卡
	err = makeupCard.UpdateUserMakeupCard(in.Uid)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("扣除补签卡失败")
	}

	// 更新用户积分
	ui := &model_mysql.User{}
	err = ui.UpdateUserIntegral(in.Uid, int64(points))
	if err != nil {
		tx.Rollback()
		return nil, errors.New("更新积分失败")
	}
	makeup := 1
	//添加用户补签到记录
	err = dao_mysql.UserSignIn(in.Uid, makeupDate, makeup)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("用户签到失败")
	}

	//恢复用户连续签到日期
	var context int64
	num := global.DB.Debug().Model(&model_mysql.UserSign{}).Where("uid = ? ", in.Uid).Count(&context) // 修正 SQL 参数并传递指针
	if num.Error != nil {
		tx.Rollback()
		return nil, errors.New("查询日期失败")
	}
	err = dao_mysql.RestoreUserSignInTimes(in.Uid, context)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("恢复用户连续签到日期失败")
	}

	// 11. 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, errors.New("提交事务失败")
	}

	return &user.UserReSigningResponse{
		Success: true,
		Message: "补签成功",
		Points:  int64(points),
	}, nil

}
