package controller

import (
	"blog/dao/mapper"
	"blog/model/vo"
	"blog/service"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type CategoryController struct {
	service service.CategoryService
}

func (c CategoryController) Configure(r router.Party) {
	c.service = service.NewService[service.CategoryServiceImpl](service.WithCategoryDao(mapper.NewCategoryDao()))
	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&c)
	})
}

// GetCategory 获取分类
func (c CategoryController) GetCategory(ctx *context.Context) *vo.Result {
	show := ctx.URLParamBoolDefault("show", true)
	mid := ctx.URLParamDefault("mid", "0")
	list := c.service.CategoryService(show, mid)
	if len(list) == 0 {
		return vo.Fail(vo.WithMessage("分类不存在"))
	}
	return vo.Success(vo.WithData(list))
}
