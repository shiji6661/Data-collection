package trigger

import (
	"Api/internal/handler"
	"Api/internal/request"
	"Api/internal/response"
	"github.com/gin-gonic/gin"
	"merchant_srv/proto_merchant/merchant"
)

// 商家注册
func MerchantRegister(c *gin.Context) {
	var data request.MerchantRegisterRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	register, err := handler.MerchantRegister(c, &merchant.MerChantRegisterRequest{
		UserName:       data.UserName,
		UserPassword:   data.UserPassword,
		MerchantPhone:  data.MerchantPhone,
		MerchantEmail:  data.MerchantEmail,
		MerchantAvatar: data.MerchantAvatar,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// 商家登录
func MerchantLogin(c *gin.Context) {
	var data request.MerchantLoginRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	register, err := handler.MerchantLogin(c, &merchant.MerchantLoginRequest{
		LoginType:     merchant.MerchantLoginType(data.LoginType),
		UserName:      data.MerchantName,
		UserPassword:  data.MerchantPassword,
		MerchantPhone: data.MerchantPhone,
		MerchantEmail: data.MerchantEmail,
		Code:          data.Code,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, register)
}

// 商家端短信发送
func MerchantSms(c *gin.Context) {
	var data request.SendSmsRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	sms, err := handler.MerchantSms(c, &merchant.SendSmsRequest{
		Phone:  data.Phone,
		Source: data.Source,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, sms)
}

// 商家端短信发送
func MerchantEmail(c *gin.Context) {
	var data request.SendEmailRequest
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	sms, err := handler.MerchantEmail(c, &merchant.SendEmailRequest{
		Email: data.Email,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, sms)
}

// 商家核销二维码
func MerchantParse(c *gin.Context) {
	var data request.ParseCodeReq
	if err := c.ShouldBind(&data); err != nil {
		response.ResponseError400(c, err.Error())
		return
	}
	userId := c.GetUint("userId")
	parse, err := handler.MerchantParse(c, &merchant.ParseCodeRequest{
		MerId:    int64(userId),
		Filepath: data.FilePath,
	})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, parse)
}

// 按日统计店铺数据
func StatisticsStoreDailyData(c *gin.Context) {
	userId := c.GetUint("userId")
	parse, err := handler.StatisticsStoreDailyData(c, &merchant.StatisticsStoreDailyDataRequest{MerId: int64(userId)})
	if err != nil {
		response.ResponseError(c, err.Error())
		return
	}
	response.ResponseSuccess(c, parse)
}
