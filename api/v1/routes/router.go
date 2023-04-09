package routes

import (
	controller2 "blog/controller"
	_ "blog/docs"
	"blog/middleware/logger"
	"github.com/arl/statsviz"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/monitor"
	"strings"
	"time"
)

func InitRouterV1(app *iris.Application) {
	app.ConfigureContainer(func(api *iris.APIContainer) {
		app.Get("/", func(ctx *context.Context) {
			ctx.Redirect("http://192.168.8.1:3000/debug/")
		})
		root := app.Party("/api/v1")
		root.PartyConfigure("/article", new(controller2.ArticleController))
		root.PartyConfigure("/user", new(controller2.UserController))
		root.PartyConfigure("/menu", new(controller2.MenuchildController))
		root.PartyConfigure("/category", new(controller2.CategoryController))
		root.PartyConfigure("/tag", new(controller2.TagController))
	})
	// 内存图
	Statsviz(app)
	// 监控
	Monitor(app)
}
func Monitor(app iris.Party) {
	// Initialize and start the monitor middleware.
	m := monitor.New(monitor.Options{
		RefreshInterval:     2 * time.Second,
		ViewRefreshInterval: 2 * time.Second,
		ViewTitle:           "MyServer Monitor",
	})
	// Manually stop monitoring on CMD/CTRL+C.
	iris.RegisterOnInterrupt(m.Stop)

	// Serve the actual server's process and operating system statistics as JSON.
	app.Post("/monitor", m.Stats)
	// Render with the default page.
	app.Get("/monitor", m.View)
}
func Statsviz(app iris.Party) {
	// Register a router wrapper for this one.
	statsvizPath := "/debug"
	serveRoot := statsviz.IndexAtRoot(statsvizPath)
	serveWS := statsviz.NewWsHandler(time.Second)
	app.UseRouter(func(ctx iris.Context) {
		// You can optimize this if branch, I leave it to you as an exercise.
		if strings.HasPrefix(ctx.Path(), statsvizPath+"/ws") {
			serveWS(ctx.ResponseWriter(), ctx.Request())
		} else if strings.HasPrefix(ctx.Path(), statsvizPath) {
			logger.Log.Info("Statsviz")
			serveRoot(ctx.ResponseWriter(), ctx.Request())
		} else {
			ctx.Next()
		}
	})
}
