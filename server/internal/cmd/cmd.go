package cmd

import (
	"context"
	"server/internal/router"
	"server/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindMiddleware("/*any", []ghttp.HandlerFunc{
				service.Middleware().Ctx,
				service.Middleware().Cors,
				service.Middleware().ResponseHandler,
			}...)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 注册后台路由
				router.Admin(ctx, group)

				// 注册Api路由
				router.Api(ctx, group)
			})
			s.Run()
			return nil
		},
	}
)
