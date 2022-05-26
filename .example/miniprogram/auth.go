package main

import (
	"gitee.com/wallesoft/ewa/miniprogram"
	"github.com/gogf/gf/v2/frame/g"
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
	//mp, err := miniprogram.Config(c)
	m := miniprogram.Instance(c)
	g.Dump(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//fmt.Println(miniprogram.Miniprogram)
}
