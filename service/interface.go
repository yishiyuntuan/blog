package service

import (
	. "blog/model/entity"

	"github.com/go-sonic/sonic/model/entity"
)

type Option func(opts any)

type ArticleService interface {
	GetsArticleList(cid, mid, tid string, size, num int) ([]*Article, int64)
	GetArticle(id uint64) *Article
}

type CategoryService interface {
	GetCategory(isShow bool) *entity.Category
}

type UserService interface {
	// 判断用户是否存在
	IsExist(username string) bool
	Register(username string, password []byte) error
}
