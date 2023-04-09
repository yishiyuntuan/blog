package service

import (
	"blog/dao/mapper"
	"blog/model/entity"

	"github.com/tdewolff/parse/v2/strconv"
)

type CategoryServiceImpl struct {
	categoryDao mapper.CategoryDao
}

func WithCategoryDao(categoryDao mapper.CategoryDao) Option {
	return func(u any) {
		impl, ok := u.(*CategoryServiceImpl)
		if ok {
			impl.categoryDao = categoryDao
		}
	}
}
func (c CategoryServiceImpl) CategoryService(isShow bool, mid string) []*entity.Category {
	menuId, _ := strconv.ParseUint([]byte(mid))
	return c.categoryDao.Category(isShow, menuId)
}
