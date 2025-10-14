// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 1.0.0
package sys

import (
	"context"
	"server/api/admin/hguser"
	"server/internal/consts"
	"server/internal/model/input/sysin"
	"server/internal/service"
)

var (
	HgUser = cHgUser{}
)

type cHgUser struct{}

// List 查看租户列表
func (c *cHgUser) List(ctx context.Context, req *hguser.ListReq) (res *hguser.ListRes, err error) {
	list, totalCount, err := service.SysHgUser().List(ctx, &req.HgUserListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.HgUserListModel{}
	}

	res = new(hguser.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)

	// 添加字典依赖

	res.SysNormalDisableOption = consts.DictSysNormalDisableOption

	res.SysLoginStatusOption = consts.DictSysLoginStatusOption

	return
}

// Export 导出租户列表
func (c *cHgUser) Export(ctx context.Context, req *hguser.ExportReq) (res *hguser.ExportRes, err error) {
	err = service.SysHgUser().Export(ctx, &req.HgUserListInp)
	return
}

// Edit 更新租户
func (c *cHgUser) Edit(ctx context.Context, req *hguser.EditReq) (res *hguser.EditRes, err error) {
	err = service.SysHgUser().Edit(ctx, &req.HgUserEditInp)
	return
}

// View 获取指定租户信息
func (c *cHgUser) View(ctx context.Context, req *hguser.ViewReq) (res *hguser.ViewRes, err error) {
	data, err := service.SysHgUser().View(ctx, &req.HgUserViewInp)
	if err != nil {
		return
	}

	res = new(hguser.ViewRes)
	res.HgUserViewModel = data
	return
}

// Delete 删除租户
func (c *cHgUser) Delete(ctx context.Context, req *hguser.DeleteReq) (res *hguser.DeleteRes, err error) {
	err = service.SysHgUser().Delete(ctx, &req.HgUserDeleteInp)
	return
}