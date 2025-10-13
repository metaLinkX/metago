package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/consts"
	"server/library/stack"
	"server/library/system"
)

type Response struct {
	Code      int         `json:"code" example:"0" description:"状态码"`
	Message   string      `json:"message,omitempty" example:"操作成功" description:"提示消息"`
	Data      interface{} `json:"data,omitempty" description:"数据集"`
	Error     interface{} `json:"error,omitempty" description:"错误信息"`
	Timestamp int64       `json:"timestamp" example:"1757661092" description:"服务器时间戳"`
	TraceID   string      `json:"traceID" v:"0" example:"d1bb93048bc5c9164cdee845dcb7f820" description:"链路ID"`
}

func json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	res := &Response{
		Code:      code,
		Message:   message,
		Timestamp: gtime.Timestamp(),
		TraceID:   gctx.CtxId(r.Context()),
	}

	// 如果不是正常的返回，则将data转为error
	if gcode.CodeOK.Code() == code {
		res.Data = responseData
	} else {
		res.Error = responseData
	}

	// 清空响应
	r.Response.ClearBuffer()

	// 写入响应
	r.Response.WriteJson(res)

	// 加入到上下文 todo 用户hook 日志
}

func JsonReturn(r *ghttp.Request) {
	code, message, data := parseResponse(r)
	json(r, code, message, data)
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	json(r, code, message, data...)
	r.Exit()
}

func xml(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	res := &Response{
		Code:      code,
		Message:   message,
		Timestamp: gtime.Timestamp(),
		TraceID:   gctx.CtxId(r.Context()),
	}

	// 如果不是正常的返回，则将data转为error
	if gcode.CodeOK.Code() == code {
		res.Data = responseData
	} else {
		res.Error = responseData
	}

	// 清空响应
	r.Response.ClearBuffer()

	// 写入响应
	r.Response.WriteXml(gconv.Map(res))

	// 加入到上下文 todo 用户hook 日志
}
func XmlReturn(r *ghttp.Request) {
	code, message, data := parseResponse(r)
	xml(r, code, message, data)
}

// HtmlReturn html模板响应
func HtmlReturn(r *ghttp.Request) {
	code, message, resp := parseResponse(r)
	if code == gcode.CodeOK.Code() {
		return
	}
	r.Response.ClearBuffer()
	_ = r.Response.WriteTplContent(system.DefaultErrorTplContent(r.Context()), g.Map{"code": code, "message": message, "stack": resp})
}

func GetContentType(r *ghttp.Request) (contentType string) {
	contentType = r.Response.Header().Get("Content-Type")
	if contentType != "" {
		return
	}

	mime := gmeta.Get(r.GetHandlerResponse(), "mime").String()
	if mime == "" {
		contentType = consts.HTTPContentTypeJson
	} else {
		contentType = mime
	}
	return
}

func parseResponse(r *ghttp.Request) (code int, message string, resp interface{}) {
	ctx := r.Context()
	err := r.GetError()
	if err == nil {
		return gcode.CodeOK.Code(), "操作成功", r.GetHandlerResponse()
	}

	// 是否输出错误堆栈到页面
	message = gerror.Current(err).Error()
	if system.Debug(ctx) {
		resp = stack.ParseErrStack(err)
	}
	// todo
	code = gerror.Code(err).Code()
	// 记录异常日志
	// 如果你想对错误做不同的处理，可以通过定义不同的错误码来区分
	// 默认-1为安全可控错误码只记录文件日志，非-1为不可控错误，记录文件日志+服务日志并打印堆栈
	if code == gcode.CodeNil.Code() {
		g.Log().Stdout(false).Infof(ctx, "exception:%v", err)
	} else {
		g.Log().Errorf(ctx, "exception:%v", err)
	}
	return
}
