// Package admin

package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/model/input/adminin"
	"server/internal/service"
	"server/library/contexts"
)

type sAdminMember struct {
}

func NewAdminMember() *sAdminMember {
	return &sAdminMember{}
}

func init() {
	service.RegisterAdminMember(NewAdminMember())
}

// GetOneById 换绑邮箱
func (s *sAdminMember) GetOneById(ctx context.Context, id int64) (user *entity.AdminMember, err error) {
	err = dao.AdminMember.Ctx(ctx).WherePri(id).Scan(&user)
	return
}

// LoginMemberInfo 获取登录用户信息
func (s *sAdminMember) LoginMemberInfo(ctx context.Context) (res *adminin.LoginMemberInfoModel, err error) {
	var memberId = contexts.GetUserId(ctx)
	if memberId <= 0 {
		err = gerror.New("用户身份异常，请重新登录！")
		return
	}

	if err = dao.AdminMember.Ctx(ctx).WherePri(memberId).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取用户信息失败，请稍后重试！")
		return
	}

	if res == nil {
		err = gerror.New("用户不存在！")
		return
	}

	// 细粒度权限
	permissions, err := service.AdminMenu().LoginPermissions(ctx, memberId)
	if err != nil {
		return
	}
	res.Permissions = permissions
	res.Mobile = gstr.HideStr(res.Mobile, 40, `*`)
	res.Email = gstr.HideStr(res.Email, 40, `*`)
	return
}
