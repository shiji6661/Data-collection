package logic

import (
	user2 "common/utils/user"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"user_srv/dao/dao_mysql"
	"user_srv/dao/dao_redis"
	"user_srv/proto_user/user"
)

// todo:邮件发送逻辑
func SendEmail(in *user.SendEmailRequest) (*user.SendEmailResponse, error) {
	email, err := dao_mysql.FindUserByEmail(in.UserEmail)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if email.ID == 0 {
		return nil, errors.New("该账号未绑定邮箱")
	}
	code := rand.Intn(9000) + 1000
	dao_redis.SendSmsSet("Email"+in.UserEmail, code)
	err = user2.SendEmail(in.UserEmail, strconv.Itoa(code))
	if err != nil {
		log.Println("邮件发送失败")
	}

	return &user.SendEmailResponse{Greet: "邮件发送成功"}, nil
}
