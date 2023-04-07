package payment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Client struct {
	*gclient.Client
	payment   *Payment
	UrlValues url.Values
	// Logger  *log.Logger
	BaseUri string
}
type ClientResponse struct {
	*gclient.Response
}

const (
	AUTH_TYPE = "WECHATPAY2-SHA256-RSA2048"
)

//RequestJson v3
func (c *Client) RequestJson(ctx context.Context, method string, endpoint string, data ...interface{}) *Response {
	body := ""
	if len(data) > 0 {
		switch data[0].(type) {
		case string, []byte:
			body = gconv.String(data[0])
		default:
			if b, err := json.Marshal(data[0]); err != nil {
				c.payment.Logger.Errorf(ctx, "Request json marshal err: %s", err.Error())
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
	}).ContentJson().DoRequest(ctx, method, url, data...)
	if err != nil {
		c.handleErrorLog(err, response.Raw())
	}
	debugRaw := response.Raw()
	res := &Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       response.ReadAll(),
	}

	if res.StatusCode == 200 || res.StatusCode == 204 {
		//response.StatusCode == 200
		//验签
		err := c.payment.VerifySignature(ctx, res.Header, res.Body)
		if err != nil {
			c.handleErrorLog(err, debugRaw)
		} else {
			c.handleAccessLog(debugRaw)
		}
	} else {
		c.handleErrorLog(errors.New("payment.v3.请求错误"), debugRaw)
	}

	return res
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
	signature, err := c.payment.rsaEncrypt(gvar.New(message).Bytes())
	if err != nil {
		panic(fmt.Sprintf("[Erro] client signature error: method %s,endpoint %s", method, endpoint))
		// c.payment.Logger.Errorf("client signature error: method %s,endpoint %s", method, endpoint)
	}
	return signature
}

// func (c *Client) rsaEncrypt(originData []byte) (string, error) {
// 	h := crypto.Hash.New(crypto.SHA256)
// 	h.Write(originData)
// 	hashed := h.Sum(nil)
// 	signedData, err := rsa.SignPKCS1v15(rand.Reader, c.payment.config.PrivateCer.(*rsa.PrivateKey), crypto.SHA256, hashed)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.StdEncoding.EncodeToString(signedData), nil
// }

// func (c *Client) RsaDecrypt(ciphertext string) (string, error) {
// 	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
// 	rng := rand.Reader
// 	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, c.payment.config.PrivateCer.(*rsa.PrivateKey), cipherdata, nil)
// 	if err != nil {
// 		c.payment.Logger.Errorf("Error from decryption: %s\n", err)
// 		return "", err
// 	}
// 	return string(plaintext), nil
// }

// v2 requst with cert tls
func (c *Client) RequestV2(ctx context.Context, method string, endpoint string, data ...interface{}) *Response {
	_, url := c.getUri(endpoint)
	err := c.SetTLSKeyCrt(c.payment.config.CertPath, c.payment.config.KeyPath)
	if err != nil {
		c.payment.Logger.Errorf(ctx, "set tls key crt err: %s", err.Error())
	}
	response, err := c.DoRequest(ctx, method, url, data...)

	if err != nil {
		c.handleErrorLog(err, response.Raw())
	}

	raw := gjson.New(response.Raw())
	res := &Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       response.ReadAll(),
	}

	if raw.Contains("return_code") && raw.Get("return_code").String() != "SUCCESS" {
		c.handleErrorLog(errors.New("payment.v2.请求错误"), raw.MustToXmlString())
	}
	c.handleAccessLog(raw.MustToXmlString())
	return res

}

func (c *Client) handleAccessLog(raw string) {
	if !c.payment.Logger.AccessLogEnabled {
		return
	}
	c.payment.Logger.File(c.payment.Logger.AccessLogPattern).
		Stdout(c.payment.Logger.LogStdout).
		Printf(context.TODO(), "\n=============Response Raw============\n\n %s \n ", raw)
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
		Print(context.TODO(), content)
}
