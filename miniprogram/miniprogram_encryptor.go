package miniprogram

import (
	"crypto/aes"
	"crypto/cipher"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func DecryptData(encryptedData []byte, sessionKey string, iv string) (*gjson.Json, error) {
	decoding := gbase64.MustDecode(encryptedData)
	//cipherText
	// if err != nil {
	// 	return nil, encryptor.NewError(encryptor.ERROR_BASE64_DECODE, err.Error())
	// }
	ivDecoded := gbase64.MustDecode(gconv.Bytes(iv))
	// if err != nil {
	// 	return nil, encryptor.NewError(encryptor.ERROR_BASE64_DECODE, err.Error())
	// }
	aesKey := gbase64.MustDecode(gconv.Bytes(sessionKey))
	// if err != nil {
	// 	return nil, encryptor.NewError(encryptor.ERROR_BASE64_DECODE, err.Error())
	// }
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, encryptor.NewError(encryptor.ERROR_DECRYPT_AES, err.Error())
	}
	blockModel := cipher.NewCBCDecrypter(block, ivDecoded)
	plainText := make([]byte, len(decoding))
	blockModel.CryptBlocks(plainText, decoding)
	plainText, err = encryptor.PKCS7Unpad(plainText, 32)
	if err != nil {
		return nil, err
	}
	g.Dump(plainText)
	return gjson.New(plainText), nil
}
