package main

import (
	"blog/routes"

	"github.com/kataras/iris/v12"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	// model.InitDb()
	// db.InitRedis()
	// routes.InitRouter()
	app := iris.Default()
	routes.Monitor(app)
	routes.InitRouterV2(app)
	routes.DebugStatisviz(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}