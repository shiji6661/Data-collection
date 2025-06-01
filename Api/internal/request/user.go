package request

// TODO 用户注册
type UserRegister struct {
	UserName     string `form:"user_name" json:"user_name"  binding:"required"`
	UserPhone    string `form:"user_phone" json:"user_phone"  binding:"required"`
	UserPassword string `form:"user_password" json:"user_password"  binding:"required"`
	InviteCode   string `form:"invite_code" json:"invite_code"  `
}

// todo 短信发送
type SendSms struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Source string `json:"source" form:"source" binding:"required"`
}

// todo 邮件发送
type SendEmail struct {
	UserEmail string `json:"user_email" form:"user_email" binding:"required"`
}

// todo 用户登录
type UserLogin struct {
	LoginType    int64  `json:"login_type" form:"login_type"`
	UserName     string `json:"user_name" form:"user_name"`
	UserPassword string `json:"user_password" form:"user_password"`
	UserPhone    string `json:"user_phone" form:"user_phone"`
	MobileCode   string `json:"mobile_code" form:"mobile_code"`
	UserEmail    string `json:"user_email" form:"user_email"`
	EmailCode    string `json:"email_code" form:"email_code"`
}

// todo 发送站内信

type SendMessage struct {
	ReceiverId int64  `json:"receiver_id" binding:"required"`
	Context    string `json:"context" binding:"required"`
}

type FindMessage struct {
}

// todo 用户修改信息
type UserModify struct {
	UserName            string `json:"user_name" form:"user_name" binding:"omitempty"`                         //用户名
	UserPassword        string `json:"user_password" form:"user_password" binding:"omitempty"`                 //用户密码
	Password            string `json:"password" form:"password" binding:"omitempty"`                           //确认密码
	UserPhone           string `json:"user_phone" form:"user_phone" binding:"omitempty"`                       //用户手机号码
	UserEmail           string `json:"user_email" form:"user_email" binding:"omitempty"`                       //用户邮箱
	UserProvince        string `json:"user_province" form:"user_province" binding:"omitempty"`                 //用户所在省份
	UserCity            string `json:"user_city" form:"user_city" binding:"omitempty"`                         //用户所在城市
	UserCounty          string `json:"user_county" form:"user_county" binding:"omitempty"`                     //用户所在国家
	UserDetailedAddress string `json:"user_detailed_address" form:"user_detailed_address" binding:"omitempty"` //用户详细地址
}

// todo 实名认证
type RealName struct {
	Name string `json:"name" form:"name" binding:"required"` //用户真实姓名
	Card string `json:"card" form:"card" binding:"required"` //身份证号码
}

// todo 用户查看订单详情
type UserViewOrderInfo struct {
	OrderId int64 `json:"order_id" form:"order_id" binding:"required"` //订单id
}

// todo 展示会员权益
type UserLevelRights struct {
	ULid int64 `json:"u_lid" form:"u_lid" binding:"required"` //会员id
}

// todo 用户申请发票
type UserApplyInvoice struct {
	OrderId   int64  `form:"order_id" binding:"required"`   //订单id
	TitleType int64  `form:"title_type" binding:"required"` //抬头类型： 1-个人 2-企业
	Title     string `form:"title" binding:"required"`      //抬头名称
	TaxId     string `form:"tax_id" binding:"required"`     //纳税人识别号
}

// todo 用户补签
type UserReSigning struct {
	SignDate string `json:"sign_date" form:"sign_date" binding:"required"` //会员id
}

// todo:用户修改发票
type UpdateInvoiceRequest struct {
	InvoiceId int64  `form:"invoice_id" binding:"required"` //发票id
	Type      int64  `form:"type" binding:"required"`       //抬头类型： 1-个人 2-企业
	TitleType int64  `form:"title_type" binding:"required"` //抬头类型： 1-个人 2-企业
	Title     string `form:"title" binding:"required"`      //抬头名称
	TaxId     string `form:"tax_id" binding:"required"`     //纳税人识别号
}

// TODO:用户领取优惠卷
type UserReceiceCoupon struct {
	Cid int64 `form:"c_id" binding:"required"`
}

// TODO 佣金排行榜
type CommissionListReq struct {
	Page int64 `form:"page" binding:"required"`
	Size int64 `form:"size" binding:"required"`
}
