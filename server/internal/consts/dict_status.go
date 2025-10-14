package consts

import "server/internal/model"

const (
	DictStatusEnabled int = 1  // 启用
	DictStatusDisable int = -1 // 禁用
)

var (
	DictStatusMap = map[int]string{
		DictStatusDisable: "禁用",
		DictStatusEnabled: "启用",
	}
	DictStatusOption = []model.Option{
		{Label: "禁用", Value: DictStatusDisable},
		{Label: "启用", Value: DictStatusEnabled},
	}
	DictStatusSlice = []int{DictStatusDisable, DictStatusEnabled}
)
