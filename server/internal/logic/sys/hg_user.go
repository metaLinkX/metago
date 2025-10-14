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

type sSysHgUser struct{}

func NewSysHgUser() *sSysHgUser {
	return &sSysHgUser{}
}

func init() {
	service.RegisterSysHgUser(NewSysHgUser())
}

// Model 租户ORM模型
func (s *sSysHgUser) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.AdminModel(dao.HgUser.Ctx(ctx), option...)
}

// List 获取租户列表
func (s *sSysHgUser) List(ctx context.Context, in *sysin.HgUserListInp) (list []*sysin.HgUserListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.HgUserListModel{})

	// 查询用户ID
	if in.Id != 0 {
		mod = mod.Where(dao.HgUser.Columns().Id, in.Id)
	}

	// 查询用户状态
	if in.UserStatus != 0 {
		mod = mod.Where(dao.HgUser.Columns().UserStatus, in.UserStatus)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.HgUser.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.HgUser.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取租户列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出租户
func (s *sSysHgUser) Export(ctx context.Context, in *sysin.HgUserListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.HgUserExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出租户-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.HgUserExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增租户
func (s *sSysHgUser) Edit(ctx context.Context, in *sysin.HgUserEditInp) (err error) {
	// 验证'IdCard'唯一
	if err = hgorm.IsUnique(ctx, &dao.HgUser, g.Map{dao.HgUser.Columns().IdCard: in.IdCard}, "身份证已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.HgUserUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改租户失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.HgUserInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增租户失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除租户
func (s *sSysHgUser) Delete(ctx context.Context, in *sysin.HgUserDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除租户失败，请稍后重试！")
		return
	}
	return
}

// View 获取租户指定信息
func (s *sSysHgUser) View(ctx context.Context, in *sysin.HgUserViewInp) (res *sysin.HgUserViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取租户信息，请稍后重试！")
		return
	}
	return
}