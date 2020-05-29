package main

import (
	"fmt"

	"gitee.com/wallesoft/ewa/miniprogram"
	"github.com/gogf/gf/frame/g"
)

//"../../miniprogram"

func main() {
	c := g.Map{
		"Appid":  "appid",
		"Secret": "secret",
	}
	//t := new(Test)
	//err := gconv.Struct(c, t)
	//g.Dump(t)
	err := miniprogram.Config(c)
	g.Dump(miniprogram.MiniProgram{})
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(miniprogram.Miniprogram)
}
