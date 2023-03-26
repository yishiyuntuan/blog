package controller

import (
	"blog/dao/gen"
	"blog/dao/mapper"
	"blog/model/vo"
	"blog/service"
	"blog/util/tool"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type ArticleController struct {
	Ctx iris.Context
	S   service.ArticleService
}

func NewArticleController(s service.ArticleService) *ArticleController {
	return &ArticleController{
		S: s,
	}
}

func (ac ArticleController) Configure(r router.Party) {
	// 依赖注入
	dao := mapper.NewArticleDao(gen.Article)
	articleService := service.NewArticleService(service.WithArticleDao(dao))
	r.RegisterDependency(articleService)

	// ac.S = service.NewArticleService(service.WithA(30000))
	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&ac)
	})
	// r.Get("/test",)

}

/*
func (ac ArticleController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // 和你已知的任何标准 API  调用

	// 1-> 方法
	// 2-> 路径
	// 3-> 控制器函数的名称将被解析未一个处理程序 [ handler ]
	// 4-> 任何应该在 MyCustomHandler 之前运行的处理程序[ handlers ]
	// b.Handle("GET", "/something/{id:long}", "MyCustomHandler")
}
*/

// @Summary 获取文章列表
// @Param   cid         query    int     false        "分类ID"
// @Param   mid         query    int     false        "菜单ID"
// @Param   tid         query    int     false        "标签ID"
// @Param   pageSize    query    int     false        "分页大小"
// @Param   pageNum     query    int     false        "页码"
// @Router /article/list [get]
func (ac ArticleController) GetList(ctx iris.Context) *vo.Result {
	// 分页最大数,分页偏移量
	size, num := tool.PageTool(ctx)
	cid := ctx.URLParam("cid") // 分类ID
	mid := ctx.URLParam("mid") // 菜单ID
	tid := ctx.URLParam("tid") // 标签ID

	list, _ := ac.S.GetsArticleList(cid, mid, tid, size, num)

	return vo.OK(vo.WithData(list))
}
