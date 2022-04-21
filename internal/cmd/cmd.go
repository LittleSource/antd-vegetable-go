package cmd

import (
	"context"

	"antd-vegetable-go/internal/controller"
	"antd-vegetable-go/internal/middleware/mongodb"

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
			mongodb.InitMongodb()
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, ghttp.MiddlewareCORS)
				group.Bind(
					controller.Goods,
				)
			})
			s.Run()
			return nil
		},
	}
)
