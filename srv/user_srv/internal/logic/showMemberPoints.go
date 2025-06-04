package logic

import (
	"errors"
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

// todo 会员分添加记录展示
func ShowMemberPoints(in *user.ShowMemberPointsRequest) (*user.ShowMemberPointsResponse, error) {
	list, err := dao_mysql.FindMemberPoints(in.Userid)
	if err != nil {
		return nil, errors.New("列表查询失败")
	}
	return &user.ShowMemberPointsResponse{List: list}, nil
}
