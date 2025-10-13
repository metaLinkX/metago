package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gmode"

	"server/api/admin/adminSite"
)

func (c *ControllerAdminSite) SiteConfig(ctx context.Context, req *adminSite.SiteConfigReq) (res *adminSite.SiteConfigRes, err error) {
	// todo
	res = &adminSite.SiteConfigRes{
		Version: "",
		WsAddr:  "",
		Domain:  "",
		Mode:    gmode.Mode(),
	}
	return
}
