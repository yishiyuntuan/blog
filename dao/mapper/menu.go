package mapper

import (
	"blog/dao/gen"
	"blog/middleware/logger"
	"blog/model/entity"

	"golang.org/x/sync/singleflight"
)

type meunChildDao struct {
	dao *gen.MenuchildExec
	sfg *singleflight.Group
}

func NewMenuDao() MenuDao {
	return &meunChildDao{
		dao: gen.Menuchild,
		sfg: new(singleflight.Group)}
}
func (m meunChildDao) GetMenus(leavl uint64) []*entity.Menuchild {
	find, err := m.dao.Select().Where(m.dao.ParentID.Eq(leavl)).Find()
	if err != nil {
		logger.Log.Error(err)
		return nil
	}
	return find
}
