package main

import {
	"gitee.com/wallesoft/ewa/miniprogram"
}

func main() {
	
	app := miniprogram.GetApp()
	// code 小程序获取用户登录信息
	res := app.Auth().Session(code)

}