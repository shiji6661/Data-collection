package dao_mysql

import "models/model_user/model_mysql"

// todo:查找用户填写的邀请码是否真实
func FindUserInvite(inviteCode string) (*model_mysql.User, error) {
	code, err := U.FindUserByInviteCode(inviteCode)
	if err != nil {
		return nil, err
	}

	return code, nil
}
