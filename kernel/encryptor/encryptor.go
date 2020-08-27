package encryptor

import (
	"bytes"
	"errors"

	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/gogf/gf/util/gutil"
)

type Encryptor struct {
	AppId     string
	Token     string
	AesKey    string
	blockSize int
}

const (
	ERROR_INVALID_SIGNATURE = -40001 // Signature verification failed
	ERROR_PARSE_XML         = -40002 // Parse XML failed
	ERROR_CALC_SIGNATURE    = -40003 // Calculating the signature failed
	ERROR_INVALID_AES_KEY   = -40004 // Invalid AESKey
	ERROR_INVALID_APP_ID    = -40005 // Check AppId failed
	ERROR_ENCRYPT_AES       = -40006 // AES EncryptionInterface failed
	ERROR_DECRYPT_AED       = -40007 // AES Decryption failed
	ERROR_INVALID_XML       = -40008 // Invaild XML
	EROOR_BASE64_ENCODE     = -40009 // Base64 encoding failed
	ERROR_BASE64_DECODE     = -40010 // Base64 decoding failed
	ERROR_XML_BULID         = -40011 // XML build failed
	ILLEGAL_BUFFER          = -41003 // Illegal buffer

)

func New(config map[string]interface{}) (*Encryptor, error) {
	if config == nil || len(config) == 0 {
		return nil, errors.New("Encryptor configuration cannot be empty")
	}
	config = gutil.MapCopy(config)
	var c *Encryptor
	if err := gconv.Struct(config, &c); err != nil {
		return nil, err
	}
	return c, nil
}

//Encrypt encrypt the message and return gxml
func (e *Encryptor) Encrypt(rawXML []byte) string {
	text := bytes.Join([][]byte{grand.B(16), gconv.Bytes(gconv.Uint32(len(rawXML))), rawXML, gconv.Bytes(e.AppId)}, []byte(""))

	xml, err := e.Pkcs7Pad(text)
}

// func (e *Encryptor) Decrypt()
// func (e *Encryptor) GetToken()
// func (e *Encryptor) signature()
//Pkcs7Pad pad.
func (e *Encryptor) Pkcs7Pad(text string) (string, error) {
	if e.blockSize > 256 {
		return "", errors.New("blockSize may not be more than 256")
	}
	padding := e.blockSize - (len(text) % e.blockSize)
	pattern := gstr.Chr(padding)
	return text + gstr.Repeat(pattern, padding), nil
}

//Pkcs7 unpad
func (e *Encryptor) Pkcs7Unpad(text string) string {
	pad := gstr.Ord(gstr.SubStr(text, -1))
	if pad < 1 || pad > e.blockSize {
		pad = 0
	}
	return gstr.SubStr(text, 0, (len(text) - pad))
}
