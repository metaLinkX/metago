package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/adminLogin"
)

func (c *ControllerAdminLogin) LoginConfig(ctx context.Context, req *adminLogin.LoginConfigReq) (res *adminLogin.LoginConfigRes, err error) {
	res = new(adminLogin.LoginConfigRes)
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}
	res.LoginConfig = login
	return res, err
}
