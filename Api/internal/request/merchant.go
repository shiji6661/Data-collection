package request

type MerchantRegisterRequest struct {
	UserName       string `form:"user_name" binding:"required"`
	UserPassword   string `form:"user_password" binding:"required"`
	MerchantPhone  string `form:"merchant_phone" binding:"required"`
	MerchantEmail  string `form:"merchant_email" binding:"required"`
	MerchantAvatar string `form:"merchant_avatar"`
}

// MerchantLoginRequest 商家登录请求
type MerchantLoginRequest struct {
	LoginType        int64  `form:"login_type" binding:"required"`
	MerchantName     string `form:"merchant_name"`
	MerchantPassword string `form:"merchant_password"`
	MerchantPhone    string `form:"merchant_phone"`
	MerchantEmail    string `form:"merchant_email"`
	Code             string `form:"code"`
}

// SendSmsRequest 发送短信请求
type SendSmsRequest struct {
	Phone  string `form:"phone" binding:"required"`
	Source string `form:"source" binding:"required"`
}

// SendEmail 发送邮箱请求
type SendEmailRequest struct {
	Email string `form:"email" binding:"required"`
}
type ParseCodeReq struct {
	FilePath string `form:"file_path" binding:"required"`
}
