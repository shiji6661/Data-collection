package pkg

import (
	"common/appconfig"
	"net/smtp"
)

func SendEmail(toEmail string, context string) error {
	// QQ 邮箱的 SMTP 服务器地址
	smtpServer := "smtp.qq.com"
	// QQ 邮箱 SMTP 服务器的端口，使用 587 端口需开启 TLS
	smtpPort := "587"
	// 发件人的 QQ 邮箱地址
	fromEmail := appconfig.NaCos.QqEmail.Qq
	// 发件人 QQ 邮箱的授权码，不是邮箱登录密码
	authCode := appconfig.NaCos.QqEmail.Password

	// 构建邮件消息
	message := []byte("To: " + toEmail + "\r\n" +
		"From: " + fromEmail + "\r\n" +
		"Subject: " + context + "\r\n")

	// 创建认证信息
	auth := smtp.PlainAuth("", fromEmail, authCode, smtpServer)

	// 发送邮件
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, fromEmail, []string{toEmail}, message)
	if err != nil {
		return err
	}
	return nil
}
