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
	config   	 MiniProgramConfig
	ResponseType string // 返回类型
	// Logger specifies the logger for miniprogram
	Logger *glog.Logger
}
type MiniProgramConfig struct {
	Appid  		string  // appid
	Secret 		string  // scret
}

func ConfigFromMap(config map[string]interface{}) (*MiniProgram, error) {
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
