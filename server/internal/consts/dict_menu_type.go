package consts

import "server/internal/model"

const (
	DictMenuTypeDir  int = 1 // 目录
	DictMenuTypeMenu int = 2 // 菜单
	DictMenuTypeBtn  int = 3 // 按钮
)

var (
	DictMenuTypeMap = map[int]string{
		DictMenuTypeDir:  "目录",
		DictMenuTypeMenu: "菜单",
		DictMenuTypeBtn:  "按钮",
	}
	DictMenuTypeOption = []model.MdOption{
		{Label: "目录", Value: DictMenuTypeDir},
		{Label: "菜单", Value: DictMenuTypeMenu},
		{Label: "按钮", Value: DictMenuTypeBtn},
	}
	DictMenuTypeSlice = []int{DictMenuTypeDir, DictMenuTypeMenu, DictMenuTypeBtn}
)
