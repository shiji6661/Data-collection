package trigger

import (
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"Api/pkg"
	pkg2 "common/pkg/pkg_merchant"
	user2 "common/utils/user"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"weikang/Data-collection/srv/user_srv/proto_user/user"
)

// todo:用户注册
func UserRegister(c *gin.Context) {
	var data request.UserRegister
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	var id uint64
	if data.InviteCode != "" {
		var err error
		id, err = user2.ParseInviteCode(data.InviteCode)
		if err != nil {
			response.ResponseError(c, err.Error())
			return
		}
	}
	fmt.Println(data.InviteCode)
	fmt.Println(11111)
	fmt.Println(id)
	fmt.Println(11111111111)
	register, err := handler.UserRegister(c, &user.UserRegisterRequest{
		UserName:     data.UserName,
		UserPassword: data.UserPassword,
		UserPhone:    data.UserPhone,
		UserInviteId: int64(id),
		InviteCode:   data.InviteCode,
	})
	fmt.Println(&user.UserRegisterRequest{
		UserName:     data.UserName,
		UserPassword: data.UserPassword,
		UserPhone:    data.UserPhone,
		UserInviteId: int64(id),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// todo 短信发送
func SendSms(c *gin.Context) {
	var data request.SendSms
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	sms, err := handler.SendSms(c, &user.SendSmsRequest{
		Mobile: data.Mobile,
		Source: data.Source,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, sms)
}

// todo 邮件发送
func SendEmail(c *gin.Context) {
	var data request.SendEmail
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	email, err := handler.SendEmail(c, &user.SendEmailRequest{UserEmail: data.UserEmail})
	if err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	response.ResponseSuccess(c, email)
}

// todo 用户登录
func UserLogin(c *gin.Context) {
	var data request.UserLogin
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	login, err := handler.UserLogin(c, &user.UserLoginRequest{
		Login_Type:   user.LoginType(data.LoginType),
		UserName:     data.UserName,
		UserPassword: data.UserPassword,
		UserPhone:    data.UserPhone,
		MobileCode:   data.MobileCode,
		UserEmail:    data.UserEmail,
		EmailCode:    data.EmailCode,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	claims := pkg.CustomClaims{ID: uint(login.Greet)}
	token, err := pkg.NewJWT("2209A").CreateToken(claims)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, token)
}

// todo 用户详情
func UserInfo(c *gin.Context) {
	id := c.GetUint("userId")
	info, err := handler.UserInfo(c, &user.UserInfoRequest{
		Id: int64(id),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)
}

func SendMessage(c *gin.Context) {
	var data request.SendMessage
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	message, err := handler.SendMessage(c, &user.SendMessageRequest{
		SendId:     int64(id),
		ReceiverId: data.ReceiverId,
		Context:    data.Context,
	})
	if err != nil {
		return
	}
	response.ResponseSuccess(c, message)
}

func FindMessage(c *gin.Context) {
	var data request.FindMessage
	err := c.ShouldBind(&data)
	if err != nil {
		response.ResponseError(c, err.Error())
	}
	id := c.GetUint("userId")
	message, err := handler.FindMessage(c, &user.FindMessageRequest{ReceiverId: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, message)

}
func InvitationCodeGeneration(c *gin.Context) {
	id := c.GetUint("userId")
	info, err := handler.InvitationCodeGeneration(c, &user.InvitationCodeGenerationRequest{
		UserId: int64(id),
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)
}

// todo 用户修改信息
func UserModify(c *gin.Context) {
	var data request.UserModify
	if err := c.ShouldBind(&data); err != nil {

		response.ResponseError(c, err.Error())
		return
	}
	id := c.GetUint("userId")

	message, err := handler.FindMessage(c, &user.FindMessageRequest{ReceiverId: int64(id)})
	if err != nil {
		return
	}
	response.ResponseSuccess(c, message)

	modify, err := handler.UserModify(c, &user.UserModifyRequest{
		Id:                  int64(id),
		UserName:            data.UserName,
		UserPassword:        data.UserPassword,
		Password:            data.Password,
		UserPhone:           data.UserPhone,
		UserEmail:           data.UserEmail,
		UserProvince:        data.UserProvince,
		UserCity:            data.UserCity,
		UserCounty:          data.UserCounty,
		UserDetailedAddress: data.UserDetailedAddress,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, modify)
}

// todo 文件上传
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, errors.New("文件获取失败"))
		return
	}
	ext := path.Ext(file.Filename)
	result := map[string]bool{
		".jpg": true,
		".png": true,
	}

	if !result[ext] {
		response.ResponseError(c, "图片格式应为.jpg或.png格式")
		return
	}

	if file.Size > 3*1024*1024 {
		response.ResponseError(c, "图片大小不能超过3MB")
		return
	}

	dst := "D:\\gocode\\shopps\\shopping\\common\\pkg\\pkg_merchant\\upload\\" + file.Filename

	c.SaveUploadedFile(file, dst)
	go pkg2.AliYunUpload(dst, file.Filename)

	response.ResponseSuccess(c, "文件上传成功"+pkg2.GetUrl(file.Filename))
}

// todo 实名认证
func RealName(c *gin.Context) {
	var data request.RealName
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	name, err := handler.RealName(c, &user.RealNameRequest{
		Uid:  int64(id),
		Name: data.Name,
		Card: data.Card,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, name)
}

// todo 用户查看订单
func UserViewOrder(c *gin.Context) {
	id := c.GetUint("userId")
	order, err := handler.UserViewOrder(c, &user.UserViewOrderRequest{Uid: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, order)
}

// todo 用户查看订单详情
func UserViewOrderInfo(c *gin.Context) {
	var data request.UserViewOrderInfo
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	info, err := handler.UserViewOrderInfo(c, &user.UserViewOrderInfoRequest{
		Uid:     int64(id),
		OrderId: data.OrderId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)

}

// todo 用户申请发票
func UserApplyInvoice(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.UserApplyInvoice
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	invoice, err := handler.UserApplyInvoice(c, &user.UserApplyInvoiceRequest{
		Userid:    int64(userid),
		OrderId:   data.OrderId,
		TitleType: data.TitleType,
		Title:     data.Title,
		TaxId:     data.TaxId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, invoice)
}

// todo:用户查看自己的发票
func UserInvoicesList(c *gin.Context) {
	userid := c.GetUint("userId")
	list, err := handler.UserInvoicesList(c, &user.UserInvoicesListRequest{Userid: int64(userid)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, list)
}

// todo:用户签到
func UserSignIn(c *gin.Context) {
	id := c.GetUint("userId")
	info, err := handler.UserSignIn(c, &user.UserSignInRequest{Uid: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)
}

// todo:用户修改发票
func UserUpdateInvoice(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.UpdateInvoiceRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	list, err := handler.UserUpdateInvoice(c, &user.UpdateInvoiceRequest{
		Userid:    int64(userid),
		InvoiceId: data.InvoiceId,
		Type:      data.Type,
		TitleType: data.TitleType,
		Title:     data.Title,
		TaxId:     data.TaxId,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, list)
}

// todo:用户补签
func UserReSigning(c *gin.Context) {
	id := c.GetUint("userId")
	var data request.UserReSigning
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, "请求参数错误")
		return
	}
	info, err := handler.UserReSigning(c, &user.UserReSigningRequest{
		Uid:      int64(id),
		SignDate: data.SignDate,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)
}

// todo 展示会员
func UserLevelInfo(c *gin.Context) {
	id := c.GetUint("userId")
	info, err := handler.UserLevelInfo(c, &user.UserLevelInfoRequest{Uid: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, info)
}

// todo 展示会员权益
func UserLevelRights(c *gin.Context) {
	var data request.UserLevelRights
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	id := c.GetUint("userId")
	rights, err := handler.UserLevelRights(c, &user.UserLevelRightsRequest{
		Uid:  int64(id),
		ULid: data.ULid,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}

	response.ResponseSuccess(c, rights)
}

// todo 用户使用权益展示
func UserUseRights(c *gin.Context) {
	id := c.GetUint("userId")
	rights, err := handler.UserUseRights(c, &user.UserUseRightsRequest{Uid: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, rights)
}

// todo:用户领取优惠卷
func UserReceive(c *gin.Context) {
	userid := c.GetUint("userId")
	var data request.UserReceiceCoupon
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	coupon, err := handler.UserReceiveCoupon(c, &user.UserReceiveCouponRequest{
		Userid: int64(userid),
		CSId:   data.Cid,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, coupon)
}

// todo 会员分添加记录展示
func ShowMemberPoints(c *gin.Context) {
	id := c.GetUint("userId")
	points, err := handler.ShowMemberPoints(c, &user.ShowMemberPointsRequest{Userid: int64(id)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, points)
}

// todo 佣金排行榜
func CommissionList(c *gin.Context) {
	var data request.CommissionListReq
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	rank, err := handler.CommissionRank(c, &user.CommissionListRequest{
		Page: data.Page,
		Size: data.Size,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, rank)
}

// TODO 邀请用户排行榜
func InviteUserList(c *gin.Context) {
	rank, err := handler.InviteRank(c, &user.InviteUserListRequest{})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, rank)
}
