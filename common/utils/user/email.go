package user

import (
	"common/appconfig"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(toEmail string, context string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = appconfig.NaCos.QqEmail.Qq
	// 设置接收方的邮箱
	e.To = []string{toEmail}
	//设置主题
	e.Subject = "验证码"
	//设置文件发送的内容

	e.Text = []byte(context)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", appconfig.NaCos.QqEmail.Qq, appconfig.NaCos.QqEmail.Password, "smtp.qq.com"))
	if err != nil {
		return err
	}
	return nil
}
