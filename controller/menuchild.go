package controller

import (
	"blog/dao/mapper"
	"blog/middleware/logger"
	"blog/model/vo"
	"blog/service"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type MenuchildController struct {
	service service.MenuChildService
}

func (m MenuchildController) Configure(r router.Party) {
	m.service = service.NewService[service.MenuChildServiceImpl](service.WithMenuChildDao(mapper.NewMenuDao()))

	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&m)
	})

	r.Get("/{leavl:path}", m.Menu)
}

// func (m MenuchildController) BeforeActivation(b mvc.BeforeActivation) {
// 	b.Handle("GET", "/{leavl:path}", "Menu")
// }

// Menu 获取菜单
// @Summary 获取菜单
// @Description 获取菜单
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param leavl path string true "菜单层级"
// @Success 200 {object} Result
// @Router /api/v1/menu/{leavl}/.../{leavl} [get]
func (m MenuchildController) Menu(ctx *context.Context) {
	path := strings.Split(ctx.Params().Get("leavl"), "/")
	mids := make([]uint64, len(path))
	for i, v := range path {
		parseUint, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			ctx.JSON(vo.Fail(vo.WithMessage("路径不存在")))
			return
		}
		// parseUint, _ := strconv.ParseUint([]byte(v))
		// logger.Log.Debug(parseUint)
		mids[i] = parseUint
	}

	logger.Log.Info(path)
	menu := m.service.GetMenuChilds(mids[len(mids)-1])
	if len(menu) == 0 {
		ctx.JSON(vo.Fail(vo.WithMessage("获取菜单失败")))
		return
	}
	// 分页最大数,分页偏移量
	ctx.JSON(vo.Success(vo.WithData(menu)))
}

func (m MenuchildController) Get() {
	logger.Log.Info("get")
}
