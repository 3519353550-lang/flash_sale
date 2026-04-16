package configs

type Nacos struct {
	Host      string
	Port      int
	Namespace string
	Group     string
	DataId    string
	Username  string
	Password  string
}

type Configs struct {
	Mysql struct {
		Host     string
		Port     int64
		User     string
		Password string
		Database string
	}
	Redis struct {
		Host     string
		Port     int64
		Password string
		Database int
	}
	SendSms struct {
		Account  string
		Password string
	}
	QiNiuYun struct {
		AccessKey string
		SecretKey string
		Bucket    string
		Url       string
	}
	Elastic struct {
		Host string
	}
	AilPay struct {
		AppId      string
		PrivateKey string
		NotifyUrl  string
		ReturnUrl  string
	}
}
