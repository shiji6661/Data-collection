package logic

import (
	"common/global"
	user2 "common/utils/user"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

// todo 用户邀请码生成
func InvitationCodeGeneration(in *user.InvitationCodeGenerationRequest) (*user.InvitationCodeGenerationResponse, error) {
	Invitation, err := user2.GenerateInviteCode(uint64(in.UserId))
	global.DB.Model(&model_mysql.User{}).Where("id=?", in.UserId).Update("invite_code", Invitation)
	if err != nil {
		return nil, err
	}
	return &user.InvitationCodeGenerationResponse{InvitationCode: Invitation}, nil
}
