package handler

import (
	"Api/client"
	"context"
	"github.com/gin-gonic/gin"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:用户账号密码注册
func UserRegister(ctx context.Context, i *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		register, err := in.UserRegister(ctx, i)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserRegisterResponse), nil
}

// todo:短信发送
func SendSms(ctx context.Context, i *user.SendSmsRequest) (*user.SendSmsResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		sms, err := in.SendSms(ctx, i)
		if err != nil {
			return nil, err
		}
		return sms, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.SendSmsResponse), nil
}

// todo: 邮件发送
func SendEmail(ctx context.Context, i *user.SendEmailRequest) (*user.SendEmailResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		email, err := in.SendEmail(ctx, i)
		if err != nil {
			return nil, err
		}
		return email, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.SendEmailResponse), nil
}

// todo: 用户登录
func UserLogin(ctx context.Context, i *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		login, err := in.UserLogin(ctx, i)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserLoginResponse), nil
}

// todo 用户详情
func UserInfo(ctx context.Context, i *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.UserInfo(ctx, i)
		if err != nil {
			return nil, err
		}
		return info, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserInfoResponse), nil
}

// todo 用户站内信
func SendMessage(ctx context.Context, i *user.SendMessageRequest) (*user.SendMessageResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		message, err := in.SendMessage(ctx, i)
		if err != nil {
			return nil, err
		}
		return message, err
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.SendMessageResponse), nil
}
func InvitationCodeGeneration(ctx *gin.Context, u *user.InvitationCodeGenerationRequest) (*user.InvitationCodeGenerationResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.InvitationCodeGeneration(ctx, u)
		if err != nil {
			return nil, err
		}
		return info, nil

	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.InvitationCodeGenerationResponse), nil
}

func FindMessage(ctx context.Context, i *user.FindMessageRequest) (*user.FindMessageResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		message, err := in.FindMessage(ctx, i)
		if err != nil {
			return nil, err
		}
		return message, err
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.FindMessageResponse), nil
}

// todo 用户修改信息
func UserModify(ctx context.Context, i *user.UserModifyRequest) (*user.UserModifyResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		modify, err := in.UserModify(ctx, i)
		if err != nil {
			return nil, err
		}
		return modify, nil

	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserModifyResponse), nil
}

// todo 实名认证
func RealName(ctx context.Context, i *user.RealNameRequest) (*user.RealNameResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		name, err := in.RealName(ctx, i)
		if err != nil {
			return nil, err
		}
		return name, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.RealNameResponse), nil
}

// todo 用户查看订单
func UserViewOrder(ctx context.Context, i *user.UserViewOrderRequest) (*user.UserViewOrderResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		order, err := in.UserViewOrder(ctx, i)
		if err != nil {
			return nil, err
		}
		return order, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserViewOrderResponse), nil
}

// todo 用户查看订单详情
func UserViewOrderInfo(ctx context.Context, i *user.UserViewOrderInfoRequest) (*user.UserViewOrderInfoResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.UserViewOrderInfo(ctx, i)
		if err != nil {
			return nil, err
		}
		return info, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserViewOrderInfoResponse), nil
}

// todo: 用户申请发票
func UserApplyInvoice(ctx context.Context, i *user.UserApplyInvoiceRequest) (*user.UserApplyInvoiceResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		invoice, err := in.UserApplyInvoice(ctx, i)
		if err != nil {
			return nil, err
		}
		return invoice, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserApplyInvoiceResponse), nil
}

// todo: 用户查看自己的发票
func UserInvoicesList(ctx context.Context, i *user.UserInvoicesListRequest) (*user.UserInvoicesListResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		invoice, err := in.UserInvoicesList(ctx, i)
		if err != nil {
			return nil, err
		}
		return invoice, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserInvoicesListResponse), nil
}

// todo:用户签到
func UserSignIn(ctx context.Context, i *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.UserSignIn(ctx, i)
		if err != nil {
			return nil, err
		}
		return info, nil
	})
	if err != nil {
		return nil, err
	}

	return userClient.(*user.UserSignInResponse), nil

}

// todo: 用户修改发票
func UserUpdateInvoice(ctx context.Context, i *user.UpdateInvoiceRequest) (*user.UpdateInvoiceResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		invoice, err := in.UpdateInvoice(ctx, i)
		if err != nil {
			return nil, err
		}
		return invoice, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UpdateInvoiceResponse), nil
}

// todo:用户补签
func UserReSigning(ctx context.Context, i *user.UserReSigningRequest) (*user.UserReSigningResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.UserReSigning(ctx, i)
		if err != nil {
			return nil, err
		}
		return info, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserReSigningResponse), nil
}

// todo 展示会员
func UserLevelInfo(ctx context.Context, i *user.UserLevelInfoRequest) (*user.UserLevelInfoResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		info, err := in.UserLevelInfo(ctx, i)
		if err != nil {
			return nil, err
		}
		return info, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserLevelInfoResponse), nil
}

// todo 展示会员权益
func UserLevelRights(ctx context.Context, i *user.UserLevelRightsRequest) (*user.UserLevelRightsResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		rights, err := in.UserLevelRights(ctx, i)
		if err != nil {
			return nil, err
		}
		return rights, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserLevelRightsResponse), nil
}

// todo 用户使用权益展示
func UserUseRights(ctx context.Context, i *user.UserUseRightsRequest) (*user.UserUseRightsResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		rights, err := in.UserUseRights(ctx, i)
		if err != nil {
			return nil, err
		}
		return rights, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.UserUseRightsResponse), nil
}

// TODO: 用户领取优惠卷
func UserReceiveCoupon(ctx context.Context, i *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		coupon, err := in.UserReceiveCoupon(ctx, i)
		if err != nil {
			return nil, err
		}
		return coupon, nil
	})
	if err != nil {
		return nil, err
	}

	return userClient.(*user.UserReceiveCouponResponse), nil
}

// todo 会员分添加记录展示
func ShowMemberPoints(ctx context.Context, i *user.ShowMemberPointsRequest) (*user.ShowMemberPointsResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		points, err := in.ShowMemberPoints(ctx, i)
		if err != nil {
			return nil, err
		}
		return points, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.ShowMemberPointsResponse), nil
}

// TODO 佣金排行榜
func CommissionRank(ctx context.Context, i *user.CommissionListRequest) (*user.CommissionListResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		rank, err := in.CommissionList(ctx, i)
		if err != nil {
			return nil, err
		}
		return rank, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.CommissionListResponse), nil
}

// todo 邀请用户排行榜
func InviteRank(ctx context.Context, i *user.InviteUserListRequest) (*user.InviteUserListResponse, error) {
	userClient, err := client.UserClient(ctx, func(ctx context.Context, in user.UserClient) (interface{}, error) {
		rank, err := in.InviteUserList(ctx, i)
		if err != nil {
			return nil, err
		}
		return rank, nil
	})
	if err != nil {
		return nil, err
	}
	return userClient.(*user.InviteUserListResponse), nil
}
