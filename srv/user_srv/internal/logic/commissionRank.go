package logic

import (
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

func CommissionRank(in *user.CommissionListRequest) (*user.CommissionListResponse, error) {
	u := model_mysql.User{}
	if in.Page <= 0 {
		in.Page = 1
	}
	switch {
	case in.Size > 100:
		in.Size = 100
	case in.Size <= 0:
		in.Size = 10
	}
	commissionList, err := u.CommissionList(int(in.Page), int(in.Size))
	if err != nil {
		return nil, err
	}
	var list []*user.CommissionList
	for _, u2 := range commissionList {
		lists := &user.CommissionList{
			UserId:       int64(u2.ID),
			UserName:     u2.UserName,
			SubUserPrice: float32(u2.SubUserPrice),
		}
		list = append(list, lists)
	}
	return &user.CommissionListResponse{
		List: list,
	}, nil
}
