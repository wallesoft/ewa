package payment

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

type Client struct {
	payment *Payment
}

//Request
func (c *Client) Request() {

}

func (this *Client) rsaEncrypt(origData []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	h.Write(origData)
	hashed := h.Sum(nil)
	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, this.Priv.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

func (this *Client) RsaDecrypt(ciphertext string) (string, error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, this.Priv.(*rsa.PrivateKey), cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return "", err
	}
	return string(plaintext), nil
}
