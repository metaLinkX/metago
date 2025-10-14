package consts

import (
	"server/internal/model"
)

const (
	DictSysNormalDisableZC int = 1 // 正常
	DictSysNormalDisableTY int = 2 // 停用
)

var (
	DictSysNormalDisableMap = map[int]string{
		DictSysNormalDisableZC: "正常",
		DictSysNormalDisableTY: "停用",
	}
	DictSysNormalDisableOption = []model.Option{
		{
			Key:       DictSysNormalDisableZC,
			Label:     "正常",
			Value:     DictSysNormalDisableZC,
			ValueType: "int",
			Type:      "sys_normal_disable",
			ListClass: "default",
			Extra:     nil,
		},
		{
			Key:       DictSysNormalDisableTY,
			Label:     "停用",
			Value:     DictSysNormalDisableTY,
			ValueType: "int",
			Type:      "sys_normal_disable",
			ListClass: "info",
			Extra:     nil,
		},
	}
)
