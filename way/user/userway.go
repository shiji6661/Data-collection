package user

import (
	"models/model_user/model_mysql"
)

// todo:根据用户id查看

func CommonGetUserIdInfo(userId int64) (u *model_mysql.User, err error) {
	u = &model_mysql.User{}
	err = u.FindUserById(userId)
	return u, nil
}
func CommonGetUserInfoById(userId int64) (user *model_mysql.UserDetail, err error) {
	user, err = user.FindUserDetailById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
