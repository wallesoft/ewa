package payment

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"gitee.com/wallesoft/ewa/kernel/log"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

const (
	BASE_API_URI = "https://api.mch.weixin.qq.com"
)

//Config
type Config struct {
	AppID string `json:"app_id"` //公众号
	MchID string `json:"mch_id"` //商户号
	//v2接口
	Key      string `json:"key"`       //API 秘钥
	CertPath string `json:"cert_path"` //API 证书路径 绝对路径！！！
	KeyPath  string `json:"key_path"`  //绝对路径！！！
	Sandbox  bool   `json:"sandbox"`   //沙盒模式
	//v3接口
	SerialNo    string            `json:"serial_no"`         //商户证书编号
	PriCertPath string            `json:"private_cert_path"` //私钥证书路径 绝对路径！！！
	PubCertPath string            `json:"public_cert_Path"`  //公钥证书路径 绝对路径！！！
	PublicCer   *x509.Certificate //商户证书公钥
	PrivateCer  interface{}       //商户证书私钥

	PFSerialNo     string            //平台证书编号
	PFPublicCer    *x509.Certificate //平台证书公钥
	PFCertSavePath string            //平台证书保存路径 绝对路径！！！
	PFCertPrefix   string            //平台证书保存前缀
	ApiV3Key       string            //apiv3 秘钥

	//公用
	NotifyUrl string `json:"notify_url"` //默认回调地址

	//服务商参数
	SubMchID string `json:"sub_mch_id"`
	SubAppID string `json:"sub_appid"`
	//others
	Logger  *log.Logger
	LogPath string
}

func (p *Payment) setConfig(config Config, compatible ...bool) Config {
	var err error
	if config.MchID == "" {
		panic("商户号无效 MchiID: nil")
	}
	if config.SerialNo == "" {
		panic("证书编号无效 SerialNo: nil")
	}
	if priCertData := gfile.GetBytes(config.PriCertPath); priCertData == nil {
		panic(fmt.Sprintf("私钥证书读取失败,config PriCertPath: %s", config.PriCertPath))
	} else {
		if block, _ := pem.Decode(priCertData); block == nil || block.Type != "PRIVATE KEY" {
			panic("私钥PEM解码失败")
		} else {

			config.PrivateCer, err = x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				panic(err.Error())
			}
		}
	}
	if pubCertData := gfile.GetBytes(config.PubCertPath); pubCertData == nil {
		panic(fmt.Sprintf("公钥证书读取失败,config PubCertPath: %s", config.PubCertPath))
	} else {
		if block, _ := pem.Decode(pubCertData); block == nil || block.Type != "CERTIFICATE" {
			panic("公钥PEM解码失败")
		} else {
			config.PublicCer, err = x509.ParseCertificate(block.Bytes)
			if err != nil {
				panic(err.Error())
			}
		}
	}
	if config.PFCertPrefix == "" {
		config.PFCertPrefix = "wechatpay_"
	}
	if config.PFCertSavePath == "" {
		config.PFCertSavePath = "/etc/wechatpay/"
	}
	return config
}
func (p *Payment) getClient() *Client {
	return &Client{
		Client:  ghttp.NewClient(),
		BaseUri: p.getBaseUri(),
		payment: p,
	}
}

func (p *Payment) getBaseUri() string {
	return "https://api.mch.weixin.qq.com"
}
