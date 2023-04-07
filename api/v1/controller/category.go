package controller

import (
	"github.com/go-sonic/sonic/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type CategoryController struct {
	ctx iris.Context
	s   service.CategoryService
}

func (c CategoryController) Configure(parent router.Party) {
	// TODO implement me
	// panic("implement me")
}

// GetCategory 获取分类

// GetTags 获取全部标签
