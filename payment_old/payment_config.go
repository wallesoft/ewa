package payment

//Config
type Config struct {
	AppID     string `json:"app_id"`
	MchID     string `json:"mch_id"`
	Key       string `json:"key"`        //API 秘钥
	CertPath  string `json:"cert_path"`  //API 证书路径 绝对路径！！！
	KeyPath   string `json:"key_path"`   //绝对路径！！！
	NotifyUrl string `json:"notify_url"` //默认回调地址
	Sandbox   bool   `json:"sandbox"`    //沙河模式

	//服务商参数
	SubMchID string `json:"sub_mch_id"`
	SubAppID string `json:"sub_appid"`
}
