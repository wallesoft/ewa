package payment

import "crypto/x509"

const (
	BASE_API_URI = "https://api.mch.weixin.qq.com"
)

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

	//v3接口
	SerialNo    string            //商户证书编号
	PublicCer   *x509.Certificate //商户证书公钥
	PrivateCer  interface{}       //商户证书私钥
	PFSerialNo  string            //平台证书编号
	PFPublicCer *x509.Certificate //平台证书公钥
}

func (p *Payment) GetClient(endpoint string, method string) *Client {
	return &Client{
		BaseUri: p.getBaseUri(),
	}
}

func (p *Payment) getBaseUri() string {
	return "https://api.mch.weixin.qq.com"
}
