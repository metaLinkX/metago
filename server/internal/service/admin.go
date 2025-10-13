// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/adminRole"
	"server/internal/model/entity"
	"server/internal/model/input/adminin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IAdminMember interface {
		// GetOneById 换绑邮箱
		GetOneById(ctx context.Context, id int64) (user *entity.AdminMember, err error)
		// LoginMemberInfo 获取登录用户信息
		LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error)
	}
	IAdminMenu interface {
		// List 获取菜单列表
		List(ctx context.Context, in *adminin.MenuListInp) (res *adminin.MenuListModel, err error)
		// LoginPermissions 获取登录成功后的细粒度权限
		LoginPermissions(ctx context.Context, memberId int64) (lists adminin.MemberLoginPermissions, err error)
		// GetMenuList 获取菜单列表
		GetMenuList(ctx context.Context, memberId int64) (res *adminRole.DynamicRes, err error)
		// GetFastList 获取菜单列表
		GetFastList(ctx context.Context) (res map[int64]*entity.AdminMenu, err error)
		HasMenus(ctx context.Context, menuNames []string) ([]gdb.Value, error)
	}
	IAdminRole interface {
		// List 获取列表
		List(ctx context.Context, in *adminin.RoleListInp) (res *adminin.RoleListModel, totalCount int, err error)
	}
)

var (
	localAdminMember IAdminMember
	localAdminMenu   IAdminMenu
	localAdminRole   IAdminRole
)

func AdminMember() IAdminMember {
	if localAdminMember == nil {
		panic("implement not found for interface IAdminMember, forgot register?")
	}
	return localAdminMember
}

func RegisterAdminMember(i IAdminMember) {
	localAdminMember = i
}

func AdminMenu() IAdminMenu {
	if localAdminMenu == nil {
		panic("implement not found for interface IAdminMenu, forgot register?")
	}
	return localAdminMenu
}

func RegisterAdminMenu(i IAdminMenu) {
	localAdminMenu = i
}

func AdminRole() IAdminRole {
	if localAdminRole == nil {
		panic("implement not found for interface IAdminRole, forgot register?")
	}
	return localAdminRole
}

func RegisterAdminRole(i IAdminRole) {
	localAdminRole = i
}
