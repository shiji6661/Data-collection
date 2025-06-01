package logic

import (
	"user_srv/dao/dao_mysql"
	"user_srv/proto_user/user"
)

// todo 用户详情
func UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userinfo, err := dao_mysql.GetUserInfo(in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserInfoResponse{
		Nickname:   userinfo.Nickname,
		Headimgurl: userinfo.Headimgurl,
		Sex:        userinfo.Sex,
		City:       userinfo.City,
		Language:   userinfo.Language,
		Province:   userinfo.Province,
		Country:    userinfo.Country,
		Stair:      userinfo.Stair,
		Second:     userinfo.Second,
		UserPhone:  userinfo.UserPhone,
		UserEmail:  userinfo.UserEmail,
		UserState:  userinfo.UserState,
	}, nil
}
