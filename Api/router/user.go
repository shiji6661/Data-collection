package router

import (
	"Api/internal/middleware"
	"Api/internal/trigger"
	"Api/pkg"
	"github.com/gin-gonic/gin"
)

func LoadUser(r *gin.Engine) {
	r.Use(middleware.Logger())
	user := r.Group("/user")
	{
		user.POST("/register", trigger.UserRegister) // todo:用户账号密码电话号码注册

		user.POST("/login", trigger.UserLogin) // todo:用户账号密码短信邮箱验证码登录

		user.POST("/send/sms", trigger.SendSms) // todo:短信验证码

		user.POST("/send/email", trigger.SendEmail) // todo:邮件发送

		user.GET("commission/list", trigger.CommissionList) // todo:佣金排行榜

		user.GET("invite/rank", trigger.InviteUserList) // todo:邀请用户排行榜

		user.Use(pkg.JWTAuth("2209A"))
		user.POST("/sendMessage", trigger.SendMessage) //todo:用户发送站内信

		user.POST("/findMessage", trigger.FindMessage) //todo:用户接收站内信

		user.POST("/invitation/code", trigger.InvitationCodeGeneration) //todo:用户邀请码生成

		user.POST("/info", trigger.UserInfo) //todo:用户详情

		user.POST("/modify", trigger.UserModify) //todo:用户修改信息

		user.POST("/uploadFile", trigger.UploadFile) //todo:文件上传

		user.POST("/realName", trigger.RealName) //todo:实名认证

		user.GET("/viewOrder", trigger.UserViewOrder) //todo:用户查看订单

		user.POST("/viewOrderInfo", trigger.UserViewOrderInfo) //todo:用户查看订单详情

		user.POST("/apply/Invoice", trigger.UserApplyInvoice) //todo:用户申请发票

		user.POST("/invoices/list", trigger.UserInvoicesList) //todo:用户查看自己的发票

		user.POST("/update/Invoice", trigger.UserUpdateInvoice) //todo:用户修改发票

		user.GET("/levelInfo", trigger.UserLevelInfo) //todo:展示会员

		user.POST("/levelRights", trigger.UserLevelRights) //todo:展示会员权益

		user.POST("/sign/in", trigger.UserSignIn) //todo:用户签到

		user.POST("/re/sign", trigger.UserReSigning) //todo:用户补签

		user.POST("/receive/coupon", trigger.UserReceive) // todo:用户领取优惠卷

		user.GET("/useRights", trigger.UserUseRights) //todo:展示用户使用的权益

		user.GET("/showMemberPoints", trigger.ShowMemberPoints) //todo 会员分添加记录展示
	}

}
