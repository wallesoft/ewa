// Copyright 2020 ewa Author(https://gitee.com/wallesoft/ewa)
// ewa (https://wallesoft.gitee.io/ewa)
//
// 小程序模块

package miniprogram

import (
	"errors"

	"github.com/gogf/gf/util/gconv"
)

type MiniProgram struct {
	Appid        string // appid
	Secret       string // secret
	ResponseType string // 返回类型
	//@todo  log
	*Auth
}

func Config(config map[string]interface{}) error {
	if config == nil || len(config) == 0 {
		return errors.New("Miniprogram configuration cannot be empty")
	}

	mp := new(MiniProgram)
	err := gconv.Struct(config, mp)
	if err != nil {
		return err
	}

	return nil
}
