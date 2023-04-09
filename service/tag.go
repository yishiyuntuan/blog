package service

import (
	"blog/dao/mapper"
	"blog/model/entity"
)

type TagServiceImpl struct {
	mapper.TagDao
}

func WithTagDao(tagDao mapper.TagDao) Option {
	return func(u any) {
		impl, ok := u.(*TagServiceImpl)
		if ok {
			impl.TagDao = tagDao
		}
	}
}

func (t TagServiceImpl) GetTagAll() []*entity.Tags {
	return t.TagList()
}

func (t TagServiceImpl) GetArticleTag(id uint64) []*entity.Tags {
	return t.ArticleTag(id)
}
