package logic

import (
	"common/global"
	user2 "common/utils/user"
	"errors"
	"fmt"
	"models/model_user/model_mysql"
	"time"
	"weikang/Data-collection/srv/user_srv/dao/dao_mysql"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

var U *model_mysql.User

// todo:用户注册逻辑
func UserRegister(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	var id *model_mysql.User
	fmt.Println(in.InviteCode)
	users, err := dao_mysql.FindUserByUserName(in.UserName)
	if users.ID != 0 {
		return nil, errors.New("用户已存在")
	}
	var separator uint64

	if in.InviteCode != "" {
		_, err = dao_mysql.FindUserInvite(in.InviteCode)
		if err != nil {
			return nil, errors.New("邀请码错误")
		}
		separator, err = user2.ParseInviteCode(in.InviteCode)
		if err != nil {
			return nil, errors.New("获取邀请码关联的用户ID失败")
		}
	}

	tx := dao_mysql.BeginTransaction()

	tx.Begin()

	createUser, err := dao_mysql.CreateUser(in)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("用户注册失败")
	}

	// GenerateInviteCode  todo 根据用户 ID 生成邀请码
	InviteCode, err := user2.GenerateInviteCode(uint64(createUser.ID))
	global.DB.Model(&model_mysql.User{}).Where("id=?", createUser.ID).Update("invite_code", InviteCode)

	// 创建会员信息
	member := &model_mysql.Member{}
	err = member.CreateMemberInfo(int64(createUser.ID), 0, 1, time.Now())
	if err != nil {
		tx.Rollback()
		return nil, errors.New("创建会员信息失败")
	}

	//se, err := user2.ParseInviteCode(in.InviteCode)
	fmt.Println(in.InviteCode)

	if in.InviteCode != "" {
		fmt.Println(separator)
		id, err = dao_mysql.FindUserById(int64(createUser.ID))
		if err != nil {
			tx.Rollback()
			return nil, errors.New("查询邀请码关联的用户信息失败")
		}

		err = dao_mysql.UpdateUserPoint(separator, int(id.UserPoint+5))
		if err != nil {
			tx.Rollback()
			return nil, errors.New("更新邀请码关联的用户积分失败")
		}

		member = &model_mysql.Member{}
		err = member.AddPointsByInvitation(int64(separator), int64(createUser.ID))
		if err != nil {
			tx.Rollback()
			return nil, errors.New("会员分记录添加失败")
		}
	}
	tx.Commit()
	return &user.UserRegisterResponse{UserId: int64(createUser.ID)}, nil
}
