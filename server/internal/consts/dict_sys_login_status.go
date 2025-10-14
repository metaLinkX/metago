package consts

import (
	"server/internal/model"
)

const (
	DictSysLoginStatusDLCG int = 1 // 登录成功
	DictSysLoginStatusDLSB int = 2 // 登录失败
)

var (
	DictSysLoginStatusMap = map[int]string{
		DictSysLoginStatusDLCG: "登录成功",
		DictSysLoginStatusDLSB: "登录失败",
	}
	DictSysLoginStatusOption = []model.Option{
		{
			Key:       DictSysLoginStatusDLCG,
			Label:     "登录成功",
			Value:     DictSysLoginStatusDLCG,
			ValueType: "int",
			Type:      "sys_login_status",
			ListClass: "error",
			Extra:     nil,
		},
		{
			Key:       DictSysLoginStatusDLSB,
			Label:     "登录失败",
			Value:     DictSysLoginStatusDLSB,
			ValueType: "int",
			Type:      "sys_login_status",
			ListClass: "warning",
			Extra:     nil,
		},
	}
)