package payment

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/os/gfile"
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
func (p *Payment) VerifySignature(response *ClientResponse) error {

	serialNo := response.Header.Get("Wechatpay-Serial")
	p.setPFPublicCert(serialNo)

	signatureStr := p.getSignatureStr(response.Header.Get("Wechatpay-Timestamp"), response.Header.Get("Wechatpay-Nonce"), response.ReadAllString())
	signature := gbase64.MustDecodeString(response.Header.Get("Wechatpay-Signature"))
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(signatureStr)
	hashed := h.Sum(nil)
	//证书的问题
	ok := rsa.VerifyPKCS1v15(p.config.PFPublicCer.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed, signature)
	if ok != nil {
		return rsa.ErrVerification
	}
	return nil
}

func (p *Payment) getSignatureStr(timestamp, nonce, body string) []byte {
	return gvar.New(fmt.Sprintf("%s\n%s\n%s\n", timestamp, nonce, body)).Bytes()
}
func (p *Payment) setPFPublicCert(serialNo string) {
	var err error
	p.config.PFSerialNo = serialNo
	if certData := gfile.GetBytes(p.config.PFCertSavePath + "pf_wechatpay_" + serialNo + ".pem"); certData == nil {
		panic("平台公钥读取失败")
	} else {
		if block, _ := pem.Decode(certData); block == nil || block.Type != "CERTIFICATE" {
			panic("平台公钥PEM解码失败")
		} else {
			p.config.PFPublicCer, err = x509.ParseCertificate(block.Bytes)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
