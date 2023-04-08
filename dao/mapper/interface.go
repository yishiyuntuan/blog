package mapper

import "blog/model/entity"

// ArticleDao defining the dao interface
type ArticleDao interface {
	GetListByCid(cid string, size, num int) ([]*entity.Article, int64)
	GetListByTid(tid string, size, num int) ([]*entity.Article, int64)
	GetArticleByID(id uint64) *entity.Article
}

// UserDao defining the dao interface
type UserDao interface {
	IsExist(username string) bool
	Register(user *entity.User) error
}
