package admin

import (
	"context"
	"server/api/admin/adminLogin"
	"server/internal/service"
	"server/library/captcha"
)

func (c *ControllerAdminLogin) LoginCaptcha(ctx context.Context, req *adminLogin.LoginCaptchaReq) (res *adminLogin.LoginCaptchaRes, err error) {
	loginConf, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}
	cid, base64 := captcha.Generate(ctx, loginConf.CaptchaType)
	res = &adminLogin.LoginCaptchaRes{Cid: cid, Base64: base64}
	return
}
