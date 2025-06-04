package dao_mysql

import (
	"errors"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

var MPR *model_mysql.MemberPointsRecord

func FindMemberPoints(uid int64) (newList []*user.PointsList, err error) {
	getPoints, err := MPR.GetPoints(uid)
	if err != nil {
		return nil, errors.New("列表查询失败")
	}
	var list []*user.PointsList
	for _, i := range getPoints {
		list = append(list, &user.PointsList{
			Points:     int64(i.Points),
			Type:       int64(i.Type),
			OrderId:    i.OrderId,
			Amount:     float32(i.Amount),
			InvitedUid: int64(i.InvitedUid),
			CreatedAt:  i.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return list, nil
}
