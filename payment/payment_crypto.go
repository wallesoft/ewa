package payment

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"fmt"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/net/ghttp"
)

const (
	KEY_LENGTH_BYTE      = 32
	AUTH_TAG_LENGTH_BYTE = 16
)

//aes-256-gcm
func (p *Payment) GCMDencryter(associateData, cipherText, nonce string) ([]byte, error) {
	key := []byte(p.config.ApiV3Key)
	if len(key) != KEY_LENGTH_BYTE {
		panic("无效的ApiV3Key，长度应该为32字节")
	}

	cipherData := gbase64.MustDecodeString(cipherText)
	if len(cipherData) <= AUTH_TAG_LENGTH_BYTE {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, []byte(nonce), cipherData, []byte(associateData))
	return plaintext, err
}

//应答及回调验签
func (p *Payment) VerifySignature(response *ghttp.ClientResponse) {

	serialNo := response.Header.Get("Wechatpay-Serial")

	signatureStr := p.getSignatureStr(response.Header.Get("Wechatpay-Timestamp"), response.Header.Get("Wechatpay-Nonce"), response.ReadAllString())
	signature := gbase64.MustDecodeString(response.Header.Get("Wechatpay-Signature"))
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(signatureStr)
	hashed := h.Sum(nil)
	//证书的问题
	ok := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signature)

}

func (p *Payment) getSignatureStr(timestamp, nonce, body string) []byte {
	return gvar.New(fmt.Sprintf("%s\n%s\n%s\n", timestamp, nonce, body)).Bytes()
}

// func (p *Payment) getRsaEncrypt(data []byte, )
