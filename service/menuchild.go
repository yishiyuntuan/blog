package service

import (
	"blog/dao/mapper"
	"blog/model/entity"
)

type MenuChildServiceImpl struct {
	// Service
	menuDao mapper.MenuDao
}

func WithMenuChildDao(menuDao mapper.MenuDao) Option {
	return func(u any) {
		impl, ok := u.(*MenuChildServiceImpl)
		if ok {
			impl.menuDao = menuDao
		}
	}
}
func (m MenuChildServiceImpl) GetMenuChilds(leavl uint64) []*entity.Menuchild {
	return m.menuDao.GetMenus(leavl)
}
