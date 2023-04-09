package service

import (
	"blog/model/dto"
	. "blog/model/entity"
)

type Option func(opts any)

type ArticleService interface {
	GetsArticleList(cid, mid, tid string, size, num int) ([]*Article, int64)
	GetArticle(id uint64) *dto.ArticleContent
}

type UserService interface {
	// 判断用户是否存在
	IsExist(username string) bool
	Register(user User) error
}

type CategoryService interface {
	CategoryService(isShow bool, mid string) []*Category
}

type MenuChildService interface {
	GetMenuChilds(leavl uint64) []*Menuchild
}
type TagService interface {
	GetTagAll() []*Tags
	GetArticleTag(id uint64) []*Tags
}

type service interface {
	ArticleServiceImpl | UserServiceImpl | MenuChildServiceImpl | CategoryServiceImpl | TagServiceImpl
}

func NewService[T service](opt ...Option) T {
	s := new(T)
	for _, f := range opt {
		f(s)
	}
	return *s
}
