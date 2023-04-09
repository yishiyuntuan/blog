package controller

import (
	"blog/dao/mapper"
	"blog/model/vo"
	"blog/service"
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type TagController struct {
	service service.TagService
}

func (t TagController) Configure(r router.Party) {
	t.service = service.NewService[service.TagServiceImpl](service.WithTagDao(mapper.NewTagDao()))

	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&t)
	})
}
func (t TagController) GetList() *vo.Result {
	list := t.service.GetTagAll()
	if len(list) == 0 {
		return vo.Fail(vo.WithMessage("查询失败"))
	}
	return vo.Success(vo.WithData(list))
}

// 获取文章的标签
func (t TagController) GetArticleTag(ctx iris.Context) *vo.Result {
	id := ctx.URLParamUint64("id")
	list := t.service.GetArticleTag(id)
	if len(list) == 0 {
		return vo.Fail(vo.WithMessage("查询失败"))
	}
	return vo.Success(vo.WithData(list))
}
