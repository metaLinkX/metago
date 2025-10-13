package consts

import "server/internal/model"

const (
	DictYes = 1
	DictNot = -1
)

var (
	DictYesNotMap = map[int]string{
		DictYes: "是",
		DictNot: "否",
	}
	DictOption = []model.MdOption{
		{Label: "是", Value: DictYes},
		{Label: "否", Value: DictNot},
	}
)
