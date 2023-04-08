package controller

import "blog/service"

type ArticleController struct {
	Service service.ArticleService
}

func NewArticleController(s service.ArticleService) *ArticleController {
	return &ArticleController{
		Service: s,
	}
}

type UserController struct {
	Service service.UserService
}
