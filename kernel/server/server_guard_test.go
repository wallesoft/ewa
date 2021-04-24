package server_test

import (
	"testing"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gconv"
)

var decrypted = []byte(``)
var bodyRaw = []byte(``)

func Test_parseXmlMessage(t *testing.T) {
	data := []byte(``)
	if m, err := gxml.DecodeWithoutRoot(data); err != nil {
		panic(err)
	} else {
		gtest.C(t, func(t *gtest.T) {
			if val, ok := m["AppId"]; ok {
				t.Assert(gconv.String(val), "")
			}
			if encrypted, ok := m["Encrypt"]; ok {
				g.Dump(encrypted)
			}
		})

	}

}

func Test_DecryptMessage(t *testing.T) {
	a := []string{}
	gtest.C(t, func(t *gtest.T) {
		t.Assert(encryptor.Signature(a), "")
	})
}
func Test_GetMessage(t *testing.T) {

}
