// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"server/api/admin/adminLogin"
	"server/api/admin/adminMember"
	"server/api/admin/adminMenu"
	"server/api/admin/adminRole"
	"server/api/admin/adminSite"
	"server/api/admin/genCodes"
)

type IAdminAdminLogin interface {
	LoginCaptcha(ctx context.Context, req *adminLogin.LoginCaptchaReq) (res *adminLogin.LoginCaptchaRes, err error)
	LoginConfig(ctx context.Context, req *adminLogin.LoginConfigReq) (res *adminLogin.LoginConfigRes, err error)
	AccountLogin(ctx context.Context, req *adminLogin.AccountLoginReq) (res *adminLogin.AccountLoginRes, err error)
}

type IAdminAdminMember interface {
	Info(ctx context.Context, req *adminMember.InfoReq) (res *adminMember.InfoRes, err error)
}

type IAdminAdminMenu interface {
	List(ctx context.Context, req *adminMenu.ListReq) (res *adminMenu.ListRes, err error)
}

type IAdminAdminRole interface {
	Dynamic(ctx context.Context, req *adminRole.DynamicReq) (res *adminRole.DynamicRes, err error)
}

type IAdminAdminSite interface {
	SiteConfig(ctx context.Context, req *adminSite.SiteConfigReq) (res *adminSite.SiteConfigRes, err error)
}

type IAdminGenCodes interface {
	List(ctx context.Context, req *genCodes.ListReq) (res *genCodes.ListRes, err error)
	View(ctx context.Context, req *genCodes.ViewReq) (res *genCodes.ViewRes, err error)
	Edit(ctx context.Context, req *genCodes.EditReq) (res *genCodes.EditRes, err error)
	Delete(ctx context.Context, req *genCodes.DeleteReq) (res *genCodes.DeleteRes, err error)
	MaxSort(ctx context.Context, req *genCodes.MaxSortReq) (res *genCodes.MaxSortRes, err error)
	Status(ctx context.Context, req *genCodes.StatusReq) (res *genCodes.StatusRes, err error)
	Selects(ctx context.Context, req *genCodes.SelectsReq) (res *genCodes.SelectsRes, err error)
	TableSelect(ctx context.Context, req *genCodes.TableSelectReq) (res *genCodes.TableSelectRes, err error)
	ColumnSelect(ctx context.Context, req *genCodes.ColumnSelectReq) (res *genCodes.ColumnSelectRes, err error)
	ColumnList(ctx context.Context, req *genCodes.ColumnListReq) (res *genCodes.ColumnListRes, err error)
	Preview(ctx context.Context, req *genCodes.PreviewReq) (res *genCodes.PreviewRes, err error)
	Build(ctx context.Context, req *genCodes.BuildReq) (res *genCodes.BuildRes, err error)
}
