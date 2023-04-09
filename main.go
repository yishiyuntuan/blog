package main

import (
	"github.com/kataras/iris/v12"

	"blog/api/v1/routes"
	"blog/dao/mapper"
	"blog/middleware/logger"
	"blog/service"
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
	app := iris.Default()
	app.UseRouter(logger.LoggerHandler)
	routes.InitRouterV1(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
func DI(app *iris.Application) {
	app.UseRouter()

	// 依赖注入
	app.RegisterDependency(service.NewService[service.ArticleServiceImpl](service.WithArticleDao(mapper.NewArticleDao())))
	app.RegisterDependency(service.NewService[service.CategoryServiceImpl](service.WithCategoryDao(mapper.NewCategoryDao())))
	// app.RegisterDependency(service.NewService[service.MenuChildServiceImpl](service.WithMenuChildDao(mapper.NewMenuDao(gen.Menuchild))))
	app.RegisterDependency(service.NewService[service.TagServiceImpl](service.WithTagDao(mapper.NewTagDao())))

	app.RegisterDependency(service.NewService[service.UserServiceImpl](service.WithUserDao(mapper.NewUserDao())))

}
