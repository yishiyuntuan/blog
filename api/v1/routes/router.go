package routes

import (
	"blog/api/v1/controller"
	_ "blog/docs"
	"blog/middleware/logger"
	"strings"
	"time"

	"github.com/arl/statsviz"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/monitor"
)

func InitRouterV2(app *iris.Application) {
	SwaggerAPI(app)
	Statsviz(app)
	root := app.Party("/")
	Monitor(root)
	root.PartyConfigure("/article", new(controller.ArticleController))
}

func SwaggerAPI(app *iris.Application) {
	// Configure the swagger UI page.
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("http://localhost:3000/swagger/doc.json"),
		// The url pointing to API definition.
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
	)
	// Register on http://localhost:3000/swagger
	app.Get("/swagger", swaggerUI)
	// And http://localhost:3000/swagger/index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)
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

func Statsviz(app *iris.Application) {
	statsvizPath := "/debug"
	serveRoot := statsviz.IndexAtRoot(statsvizPath)
	serveWS := statsviz.NewWsHandler(time.Second)
	logger.Log.Debug("Statsviz .....")
	app.UseRouter(func(ctx iris.Context) {
		// You can optimize this if branch, I leave it to you as an exercise.
		if strings.HasPrefix(ctx.Path(), statsvizPath+"/ws") {
			serveWS(ctx.ResponseWriter(), ctx.Request())
		} else if strings.HasPrefix(ctx.Path(), statsvizPath) {
			serveRoot(ctx.ResponseWriter(), ctx.Request())
		} else {
			ctx.Next()
		}
	})
}
