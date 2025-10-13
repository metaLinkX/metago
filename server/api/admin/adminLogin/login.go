package adminLogin

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model"
)

// LoginCaptchaReq 获取登录验证码
type LoginCaptchaReq struct {
	g.Meta `path:"/login/captcha" method:"get" tags:"后台基础" summary:"获取登录验证码"`
}

type LoginCaptchaRes struct {
	Cid    string `json:"cid" dc:"验证码ID"`
	Base64 string `json:"base64" dc:"验证码"`
}

// LoginConfigReq 获取登录配置
type LoginConfigReq struct {
	g.Meta `path:"/login/config" method:"post" tags:"后台基础" summary:"获取登录配置"`
}

type LoginConfigRes struct {
	*model.LoginConfig
}

// AccountLoginReq 提交账号登录
type AccountLoginReq struct {
	g.Meta   `path:"/login/accountLogin" method:"post" tags:"后台基础" summary:"账号登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Cid      string `json:"cid"  dc:"验证码ID"`
	Code     string `json:"code" dc:"验证码"`
	IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}

type AccountLoginRes struct {
	Id       int64  `json:"id"              dc:"用户ID"`
	Username string `json:"username"        dc:"用户名"`
	Token    string `json:"token"           dc:"登录token"`
	Expires  int64  `json:"expires"         dc:"登录有效期"`
}
