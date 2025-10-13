// Package router

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/consts"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(consts.ApiPrefix, func(group *ghttp.RouterGroup) {

	})
}
