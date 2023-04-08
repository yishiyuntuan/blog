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

func (ac ArticleController) Configure(r router.Party) {
	// 依赖注入
	dao := mapper.NewArticleDao(gen.Article, gen.Q)
	articleService := service.NewArticleService(service.WithArticleDao(dao))
	r.RegisterDependency(articleService)
	r.Use(iris.Compression)
	// ac.Service = articleService.NewArticleService(articleService.WithA(30000))
	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&ac)
	})
	// r.Get("/test",)
}

// GetList
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

	list, total := ac.Service.GetsArticleList(cid, mid, tid, size, num)

	data := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	return vo.Success(vo.WithData(data))
}

// GetBy
// @Summary 获取文章详情
// @Param   id     path    int     true        "文章ID"
// @Router /article/{id} [get]
func (ac ArticleController) GetBy(id uint64) *vo.Result {
	article := ac.Service.GetArticle(id)
	return vo.Success(vo.WithData(article))
}
