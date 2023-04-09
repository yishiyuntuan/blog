package mapper

import (
	"blog/model/entity"
)

// ArticleDao defining the dao interface
type ArticleDao interface {
	GetListByCid(cid string, size, num int) ([]*entity.Article, int64)
	GetListByTid(tid string, size, num int) ([]*entity.Article, int64)
	//GetArticleByID(id uint64) *entity.Article
	GetListByMid(mid string, size int, num int) ([]*entity.Article, int64)
	GetListAll(size int, num int) ([]*entity.Article, int64)
	GetArticlePathByID(id uint64) string
	GetArticleInfoByID(id uint64) *entity.Article
}

// UserDao defining the dao interface
type UserDao interface {
	IsExist(username string) bool
	Register(user *entity.User) error
}

type MenuDao interface {
	GetMenus(leavl uint64) []*entity.Menuchild
}

type CategoryDao interface {
	Category(show bool, mid uint64) []*entity.Category
}
type TagDao interface {
	TagList() []*entity.Tags
	ArticleTag(id uint64) []*entity.Tags
}
