// Package admin

package admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"server/api/admin/adminRole"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/model/input/adminin"
	"server/internal/service"
	"server/library/convert"
	"server/library/validate"
)

type sAdminMenu struct{}

func NewAdminMenu() *sAdminMenu {
	return &sAdminMenu{}
}

func init() {
	service.RegisterAdminMenu(NewAdminMenu())
}

// List 获取菜单列表
func (s *sAdminMenu) List(ctx context.Context, in *adminin.MenuListInp) (res *adminin.MenuListModel, err error) {
	var models []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Order("sort asc,id desc").Scan(&models); err != nil {
		return
	}

	res = new(adminin.MenuListModel)
	res.List = s.treeList(0, models)
	return
}

// treeList 树状列表
func (s *sAdminMenu) treeList(pid int64, nodes []*entity.AdminMenu) (list []*adminin.MenuTree) {
	list = make([]*adminin.MenuTree, 0)
	for _, v := range nodes {
		if v.Pid == pid {
			item := new(adminin.MenuTree)
			item.AdminMenu = *v
			item.Label = v.Title
			item.Key = v.Id

			child := s.treeList(v.Id, nodes)
			if len(child) > 0 {
				item.Children = child
			}
			list = append(list, item)
		}
	}
	return
}

// LoginPermissions 获取登录成功后的细粒度权限
func (s *sAdminMenu) LoginPermissions(ctx context.Context, memberId int64) (lists adminin.MemberLoginPermissions, err error) {
	type Permissions struct {
		Permissions string `json:"permissions"`
	}

	var (
		allPermissions []*Permissions
		mod            = dao.AdminMenu.Ctx(ctx).Fields(dao.AdminMenu.Columns().Permissions).
				Where(dao.AdminMenu.Columns().Status, consts.DictStatusEnabled).
				WhereNot(dao.AdminMenu.Columns().Permissions, "")
	)

	// 非超管验证允许的菜单列表 todo

	if err = mod.Scan(&allPermissions); err != nil {
		return
	}

	// 无权限
	if len(allPermissions) == 0 {
		lists = append(lists, "value")
		return
	}

	for _, v := range allPermissions {
		for _, p := range gstr.Explode(`,`, v.Permissions) {
			lists = append(lists, p)
		}
	}

	lists = convert.UniqueSlice(lists)
	return
}

// GetMenuList 获取菜单列表
func (s *sAdminMenu) GetMenuList(ctx context.Context, memberId int64) (res *adminRole.DynamicRes, err error) {
	var (
		allMenus []*adminin.MenuRouteSummary
		menus    []*adminin.MenuRouteSummary
		treeMap  = make(map[string][]*adminin.MenuRouteSummary)
		mod      = dao.AdminMenu.Ctx(ctx).Where(dao.AdminMenu.Columns().Status, consts.DictStatusEnabled).WhereIn(dao.AdminMenu.Columns().Type, []int{consts.DictMenuTypeMenu, consts.DictMenuTypeDir})
	)

	if err = mod.Order(dao.AdminMenu.Columns().Sort, dao.AdminMenu.Columns().Id, "desc").Scan(&allMenus); err != nil || len(allMenus) == 0 {
		return
	}

	// 生产环境下隐藏一些菜单 todo 应该设置那个模式 有 这些菜单 而不是
	if gmode.IsProduct() || gmode.IsTesting() {
		newMenus := make([]*adminin.MenuRouteSummary, 0)
		devMenus := []string{"Develops", "doc"} // 如果你还有其他需要在生产环境隐藏的菜单，将菜单别名加入即可
		for _, menu := range allMenus {
			if !validate.InSlice(devMenus, menu.Name) {
				newMenus = append(newMenus, menu)
			}
		}
		allMenus = newMenus
	}

	for _, v := range allMenus {
		treeMap[gconv.String(v.Pid)] = append(treeMap[gconv.String(v.Pid)], v)
	}

	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getChildrenList(menus[i], treeMap)
	}

	res = new(adminRole.DynamicRes)
	res.List = append(res.List, s.genNaiveMenus(menus)...)
	return
}
func (s *sAdminMenu) genNaiveMenus(menus []*adminin.MenuRouteSummary) (sources []*adminin.MenuRoute) {
	for _, men := range menus {
		var source = new(adminin.MenuRoute)
		source.Name = men.Name
		source.Path = men.Path
		source.Redirect = men.Redirect
		source.Component = men.Component
		source.Meta = &adminin.MenuRouteMeta{
			Title:      men.Title,
			Icon:       men.Icon,
			KeepAlive:  men.KeepAlive == 1,
			Hidden:     men.Hidden == 1,
			Sort:       men.Sort,
			AlwaysShow: men.AlwaysShow == 1,
			ActiveMenu: men.ActiveMenu,
			IsRoot:     men.IsRoot == 1,
			FrameSrc:   men.FrameSrc,
			//Permissions: men.Permissions,
			Affix: men.Affix == 1,
			Type:  men.Type,
		}
		if len(men.Children) > 0 {
			source.Children = append(source.Children, s.genNaiveMenus(men.Children)...)
		}
		sources = append(sources, source)
	}
	return
}

// getChildrenList 生成菜单树
func (s *sAdminMenu) getChildrenList(menu *adminin.MenuRouteSummary, treeMap map[string][]*adminin.MenuRouteSummary) (err error) {
	menu.Children = treeMap[gconv.String(menu.Id)]
	for i := 0; i < len(menu.Children); i++ {
		if err = s.getChildrenList(menu.Children[i], treeMap); err != nil {
			return
		}
	}
	return
}

// GetFastList 获取菜单列表
func (s *sAdminMenu) GetFastList(ctx context.Context) (res map[int64]*entity.AdminMenu, err error) {
	var models []*entity.AdminMenu
	if err = dao.AdminMenu.Ctx(ctx).Scan(&models); err != nil {
		return
	}

	res = make(map[int64]*entity.AdminMenu, len(models))
	for _, v := range models {
		res[v.Id] = v
	}
	return
}

func (s *sAdminMenu) HasMenus(ctx context.Context, menuNames []string) ([]gdb.Value, error) {
	return dao.AdminMenu.Ctx(ctx).Fields(dao.AdminMenu.Columns().Name).WhereIn(dao.AdminMenu.Columns().Name, menuNames).Array()
}
