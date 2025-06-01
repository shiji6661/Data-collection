package server

import (
	"common/utils/merchant"
	user2 "common/utils/user"
	"context"
	"errors"
	"user_srv/internal/logic"
	"user_srv/proto_user/user"
)

type ServerUser struct {
	user.UnimplementedUserServer
}

// todo:用户注册
func (s ServerUser) UserRegister(ctx context.Context, in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	if in.UserName == "" || in.UserPassword == "" {
		return nil, errors.New("账号或密码为空")
	}
	res, err := logic.UserRegister(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// todo 用户登录

func (s ServerUser) UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	login, err := logic.UserLogin(in)
	if err != nil {
		return nil, err
	}
	return login, nil
}

// todo 短信发送

func (s ServerUser) SendSms(ctx context.Context, in *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	if in.Mobile == "" || in.Source == "" {
		return nil, errors.New("手机号或验证码来源为空")
	}
	if len(in.Mobile) != 11 {
		return nil, errors.New("手机号长度不对")
	}
	if !user2.ValidateMobile(in.Mobile) {
		return nil, errors.New("手机号格式不对")
	}
	sms, err := logic.SendSms(in)
	if err != nil {
		return nil, err
	}
	return sms, nil
}

// todo 邮件发送

func (s ServerUser) SendEmail(ctx context.Context, in *user.SendEmailRequest) (*user.SendEmailResponse, error) {
	if in.UserEmail == "" {
		return nil, errors.New("邮箱账号不能为空")
	}
	email, err := logic.SendEmail(in)
	if err != nil {
		return nil, err
	}
	return email, nil
}

// todo 用户详情

func (s ServerUser) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	info, err := logic.UserInfo(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

//todo:站内信发送信息

func (s ServerUser) SendMessage(ctx context.Context, in *user.SendMessageRequest) (*user.SendMessageResponse, error) {
	message, err := logic.SendMessage(in)
	if err != nil {
		return nil, err
	}
	return message, err
}

func (s ServerUser) FindMessage(ctx context.Context, in *user.FindMessageRequest) (*user.FindMessageResponse, error) {
	message, err := logic.FindMessage(in)
	if err != nil {
		return nil, err
	}
	return message, err
}

// todo 用户邀请码生成
func (s ServerUser) InvitationCodeGeneration(ctx context.Context, in *user.InvitationCodeGenerationRequest) (*user.InvitationCodeGenerationResponse, error) {
	info, err := logic.InvitationCodeGeneration(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 用户修改信息
func (s ServerUser) UserModify(ctx context.Context, in *user.UserModifyRequest) (*user.UserModifyResponse, error) {
	if in.UserPassword != in.Password {
		return nil, errors.New("两次密码输入不一致")
	}
	if !user2.ValidateMobile(in.UserPhone) {
		return nil, errors.New("手机号码格式错误")
	}
	if !merchant.Email(in.UserEmail) {
		return nil, errors.New("邮箱格式错误")
	}
	modify, err := logic.UserModify(in)
	if err != nil {
		return nil, err
	}
	return modify, nil
}

// todo 实名认证
func (s ServerUser) RealName(ctx context.Context, in *user.RealNameRequest) (*user.RealNameResponse, error) {
	name, err := logic.RealName(in)
	if err != nil {
		return nil, err
	}
	return name, nil
}

// todo 用户查看订单
func (s ServerUser) UserViewOrder(ctx context.Context, in *user.UserViewOrderRequest) (*user.UserViewOrderResponse, error) {
	order, err := logic.UserViewOrder(in)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// todo 用户查看订单详情
func (s ServerUser) UserViewOrderInfo(ctx context.Context, in *user.UserViewOrderInfoRequest) (*user.UserViewOrderInfoResponse, error) {
	info, err := logic.UserViewOrderInfo(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 展示会员
func (s ServerUser) UserLevelInfo(ctx context.Context, in *user.UserLevelInfoRequest) (*user.UserLevelInfoResponse, error) {
	info, err := logic.UserLevelInfo(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 用户签到
func (s ServerUser) UserSignIn(ctx context.Context, in *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	order, err := logic.UserSignIn(in)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// todo 用户补签
func (s ServerUser) UserReSigning(ctx context.Context, in *user.UserReSigningRequest) (*user.UserReSigningResponse, error) {
	info, err := logic.UserReSigning(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 展示会员权益
func (s ServerUser) UserLevelRights(ctx context.Context, in *user.UserLevelRightsRequest) (*user.UserLevelRightsResponse, error) {
	rights, err := logic.UserLevelRights(in)
	if err != nil {
		return nil, err
	}
	return rights, nil
}

// todo 用户申请发票
func (s ServerUser) UserApplyInvoice(ctx context.Context, in *user.UserApplyInvoiceRequest) (*user.UserApplyInvoiceResponse, error) {
	info, err := logic.UserApplyInvoice(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 用户使用权益展示
func (s ServerUser) UserUseRights(ctx context.Context, in *user.UserUseRightsRequest) (*user.UserUseRightsResponse, error) {
	rights, err := logic.UserUseRights(in)
	if err != nil {
		return nil, err
	}
	return rights, nil
}

// todo 用户查看自己的发票
func (s ServerUser) UserInvoicesList(ctx context.Context, in *user.UserInvoicesListRequest) (*user.UserInvoicesListResponse, error) {
	info, err := logic.UserInvoices(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 用户修改发票
func (s ServerUser) UpdateInvoice(ctx context.Context, in *user.UpdateInvoiceRequest) (*user.UpdateInvoiceResponse, error) {
	info, err := logic.UserUpdateInvoice(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 用户领取优惠卷
func (s ServerUser) UserReceiveCoupon(ctx context.Context, in *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
	info, err := logic.UserReceiveCoupon(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// todo 会员分添加记录展示
func (s ServerUser) ShowMemberPoints(ctx context.Context, in *user.ShowMemberPointsRequest) (*user.ShowMemberPointsResponse, error) {
	points, err := logic.ShowMemberPoints(in)
	if err != nil {
		return nil, err
	}
	return points, nil
}

// TODO 佣金排行榜
func (s ServerUser) CommissionList(ctx context.Context, in *user.CommissionListRequest) (*user.CommissionListResponse, error) {
	info, err := logic.CommissionRank(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// TODO 邀请用户排行榜
func (s ServerUser) InviteUserList(ctx context.Context, in *user.InviteUserListRequest) (*user.InviteUserListResponse, error) {
	info, err := logic.InviteUserRank(in)
	if err != nil {
		return nil, err
	}
	return info, nil
}
