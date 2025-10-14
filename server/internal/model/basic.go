package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/model/entity"
)

type Context struct {
	Module    string // 应用模块 admin｜api｜home｜websocket
	UserId    int64  // 上下文用户信息
	AdminUser *entity.AdminMember
	ApiUser   *entity.AppMember
	Response  *Response // 请求响应
	Data      g.Map     // 自定kv变量 业务模块根据需要设置，不固定
}

// Response HTTP响应
type Response struct {
	Code      int         `json:"code" example:"0" description:"状态码"`
	Message   string      `json:"message,omitempty" example:"操作成功" description:"提示消息"`
	Data      interface{} `json:"data,omitempty" description:"数据集"`
	Error     interface{} `json:"error,omitempty" description:"错误信息"`
	Timestamp int64       `json:"timestamp" example:"1640966400" description:"服务器时间戳"`
	TraceID   string      `json:"traceID" v:"0" example:"d0bb93048bc5c9164cdee845dcb7f820" description:"链路ID"`
}
