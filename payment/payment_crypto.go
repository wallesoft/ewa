package payment

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

const (
	KEY_LENGTH_BYTE      = 32
	AUTH_TAG_LENGTH_BYTE = 16
)

//aes-256-gcm
func (p *Payment) GCMDecryte(associateData, cipherText, nonce string) ([]byte, error) {
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
func (p *Payment) VerifySignature(header http.Header, body []byte) error {
	var err error
	gutil.TryCatch(func() {
		serialNo := header.Get("Wechatpay-Serial")
		p.setPFPublicCert(serialNo)
		signatureStr := p.getSignatureStr(header.Get("Wechatpay-Timestamp"), header.Get("Wechatpay-Nonce"), gvar.New(body).String())
		signature := gbase64.MustDecodeString(header.Get("Wechatpay-Signature"))
		h := crypto.Hash.New(crypto.SHA256)
		h.Write(signatureStr)
		hashed := h.Sum(nil)
		//证书的问题
		ok := rsa.VerifyPKCS1v15(p.config.PFPublicCer.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed, signature)
		if ok != nil {
			panic(rsa.ErrVerification.Error())
		}

	}, func(e error) {
		err = e
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Payment) getSignatureStr(timestamp, nonce, body string) []byte {
	return gvar.New(fmt.Sprintf("%s\n%s\n%s\n", timestamp, nonce, body)).Bytes()
}

func (p *Payment) setPFPublicCert(serialNo string) {
	var err error
	p.config.PFSerialNo = serialNo
	if certData := gfile.GetBytes(p.config.PFCertSavePath + p.config.PFCertPrefix + serialNo + ".pem"); certData == nil {
		panic(fmt.Sprintf("平台公钥读取失败,path: %s", p.config.PFCertSavePath+p.config.PFCertPrefix+serialNo+".pem"))
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

func (p *Payment) rsaEncrypt(originData []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(originData)
	hashed := h.Sum(nil)
	signedData, err := rsa.SignPKCS1v15(rand.Reader, p.config.PrivateCer.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

func (p *Payment) rsaDecrypt(ciphertext string) (string, error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, p.config.PrivateCer.(*rsa.PrivateKey), cipherdata, nil)
	if err != nil {
		// c.payment.Logger.Errorf("Error from decryption: %s\n", err)
		return "", err
	}
	return string(plaintext), nil
}

//v2排序
func (p *Payment) V2SortKey(text map[string]interface{}) string {
	tmpStrs := make([]string, 0)

	for key, val := range text {
		s := gconv.String(val)
		if s != "" {
			tmpStrs = append(tmpStrs, key+"="+s)
		}
	}
	sort.Strings(tmpStrs)
	return strings.Join(tmpStrs, "&")
}

//v2 md5签名
func (p *Payment) V2MD5(m map[string]interface{}) string {
	sortStr := p.V2SortKey(m)
	signTmp := sortStr + "&key=" + p.config.Key
	g.Dump(signTmp)
	return strings.ToUpper(gmd5.MustEncryptString(signTmp))
}

//v2加密
func (p *Payment) V2Signature(str string) string {

	h := hmac.New(sha256.New, gconv.Bytes(p.config.Key))
	h.Write(gconv.Bytes(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
