// Package admin

package admin

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/model/input/adminin"
	"server/internal/service"
)

type sAdminRole struct{}

func NewAdminRole() *sAdminRole {
	return &sAdminRole{}
}

func init() {
	service.RegisterAdminRole(NewAdminRole())
}

// List 获取列表
func (s *sAdminRole) List(ctx context.Context, in *adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error) {
	var (
		mod     = dao.AdminRole.Ctx(ctx)
		columns = dao.AdminRole.Columns()
		models  []*entity.AdminRole
		pid     int64 = 0
	)

	totalCount, err = mod.Count()
	if err != nil {
		return
	}
	if err = mod.Page(in.Page, in.PerPage).OrderDesc(columns.Sort).OrderDesc(columns.Id).Scan(&models); err != nil {
		return
	}

	res = new(adminin.RoleListModel)
	res.List = s.treeList(pid, models)
	return
}

// treeList 角色树列表
func (s *sAdminRole) treeList(pid int64, nodes []*entity.AdminRole) (list []*adminin.RoleTree) {
	list = make([]*adminin.RoleTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.RoleTree)
			item.AdminRole = *v
			item.Label = v.Name
			item.Value = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}
