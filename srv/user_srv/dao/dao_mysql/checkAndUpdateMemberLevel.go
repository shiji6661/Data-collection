package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
	"time"
)

var M *model_mysql.Member

// 检查并更新会员等级
func CheckAndUpdateMemberLevel(uid int) (int64, error) {
	// 获取用户会员信息
	member, err := M.GetMemberInfo(uid)
	if err != nil {
		return 0, errors.New("该用户不是会员")
	}

	now := time.Now()
	// 如果是新用户，初始化会员信息
	if member == nil {
		// 创建新会员记录，初始等级为0，开始时间为现在
		err = M.CreateMemberInfo(int64(uid), 0, 0, now)
		return 0, err
	}

	// 检查是否已经过了一年
	if now.Sub(member.UpdatedAt) >= 365*24*time.Hour {
		// 如果是最高等级且积分不足，降级
		if member.LevelId == 5 && member.Points < 2000 {
			err = M.UpdateMemberInfo(int64(uid), 4, 0, now) // 重置积分为0
			return 4, err
		}
		// 如果不是最高等级且积分不足，降级
		if member.LevelId < 5 && member.Points < 2000 {
			newLevel := member.LevelId - 1
			if newLevel < 0 {
				newLevel = 0
			}
			err = M.UpdateMemberInfo(int64(uid), newLevel, 0, now) // 重置积分为0
			return newLevel, err
		}
		// 如果积分足够，重置积分
		if member.Points >= 2000 {
			err = M.UpdateMemberInfo(int64(uid), member.LevelId, 0, now) // 重置积分为0但保持等级
			return member.LevelId, err
		}
	}

	// 检查是否可以升级
	if member.Points >= 2000 {
		// 如果未达到最高等级，则可以升级
		if member.LevelId < 5 {
			newLevel := member.LevelId + 1
			// 升级后重置积分为0
			err = M.UpdateMemberInfo(int64(uid), newLevel, 0, now)
			return newLevel, err
		}
		// 如果已是最高等级且达到2000积分，只重置积分
		if member.LevelId == 5 {
			err = M.UpdateMemberInfo(int64(uid), 5, 0, now)
			return 5, err
		}
	}

	return member.LevelId, nil
}
