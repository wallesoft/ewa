package encryptor

import (
	"bytes"
	"errors"
	"sort"
	"strings"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gbinary"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
)

type Encryptor struct {
	AppID     string
	Token     string
	AesKey    string
	BlockSize int
}

const (
	ERROR_INVALID_SIGNATURE = -40001 // Signature verification failed
	ERROR_PARSE_XML         = -40002 // Parse XML failed
	ERROR_CALC_SIGNATURE    = -40003 // Calculating the signature failed
	ERROR_INVALID_AES_KEY   = -40004 // Invalid AESKey
	ERROR_INVALID_APP_ID    = -40005 // Check AppID failed
	ERROR_ENCRYPT_AES       = -40006 // AES EncryptionInterface failed
	ERROR_DECRYPT_AES       = -40007 // AES Decryption failed
	ERROR_INVALID_XML       = -40008 // Invaild XML
	EROOR_BASE64_ENCODE     = -40009 // Base64 encoding failed
	ERROR_BASE64_DECODE     = -40010 // Base64 decoding failed
	ERROR_XML_BULID         = -40011 // XML build failed
	ILLEGAL_BUFFER          = -41003 // Illegal buffer

)

// func New(config map[string]interface{}) (*Encryptor, error) {
// 	if config == nil || len(config) == 0 {
// 		return nil, errors.New("Encryptor configuration cannot be empty")
// 	}
// 	config = gutil.MapCopy(config)
// 	var c *Encryptor
// 	if err := gconv.Struct(config, &c); err != nil {
// 		return nil, err
// 	}
// 	if c.AesKey != "" {
// 		aesKey, err := gbase64.DecodeToString(c.AesKey)
// 		if err != nil {
// 			return nil, err
// 		}
// 		c.AesKey = aesKey
// 	} else {
// 		return nil, errors.New("Encryptor configuration aes_key cannot be empty")
// 	}
// 	return c, nil
// }
type Config struct {
	AppID string `json:"app_id"` //appid
	//AppSecret      string `json:"app_secret"`       //appsecret
	Token          string `json:"token"` //token
	EncodingAESKey string `json:"encoding_aes_key"`
	BlockSize      int
}

func New(config Config) *Encryptor {
	aesKey := ""
	if config.EncodingAESKey != "" {
		key, err := gbase64.DecodeToString(config.EncodingAESKey + "=")
		if err != nil {
			panic(err)
		}
		aesKey = key
	}
	if config.BlockSize == 0 {
		config.BlockSize = 32
	}
	return &Encryptor{
		AppID:     config.AppID,
		Token:     config.Token,
		AesKey:    aesKey,
		BlockSize: config.BlockSize,
	}

}

//Encrypt encrypt message.
func (e *Encryptor) Encrypt(rawXML []byte, nonce string, timestamp int) ([]byte, error) {
	text := bytes.Join([][]byte{grand.B(16), gbinary.BeEncodeUint32(gconv.Uint32(len(rawXML))), rawXML, gconv.Bytes(e.AppID)}, []byte(""))
	xml := PKCS7Pad(text, e.BlockSize)
	encrypted, err := gaes.Encrypt(xml, gconv.Bytes(e.AesKey), gconv.Bytes(gstr.SubStr(e.AesKey, 0, 16)))
	if err != nil {
		return nil, err
	}
	// gbase64.Encode()
	return gbase64.Encode(encrypted), nil
}

//Decrypt decrypt message
func (e *Encryptor) Decrypt(content []byte) ([]byte, error) {
	decoding, err := gbase64.Decode(content)
	if err != nil {
		return nil, NewError(ERROR_BASE64_DECODE, err.Error())
	}
	decrypted, err := gaes.Decrypt(decoding, gconv.Bytes(e.AesKey), gconv.Bytes(gstr.SubStr(e.AesKey, 0, 16)))
	if err != nil {
		return nil, NewError(ERROR_DECRYPT_AES, err.Error())
	}
	//unpad
	result, err := PKCS7Unpad(decrypted, e.BlockSize)
	if err != nil {
		return nil, err
	}
	// len := len(result)
	contents := result[16:]
	//网络字节序为大端序
	msgLen := gbinary.BeDecodeToUint32(contents[:4])
	if gconv.String(contents[msgLen+4:]) != e.AppID {
		return nil, NewError(ERROR_INVALID_APP_ID, "Invalid appId.")
	}
	return contents[4 : msgLen+4], nil
}

//GetToken is this necessary?
func (e *Encryptor) GetToken() string {
	return e.Token
}

//Signature
func Signature(s []string) string {
	sort.Strings(s)
	return gsha1.Encrypt(strings.Join(s, ""))
}

//PKCS7Pad pad.
func PKCS7Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//PKCS7Unpad unpad.
func PKCS7Unpad(src []byte, blockSize int) ([]byte, error) {
	length := len(src)

	unpadding := int(src[length-1])
	if unpadding < 1 || unpadding > blockSize {
		unpadding = 0
	}
	padding := src[length-unpadding:]
	for i := 0; i < unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil, errors.New("invalid padding")
		}
	}
	return src[:(length - unpadding)], nil
}
