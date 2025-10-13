package admin

import (
	"context"
	"server/api/admin/adminMember"
	"server/internal/service"
)

func (c *ControllerAdminMember) Info(ctx context.Context, req *adminMember.InfoReq) (res *adminMember.InfoRes, err error) {
	data, err := service.AdminMember().LoginMemberInfo(ctx)
	if err != nil {
		return
	}

	res = new(adminMember.InfoRes)
	res.LoginMemberInfoModel = data
	return res, err
}
