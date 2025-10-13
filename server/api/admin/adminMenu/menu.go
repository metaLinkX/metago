package adminMenu

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/input/adminin"
	"server/internal/model/input/form"
)

// ListReq 获取菜单列表
type ListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"菜单" summary:"获取菜单列表"`
	adminin.MenuListInp
}

type ListRes struct {
	*adminin.MenuListModel
	form.PageRes
}
