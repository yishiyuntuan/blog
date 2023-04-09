package mapper

import (
	"blog/dao/gen"
	"blog/middleware/logger"
	"blog/model/entity"

	"golang.org/x/sync/singleflight"
)

type categoryDao struct {
	dao *gen.CategoryExec
	sfg *singleflight.Group
}

func NewCategoryDao() CategoryDao {
	return &categoryDao{
		dao: gen.Category,
		sfg: new(singleflight.Group),
	}
}
func (c categoryDao) Category(isShow bool, mid uint64) []*entity.Category {
	// TODO implement me
	find, err := c.dao.Where(c.dao.Homeshow.Is(isShow)).Where(c.dao.Mid.Eq(mid)).Find()
	if err != nil {
		logger.Log.Error(err)
		return nil
	}
	return find
}
