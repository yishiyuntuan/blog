package routes

import (
	"blog/controller"
	"blog/logger"
	"net/http"
	"time"

	_ "blog/docs"

	"github.com/arl/statsviz"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/monitor"
)

func InitRouterV2(app *iris.Application) {
	SwaggerAPI(app)
	root := app.Party("/")
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

func Monitor(app *iris.Application) {

	logger.Log.Info("Monitor")
	monitorParty := app.Party("monitor")
	{
		m := monitor.New(monitor.Options{
			RefreshInterval:     2 * time.Second,
			ViewRefreshInterval: 2 * time.Second,
			ViewTitle:           "MyServer Monitor",
		})
		// Manually stop monitoring on CMD/CTRL+C.
		iris.RegisterOnInterrupt(m.Stop)
		// Serve the actual server's process and operating system statistics as JSON.
		monitorParty.Post("/", m.Stats)
		// Render with the default page.
		monitorParty.Get("/", m.View)
		monitorParty.Any("/debug", func(context *context.Context) {
			context.Redirect("http://localhost:3001/debug/statsviz", 302)

		})
	}
}

func DebugStatisviz(app *iris.Application) {
	mux := http.NewServeMux()
	statsviz.Register(mux)
	statsSrv := &http.Server{Addr: ":3001", Handler: mux}
	logger.Log.Info("Point your browser to http://localhost:3001/debug" +
		"/statsviz\n")
	app.NewHost(statsSrv).ListenAndServe()

}
