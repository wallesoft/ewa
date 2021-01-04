package payment

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

type Client struct {
	*ghttp.Client
	payment *Payment
	// Logger  *log.Logger
	BaseUri string
}

const (
	AUTH_TYPE = "WECHATPAY2-SHA256-RSA2048"
)

//Request
// func (c *Client) Request(endpoint string, method string, data []byte) {

// }
func (c *Client) RequestJson(method string, endpoint string, data ...interface{}) {
	body := ""
	if len(data) > 0 {
		switch data[0].(type) {
		case string, []byte:
			body = gconv.String(data[0])
		default:
			if b, err := json.Marshal(data[0]); err != nil {
				c.payment.Logger.Errorf("Request json marshal err: %s", err.Error())
			} else {
				body = gconv.UnsafeBytesToStr(b)
			}
		}
	}
	authorization := c.getAuthorization(method, endpoint, body)
	response, err := c.Header(map[string]string{
		"Authorization": authorization,
	}).ContentJson().DoRequest(method, c.getUri(endpoint), data...)
	if err != nil {
		c.handleErrorLog(err, response.Raw())
	}
	g.Dump("Error.............", err)
	g.Dump(response)
}

func (c *Client) getUri(endpoint string) string {
	return c.BaseUri + endpoint
}

func (c *Client) getAuthorization(method string, endpoint string, body string) string {
	timestamp := gtime.TimestampStr()
	nonce := grand.S(32)
	signature := c.getSignature(strings.ToUpper(method), endpoint, nonce, timestamp, body)
	return fmt.Sprintf("%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%s\",serial_no=\"%s\",signature=\"%s\"", AUTH_TYPE, c.payment.config.MchID, nonce, timestamp, c.payment.config.SerialNo, signature)
}

func (c *Client) getSignature(method string, endpoint string, nonce string, timestamp string, body string) string {
	// timestamp := gtime.TimestampStr()
	// nonce := grand.S(32)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", method, endpoint, timestamp, nonce, body)
	g.Dump(message)
	signature, err := c.rsaEncrypt(gvar.New(message).Bytes())
	if err != nil {
		c.payment.Logger.Errorf("client signature error: method %s,endpoint %s", method, endpoint)
	}
	return signature
}

func (c *Client) rsaEncrypt(originData []byte) (string, error) {
	g.Dump("aaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	h := crypto.Hash.New(crypto.SHA256)
	g.Dump("cccccccccccccccccccccccc")
	h.Write(originData)
	g.Dump("bbabababababa")
	hashed := h.Sum(nil)
	g.Dump("890-00000000", c.payment.config.PrivateCer)
	signedData, err := rsa.SignPKCS1v15(rand.Reader, c.payment.config.PrivateCer.(*rsa.PrivateKey), crypto.SHA256, hashed)
	// signedData, err := rsa.SignPKCS1v15(rand.Reader, this.Priv.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

func (c *Client) RsaDecrypt(ciphertext string) (string, error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, c.payment.config.PrivateCer.(*rsa.PrivateKey), cipherdata, nil)
	if err != nil {
		c.payment.Logger.Errorf("Error from decryption: %s\n", err)
		return "", err
	}
	return string(plaintext), nil
}

func (c *Client) handleAccessLog(raw string) {
	if !c.payment.Logger.AccessLogEnabled {
		return
	}
	c.payment.Logger.File(c.payment.Logger.AccessLogPattern).
		Stdout(c.payment.Logger.LogStdout).
		Printf("\n=============Response Raw============\n\n %s \n ", raw)
}

func (c *Client) handleErrorLog(err error, raw string) {
	if !c.payment.Logger.ErrorLogEnabled {
		return
	}
	content := "\n\n [Error]:"
	if c.payment.Logger.ErrorStack {
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += err.Error()
		}
	} else {
		content += err.Error()
	}
	content += "\n =============Reponse Raw [err] ==============\n" + raw
	c.payment.Logger.
		File(c.payment.Logger.ErrorLogPattern).
		Stdout(c.payment.Logger.LogStdout).
		Print(content)
}

// func (this *Client) rsaEncrypt(origData []byte) (string, error) {
// 	h := crypto.Hash.New(crypto.SHA256)
// 	h.Write(origData)
// 	hashed := h.Sum(nil)
// 	// 进行rsa加密签名
// 	signedData, err := rsa.SignPKCS1v15(rand.Reader, this.Priv.(*rsa.PrivateKey), crypto.SHA256, hashed)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.StdEncoding.EncodeToString(signedData), nil
// }

// func (this *Client) RsaDecrypt(ciphertext string) (string, error) {
// 	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
// 	rng := rand.Reader

// 	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, this.Priv.(*rsa.PrivateKey), cipherdata, nil)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
// 		return "", err
// 	}
// 	return string(plaintext), nil
// }
