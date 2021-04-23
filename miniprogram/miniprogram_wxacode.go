package miniprogram

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/guid"
)

//小程序码
type AppCode struct {
	mp  *MiniProgram
	raw *gjson.Json
}

//错误
type AppCodeError struct {
	ErrCode int    //错误代码
	ErrMsg  string //错误信息
}

//AppCode
func (mp *MiniProgram) AppCode() *AppCode {
	return &AppCode{
		mp: mp,
	}
}

//ToJson gjson>Json @see https://pkg.go.dev/github.com/gogf/gf/encoding/gjson
func (c *AppCode) ToJson() *gjson.Json {
	return c.raw
}

//Save 保存小程序码到文件
func (c *AppCode) Save(path string, name ...string) (string, *AppCodeError) {
	var codeName string
	if len(name) > 0 {
		codeName = name[0]
	} else {
		codeName = guid.S() + ".png"
	}
	if c.raw.GetInt("errcode") != 0 {
		return "", &AppCodeError{
			ErrCode: c.raw.GetInt("errcode"),
			ErrMsg:  c.raw.GetString("errmsg"),
		}
	}
	err := gfile.PutBytes(fmt.Sprintf("%s/%s", path, codeName), c.raw.GetBytes("buffer"))
	if err != nil {
		return "", &AppCodeError{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
	}
	return codeName, nil
}

//获取小程序码
func (c *AppCode) CreateQrCode(path string, width ...int) *AppCode {
	var param = g.Map{
		"path": path,
	}
	if len(width) > 0 {
		param["width"] = width[0]
	}
	client := c.mp.getClientWithToken()
	c.raw = client.RequestJson("POST", "cgi-bin/wxaapp/createwxaqrcode", param)
	return c
}

//获取小程序码 有数量限制，详细查看@see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (c *AppCode) Get(path string, config ...g.Map) *AppCode {
	client := c.mp.getClientWithToken()
	param := make(g.Map)
	if len(config) > 0 {
		param = config[0]
	}
	param["path"] = path
	c.raw = client.RequestJson("POST", "wxa/getwxacode", param)
	return c
}

//获取小程序码 不限制数量
func (c *AppCode) GetUnlimit(scene string, config ...g.Map) *AppCode {
	client := c.mp.getClientWithToken()
	param := make(g.Map)
	if len(config) > 0 {
		param = config[0]
	}
	param["scene"] = scene
	c.raw = client.RequestJson("POST", "wxa/getwxacodeunlimit", param)
	return c
}
