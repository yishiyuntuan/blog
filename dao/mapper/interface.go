package mapper

import "blog/model/entity"

// ArticleDao defining the dao interface
type ArticleDao interface {
	GetList(cid, mid, tid string, size, num int) ([]*entity.Article, int64)
}
