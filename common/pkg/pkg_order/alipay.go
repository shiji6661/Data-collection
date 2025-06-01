package pkg_order

import (
	"common/appconfig"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type Pay interface {
	Pay1(Subject, OutTradeNo, TotalAmount string) string
}
type AliPay struct {
	AppId      string
	PrivateKey string
	NotifyURL  string
	ReturnURL  string
}

func NewAliPay() *AliPay {
	ali := appconfig.NaCos.ALiPay
	return &AliPay{
		AppId:      ali.AppId,
		PrivateKey: ali.PrivateKey,
		NotifyURL:  " http://114e286f.r32.cpolar.top/order/status/update",
		ReturnURL:  ali.ReturnUrl,
	}
}

func (a *AliPay) Pay1(Subject, OutTradeNo, TotalAmount string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyURL
	p.ReturnURL = a.ReturnURL
	p.Subject = Subject         //"标题"
	p.OutTradeNo = OutTradeNo   //"传递一个唯一单号"
	p.TotalAmount = TotalAmount //"10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}
