package middleware

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"server/internal/consts"
	"server/internal/model"
	"server/internal/service"
	"server/library/contexts"
	"server/library/jwt"
	"server/library/response"
	"strings"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Cors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// 错误状态码接管
	switch r.Response.Status {
	case 403:
		r.Response.Writeln("403 - deny")
		return
	case 404:
		r.Response.Writeln("404 - page not fund")
		return
	}

	// 已存在响应
	contentType := response.GetContentType(r)
	if contentType != consts.HTTPContentTypeStream && r.Response.BufferLength() > 0 { // && contexts.Get(r.Context()).Response != nil
		return
	}

	switch contentType {
	case consts.HTTPContentTypeHtml:
		response.HtmlReturn(r)
		return
	case consts.HTTPContentTypeXml:
		response.XmlReturn(r)
		return
	case consts.HTTPContentTypeStream:
	default:
		response.JsonReturn(r)
	}
}

// Ctx 初始化请求上下文
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	data := make(g.Map)
	if s.isMtForm(r) {
		data["request.body"] = gjson.New(nil)
	} else {
		data["request.body"] = gjson.New(r.GetBodyString())
	}
	contexts.Init(r, &model.Context{
		Data:   data,
		Module: s.getModule(r.URL.Path),
	})
	r.Middleware.Next()
}

func (s *sMiddleware) AdminAuth(r *ghttp.Request) {
	ctx := r.Context()
	jwtTokenConf, err := service.SysConfig().GetAdminJwtToken(ctx)
	if err != nil {
		response.JsonExit(r, gcode.CodeMissingConfiguration.Code(), "获取登陆token配置失败")
		return
	}
	userId, err := jwt.ParseUserId(s.getAuthorization(r), jwtTokenConf.SecretKey)
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}
	if userId == 0 {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "请重新授权登陆")
		return
	}
	user, err := service.AdminMember().GetOneById(ctx, userId)
	if err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}
	if user == nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), "请重新授权登陆")
		return
	}
	if user.Status != consts.DictStatusEnabled {
		response.JsonExit(r, gcode.CodeBusinessValidationFailed.Code(), "您已被禁用")
		return
	}

	ctxModel := contexts.Get(ctx)
	if ctxModel == nil {
		response.JsonExit(r, gcode.CodeUnknown.Code(), "获取ctxModel失败")
		return
	}
	ctxModel.UserId = userId
	ctxModel.AdminUser = user

	// 获取角色
	// todo
	// 验证路由访问权限

	r.Middleware.Next()
}

func (s *sMiddleware) isMtForm(r *ghttp.Request) bool {
	return gstr.Contains(r.GetHeader("Content-Type"), "multipart/form-data")
}
func (s *sMiddleware) getModule(path string) (module string) {
	slice := strings.Split(path, "/")
	if len(slice) < 2 {
		module = consts.ApiPrefix
		return
	}
	if slice[1] == "" {
		module = consts.ApiPrefix
		return
	}
	return slice[1]
}
func (s *sMiddleware) getAuthorization(r *ghttp.Request) string {
	// 默认从请求头获取
	var authorization = r.Header.Get("Authorization")

	// 如果请求头不存在则从get参数获取
	if authorization == "" {
		return r.Get("authorization").String()
	}
	return gstr.Replace(authorization, "Bearer ", "")
}
