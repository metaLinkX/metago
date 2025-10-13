package admin

import (
	"context"
	"server/api/admin/genCodes"
	"server/internal/service"
)

// Delete 删除
func (c *ControllerGenCodes) Delete(ctx context.Context, req *genCodes.DeleteReq) (res *genCodes.DeleteRes, err error) {
	err = service.SysGenCodes().Delete(ctx, &req.GenCodesDeleteInp)
	return
}

// Edit 更新
func (c *ControllerGenCodes) Edit(ctx context.Context, req *genCodes.EditReq) (res *genCodes.EditRes, err error) {
	data, err := service.SysGenCodes().Edit(ctx, &req.GenCodesEditInp)
	if err != nil {
		return
	}

	res = new(genCodes.EditRes)
	res.GenCodesEditModel = data
	return
}

// MaxSort 最大排序
func (c *ControllerGenCodes) MaxSort(ctx context.Context, req *genCodes.MaxSortReq) (res *genCodes.MaxSortRes, err error) {
	res = new(genCodes.MaxSortRes)
	res.GenCodesMaxSortModel, err = service.SysGenCodes().MaxSort(ctx, &req.GenCodesMaxSortInp)
	return
}

// View 获取指定信息
func (c *ControllerGenCodes) View(ctx context.Context, req *genCodes.ViewReq) (res *genCodes.ViewRes, err error) {
	data, err := service.SysGenCodes().View(ctx, &req.GenCodesViewInp)
	if err != nil {
		return
	}

	res = new(genCodes.ViewRes)
	res.GenCodesViewModel = data
	return
}

// List 查看列表
func (c *ControllerGenCodes) List(ctx context.Context, req *genCodes.ListReq) (res *genCodes.ListRes, err error) {
	list, totalCount, err := service.SysGenCodes().List(ctx, &req.GenCodesListInp)
	if err != nil {
		return
	}

	res = new(genCodes.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Status 更新状态
func (c *ControllerGenCodes) Status(ctx context.Context, req *genCodes.StatusReq) (res *genCodes.StatusRes, err error) {
	err = service.SysGenCodes().Status(ctx, &req.GenCodesStatusInp)
	return
}

// Selects 获取指定信息
func (c *ControllerGenCodes) Selects(ctx context.Context, req *genCodes.SelectsReq) (res *genCodes.SelectsRes, err error) {
	data, err := service.SysGenCodes().Selects(ctx, &req.GenCodesSelectsInp)
	if err != nil {
		return
	}

	res = new(genCodes.SelectsRes)
	res.GenCodesSelectsModel = data
	return
}

// TableSelect 数据库表选项
func (c *ControllerGenCodes) TableSelect(ctx context.Context, req *genCodes.TableSelectReq) (res *genCodes.TableSelectRes, err error) {
	data, err := service.SysGenCodes().TableSelect(ctx, &req.GenCodesTableSelectInp)
	if err != nil {
		return
	}

	res = (*genCodes.TableSelectRes)(&data)
	return
}

// ColumnSelect 表字段选项
func (c *ControllerGenCodes) ColumnSelect(ctx context.Context, req *genCodes.ColumnSelectReq) (res *genCodes.ColumnSelectRes, err error) {
	data, err := service.SysGenCodes().ColumnSelect(ctx, &req.GenCodesColumnSelectInp)
	if err != nil {
		return
	}

	res = (*genCodes.ColumnSelectRes)(&data)
	return
}

// ColumnList 表字段列表
func (c *ControllerGenCodes) ColumnList(ctx context.Context, req *genCodes.ColumnListReq) (res *genCodes.ColumnListRes, err error) {
	data, err := service.SysGenCodes().ColumnList(ctx, &req.GenCodesColumnListInp)
	if err != nil {
		return
	}

	res = (*genCodes.ColumnListRes)(&data)
	return
}

// Preview 生成预览
func (c *ControllerGenCodes) Preview(ctx context.Context, req *genCodes.PreviewReq) (res *genCodes.PreviewRes, err error) {
	data, err := service.SysGenCodes().Preview(ctx, &req.GenCodesPreviewInp)
	if err != nil {
		return
	}

	res = new(genCodes.PreviewRes)
	res.GenCodesPreviewModel = data
	return
}

// Build 生成预览
func (c *ControllerGenCodes) Build(ctx context.Context, req *genCodes.BuildReq) (res *genCodes.BuildRes, err error) {
	err = service.SysGenCodes().Build(ctx, &req.GenCodesBuildInp)
	return
}
