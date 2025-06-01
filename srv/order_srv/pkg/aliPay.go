package pkg

import (
	"common/appconfig"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type Pay interface {
	Pay(subject, outTradeNo, totalAmount string) string
}

type ALiPay struct {
	PrivateKey string
	AppId      string
	NotifyURL  string
	ReturnURL  string
}

func NewAliPay() *ALiPay {
	data := appconfig.NaCos.ALiPay
	return &ALiPay{
		PrivateKey: data.PrivateKey,
		AppId:      data.AppId,
		NotifyURL:  "https://63381bb3.r39.cpolar.top/order/call/back",
		ReturnURL:  data.ReturnUrl,
	}
}

func (a *ALiPay) Pay(subject, outTradeNo, totalAmount string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyURL
	p.ReturnURL = a.ReturnURL
	p.Subject = subject
	p.OutTradeNo = outTradeNo
	p.TotalAmount = totalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
