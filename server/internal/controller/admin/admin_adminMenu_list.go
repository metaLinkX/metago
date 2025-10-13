package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/adminMenu"
)

func (c *ControllerAdminMenu) List(ctx context.Context, req *adminMenu.ListReq) (res *adminMenu.ListRes, err error) {
	res = new(adminMenu.ListRes)
	res.MenuListModel, err = service.AdminMenu().List(ctx, &req.MenuListInp)
	return res, err
}
