package logic

import (
	"common/utils"
	"errors"
	"user_srv/dao/dao_mysql"
	"user_srv/dao/dao_redis"
	"user_srv/proto_user/user"
)

func UserLogin(in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	switch in.Login_Type {
	case user.LoginType_LOGIN_TYPE_USERNAME:
		if in.UserName == "" || in.UserPassword == "" {
			return nil, errors.New("用户名或密码不能为空")
		}
		name, err := dao_mysql.FindUserByUserName(in.UserName)
		if err != nil {
			return nil, errors.New("查询失败123")
		}
		if name.ID == 0 {
			return nil, errors.New("用户不存在")
		}
		if utils.Md5(in.UserPassword) != name.UserPassword {
			return nil, errors.New("密码错误")
		}
		// 检查并更新会员等级
		_, err = dao_mysql.CheckAndUpdateMemberLevel(int(name.ID))
		if err != nil {
			return nil, errors.New("会员等级更新失败")
		}
		return &user.UserLoginResponse{Greet: int64(name.ID)}, nil
	case user.LoginType_LOGIN_TYPE_PHONE:
		if in.UserPhone == "" || in.MobileCode == "" {
			return nil, errors.New("手机号或验证码不能为空")
		}
		err := dao_redis.SendSmsGet("SendSmsLogin"+in.UserPhone, in.MobileCode)
		if err != nil {
			return nil, errors.New("验证码错误")
		}
		phone, err := dao_mysql.FindUserByPhone(in.UserPhone)
		if err != nil {
			return nil, errors.New("登录失败")
		}
		if phone.ID == 0 {
			return nil, errors.New("该手机号还未注册")
		}
		// 检查并更新会员等级
		_, err = dao_mysql.CheckAndUpdateMemberLevel(int(phone.ID))
		if err != nil {
			return nil, errors.New("会员等级更新失败")
		}
		return &user.UserLoginResponse{Greet: int64(phone.ID)}, nil
	case user.LoginType_LOGIN_TYPE_EMAIL:
		if in.UserEmail == "" || in.EmailCode == "" {
			return nil, errors.New("邮箱或验证码不能为空")
		}
		email, err := dao_mysql.FindUserByEmail(in.UserEmail)
		if err != nil {
			return nil, err
		}
		if email.ID == 0 {
			return nil, errors.New("该邮箱还未绑定")
		}
		err = dao_redis.SendSmsGet("Email"+in.UserEmail, in.EmailCode)
		if err != nil {
			return nil, errors.New("验证码错误")
		}
		// 检查并更新会员等级
		_, err = dao_mysql.CheckAndUpdateMemberLevel(int(email.ID))
		if err != nil {
			return nil, errors.New("会员等级更新失败")
		}
		return &user.UserLoginResponse{Greet: int64(email.ID)}, nil
	default:
		return nil, errors.New("登录方式不对")
	}

}
