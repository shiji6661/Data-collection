package logic

import (
	"fmt"
	"models/model_user/model_mysql"
	"user_srv/proto_user/user"
)

func InviteUserRank(in *user.InviteUserListRequest) (*user.InviteUserListResponse, error) {
	u := model_mysql.User{}
	list, err := u.InviteUserList()
	if err != nil {
		fmt.Printf("获取邀请用户列表出错: %v", err)
		return nil, err
	}

	// 遍历并打印排序后的邀请用户列表
	for _, item := range list {
		fmt.Printf("用户ID: %d, 用户名: %s, 邀请数量: %d\n", item.UserId, item.UserName, item.InviteNum)
	}
	return &user.InviteUserListResponse{List: list}, nil
}
