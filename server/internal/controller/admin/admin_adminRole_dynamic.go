package admin

import (
	"context"
	"server/api/admin/adminRole"
	"server/internal/service"
	"server/library/contexts"
)

func (c *ControllerAdminRole) Dynamic(ctx context.Context, req *adminRole.DynamicReq) (res *adminRole.DynamicRes, err error) {
	return service.AdminMenu().GetMenuList(ctx, contexts.GetUserId(ctx))
}
