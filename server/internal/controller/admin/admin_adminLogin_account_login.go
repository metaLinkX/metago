package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"server/api/admin/adminLogin"
	"server/internal/consts"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"server/library/captcha"
	"server/library/encrypt"
	"server/library/jwt"
)

func (c *ControllerAdminLogin) AccountLogin(ctx context.Context, req *adminLogin.AccountLoginReq) (res *adminLogin.AccountLoginRes, err error) {
	login, err := service.SysConfig().GetLogin(ctx)
	if err != nil {
		return
	}
	if !req.IsLock && login.CaptchaSwitch == consts.DictYes {
		// 校验 验证码
		if !captcha.Verify(req.Cid, req.Code) {
			err = gerror.New("验证码错误")
			return
		}
	}

	var mb *entity.AdminMember
	if err = dao.AdminMember.Ctx(ctx).Where("username", req.Username).Scan(&mb); err != nil {
		return
	}
	if mb == nil {
		err = gerror.New("用户名错误")
		return
	}
	if mb.Salt == "" {
		err = gerror.New("用户信息错误")
		return
	}
	if !encrypt.CheckPassword(req.Password, mb.PasswordHash, mb.Salt) {
		err = gerror.New("用户名或密码错误")
		return
	}

	if mb.Status != consts.DictStatusEnabled {
		err = gerror.New("账号已被禁用")
		return
	}

	//_, _, err = service.AdminMember().GetLoginRoleAndDept(ctx, mb.RoleId, mb.DeptId)
	//if err != nil {
	//	return res, err
	//}

	jwtConfig, err := service.SysConfig().GetAdminJwtToken(ctx)
	if err != nil {
		return res, err
	}

	token, err := jwt.IssueUserIdToken(mb.Id, jwtConfig.SecretKey, jwtConfig.Expires)
	if err != nil {
		return res, err
	}

	return &adminLogin.AccountLoginRes{
		Id:       mb.Id,
		Username: mb.Username,
		Token:    token,
		Expires:  jwtConfig.Expires,
	}, err
}
