package miniapp

import (
	"fmt"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/guid"
)

type AppCode struct {
	App *MiniApp
	Raw []byte
}

//错误
type AppCodeError struct {
	ErrCode int    //错误代码
	ErrMsg  string //错误信息
}

func (app *MiniApp) AppCode() *AppCode {
	return &AppCode{
		App: app,
	}
}

//Save 保存小程序码到文件
func (c *AppCode) Save(path string, name ...string) (string, *AppCodeError) {
	var codeName string
	if len(name) > 0 {
		codeName = name[0]
	} else {
		codeName = guid.S() + ".png"
	}
	if gjson.Valid(c.Raw) {
		err := gjson.New(c.Raw)
		if err.GetInt("errcode") != 0 {
			return "", &AppCodeError{
				ErrCode: err.GetInt("errcode"),
				ErrMsg:  err.GetString("errmsg"),
			}
		}
	}
	err := gfile.PutBytes(fmt.Sprintf("%s/%s", path, codeName), c.Raw)
	if err != nil {
		return "", &AppCodeError{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		}
	}
	return codeName, nil
}

//@see https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code#%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
func (c *AppCode) CreateQRCode(config ...g.Map) *AppCode {
	client := c.App.GetClientWithToken()
	param := make(g.Map)

	if len(config) > 0 {
		param = config[0]
	}

	c.Raw = client.RequestRaw("POST", "apps/qrcode", param)
	return c
}
