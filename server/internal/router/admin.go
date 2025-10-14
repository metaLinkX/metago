// Package router

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/consts"
	"server/internal/controller/admin"
	"server/internal/controller/admin/sys"
	"server/internal/service"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(consts.AdminPrefix, func(group *ghttp.RouterGroup) {
		group.Bind(
			admin.NewAdminLogin(),
			admin.NewAdminSite(),
		)
		group.Middleware(service.Middleware().AdminAuth)
		group.Bind(admin.NewAdminMember(), admin.NewAdminRole(), admin.NewGenCodes(), admin.NewAdminMenu(), sys.HgUser)
	})
}
