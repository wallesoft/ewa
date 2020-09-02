package server

type Config interface {
	AppID          string `json:"app_id"`           //appid
	AppSecret      string `json:"app_secret"`       //appsecret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey
}
