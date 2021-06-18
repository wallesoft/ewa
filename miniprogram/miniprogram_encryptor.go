package miniprogram

import (
	"strings"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
)

//签名校验 @see https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func VerifySignature(raw string, key string, signature string) bool {
	return gsha1.Encrypt(strings.Join([]string{raw, key}, "")) == signature
}

//小程序敏感信息解密 @see https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
func DecryptData(encryptedData []byte, sessionKey string, iv string) (*gjson.Json, error) {
	plainText, err := encryptor.Decrypt(gbase64.MustDecode(encryptedData), gbase64.MustDecode(gconv.Bytes(sessionKey)), gbase64.MustDecode(gconv.Bytes(iv)))
	if err != nil {
		return nil, encryptor.NewError(encryptor.ERROR_DECRYPT_AES, err.Error())
	}
	// aes-128-cbc
	decode, err := encryptor.PKCS7Unpad(plainText, 16)
	if err != nil {
		return nil, err
	}
	return gjson.New(decode), nil
}
