package dao_mysql

import (
	"errors"
)

// 更新会员积分
func updateMemberPoints(uid int, points int) error {
	member, err := M.GetMemberInfo(uid)
	if err != nil {
		return err
	}

	if member == nil {
		return errors.New("会员信息不存在")
	}

	// 更新积分
	newPoints := member.Points + int64(points)
	return M.UpdateMemberPoints(uid, newPoints)
}
