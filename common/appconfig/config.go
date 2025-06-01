package appconfig

//TODO:json格式文件

var NaCos NaCosJson

// NaCos配置文件
type NaCosJson struct {
	Mysql struct {
		User     string `json:"User"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Database string `json:"Database"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Elasticsearch struct {
		Host string `json:"host"`
	} `json:"elasticsearch"`
	Group struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"group"`
	Consul struct {
		Name string `json:"Name"`
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"consul"`
	Rabbitmq struct {
		Address string `json:"address"`
	} `json:"rabbitmq"`
	ALiYun struct {
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		Address         string `json:"address"`
	} `json:"ALiYun"`
	ALiPay struct {
		PrivateKey string `json:"privateKey"`
		AppId      string `json:"appId"`
		NotifyUrl  string `json:"notifyUrl"`
		ReturnUrl  string `json:"returnUrl"`
	} `json:"ALiPay"`
	TenXunYun struct {
		SecretID  string `json:"SecretID"`
		SecretKey string `json:"SecretKey"`
	} `json:"tenXunYun"`
	QqEmail struct {
		Qq       string `json:"qq"`
		Password string `json:"password"`
	} `json:"qqEmail"`
}
