// Copyright 2020 ewa Author(https://gitee.com/wallesoft/ewa)
// ewa (https://wallesoft.gitee.io/ewa)
//
// 小程序模块

package miniprogram

import (
	"errors"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

type MiniProgram struct {
	Appid        string // appid
	Secret       string // secret
	ResponseType string // 返回类型
	// Logger specifies the logger for miniprogram
	Logger *glog.Logger
	// auth
	Auth *Auth
}

func Config(config map[string]interface{}) (*MiniProgram, error) {
	mp := new(MiniProgram)
	if config == nil || len(config) == 0 {
		return mp, errors.New("Miniprogram configuration cannot be empty")
	}
	// ResponseType

	err := gconv.Struct(config, mp)
	if err != nil {
		return mp, err
	}

	return mp, nil
}
func Instance(config map[string]interface{}) *MiniProgram {
	mp, err := Config(config)
	if err != nil {
		// log
	}
}
