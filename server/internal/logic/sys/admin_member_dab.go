// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 1.0.0
package sys

import (
	"context"
	"fmt"
	"server/internal/dao"
	"server/internal/model/input/form"
	"server/internal/model/input/sysin"
	"server/internal/service"
	"server/library/convert"
	"server/library/excel"
	"server/library/hgorm"
	"server/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysAdminMemberDab struct{}

func NewSysAdminMemberDab() *sSysAdminMemberDab {
	return &sSysAdminMemberDab{}
}

func init() {
	service.RegisterSysAdminMemberDab(NewSysAdminMemberDab())
}

// Model 管理员_用户表ORM模型
func (s *sSysAdminMemberDab) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.AdminModel(dao.AdminMember.Ctx(ctx), option...)
}

// List 获取管理员_用户表列表
func (s *sSysAdminMemberDab) List(ctx context.Context, in *sysin.AdminMemberDabListInp) (list []*sysin.AdminMemberDabListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.AdminMemberDabListModel{})

	// 查询管理员ID
	if in.Id != 0 {
		mod = mod.Where(dao.AdminMember.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status != 0 {
		mod = mod.Where(dao.AdminMember.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AdminMember.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AdminMember.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取管理员_用户表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出管理员_用户表
func (s *sSysAdminMemberDab) Export(ctx context.Context, in *sysin.AdminMemberDabListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.AdminMemberDabExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出管理员_用户表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.AdminMemberDabExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增管理员_用户表
func (s *sSysAdminMemberDab) Edit(ctx context.Context, in *sysin.AdminMemberDabEditInp) (err error) {
	// 验证'InviteCode'唯一
	if err = hgorm.IsUnique(ctx, &dao.AdminMember, g.Map{dao.AdminMember.Columns().InviteCode: in.InviteCode}, "邀请码已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.AdminMemberDabUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改管理员_用户表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.AdminMemberDabInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增管理员_用户表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除管理员_用户表
func (s *sSysAdminMemberDab) Delete(ctx context.Context, in *sysin.AdminMemberDabDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除管理员_用户表失败，请稍后重试！")
		return
	}
	return
}

// View 获取管理员_用户表指定信息
func (s *sSysAdminMemberDab) View(ctx context.Context, in *sysin.AdminMemberDabViewInp) (res *sysin.AdminMemberDabViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取管理员_用户表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新管理员_用户表状态
func (s *sSysAdminMemberDab) Status(ctx context.Context, in *sysin.AdminMemberDabStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.AdminMember.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新管理员_用户表状态失败，请稍后重试！")
		return
	}
	return
}