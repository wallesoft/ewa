package payment

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

type Client struct {
	*ghttp.Client
	payment   *Payment
	UrlValues url.Values
	// Logger  *log.Logger
	BaseUri string
}
type ClientResponse struct {
	*ghttp.ClientResponse
}

const (
	AUTH_TYPE = "WECHATPAY2-SHA256-RSA2048"
)

//RequestJson
func (c *Client) RequestJson(method string, endpoint string, data ...interface{}) *ghttp.ClientResponse {
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
	queryString, url := c.getUri(endpoint)
	authorization := c.getAuthorization(method, queryString, body)
	response, err := c.Header(map[string]string{
		"Authorization": authorization,
		"User-Agent":    "ewa-payment-client/1.0",
		"Accept":        "application/json",
	}).ContentJson().DoRequest(method, url, data...)
	if err != nil {
		c.handleErrorLog(err, response.Raw())
	}
	// g.Dump(response.StatusCode)
	if response.StatusCode == 200 || response.StatusCode == 204 {
		//response.StatusCode == 200 验签
		err := c.payment.VerifySignature(&ClientResponse{ClientResponse: response})
		if err == nil {
			c.handleAccessLog(response.Raw())

		} else {
			c.handleErrorLog(err, response.Raw())

		}
		return response
	}

	c.handleErrorLog(errors.New("payment.v3.请求错误:"+response.Status), response.Raw())
	return response
}

func (c *Client) getUri(endpoint string) (query, urlString string) {
	param := url.Values{}
	// var url string
	if c.UrlValues != nil {
		param = c.UrlValues
	}
	query = endpoint
	// url =  c.BaseUri + endpoint

	if len(param) > 0 {
		query = query + "?" + param.Encode()
	}

	urlString = c.BaseUri + query
	return
}

func (c *Client) getAuthorization(method string, endpoint string, body string) string {
	timestamp := gtime.TimestampStr()
	nonce := grand.S(32)
	signature := c.getSignature(strings.ToUpper(method), endpoint, nonce, timestamp, body)
	return fmt.Sprintf("%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%s\",serial_no=\"%s\",signature=\"%s\"", AUTH_TYPE, c.payment.config.MchID, nonce, timestamp, c.payment.config.SerialNo, signature)
}

func (c *Client) getSignature(method string, endpoint string, nonce string, timestamp string, body string) string {

	message := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", method, endpoint, timestamp, nonce, body)
	signature, err := c.rsaEncrypt(gvar.New(message).Bytes())
	if err != nil {
		c.payment.Logger.Errorf("client signature error: method %s,endpoint %s", method, endpoint)
	}
	return signature
}

func (c *Client) rsaEncrypt(originData []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(originData)
	hashed := h.Sum(nil)
	signedData, err := rsa.SignPKCS1v15(rand.Reader, c.payment.config.PrivateCer.(*rsa.PrivateKey), crypto.SHA256, hashed)
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
