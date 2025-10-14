// Package hguser
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 1.0.0
package hguser

import (
	"server/internal/model/input/sysin"

	"github.com/gogf/gf/v2/frame/g"

	"server/internal/model"
	"server/internal/model/input/form"
)

// ListReq 查询租户列表
type ListReq struct {
	g.Meta `path:"/hgUser/list" method:"get" tags:"租户" summary:"获取租户列表"`
	sysin.HgUserListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.HgUserListModel `json:"list"   dc:"数据列表"`
	// 添加字典依赖

	SysNormalDisableOption []model.Option `json:"sysNormalDisableOption" dc:"字典选项"`

	SysLoginStatusOption []model.Option `json:"sysLoginStatusOption" dc:"字典选项"`
}

// ExportReq 导出租户列表
type ExportReq struct {
	g.Meta `path:"/hgUser/export" method:"get" tags:"租户" summary:"导出租户列表"`
	sysin.HgUserListInp
}

type ExportRes struct{}

// ViewReq 获取租户指定信息
type ViewReq struct {
	g.Meta `path:"/hgUser/view" method:"get" tags:"租户" summary:"获取租户指定信息"`
	sysin.HgUserViewInp
}

type ViewRes struct {
	*sysin.HgUserViewModel
}

// EditReq 修改/新增租户
type EditReq struct {
	g.Meta `path:"/hgUser/edit" method:"post" tags:"租户" summary:"修改/新增租户"`
	sysin.HgUserEditInp
}

type EditRes struct{}

// DeleteReq 删除租户
type DeleteReq struct {
	g.Meta `path:"/hgUser/delete" method:"post" tags:"租户" summary:"删除租户"`
	sysin.HgUserDeleteInp
}

type DeleteRes struct{}