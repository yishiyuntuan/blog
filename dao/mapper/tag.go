package mapper

import (
	"blog/dao/gen"
	"blog/model/entity"

	"golang.org/x/sync/singleflight"
)

type tagDao struct {
	dao *gen.TagsExec
	at  *gen.ArticleTagsExec
	sfg *singleflight.Group
}

func NewTagDao() TagDao {
	return &tagDao{
		dao: gen.Tags,
		at:  gen.Article_tags,
		sfg: new(singleflight.Group),
	}
}

func (t tagDao) TagList() []*entity.Tags {
	find, err := t.dao.Find()
	if err != nil {
		return nil
	}
	return find
}

func (t tagDao) ArticleTag(id uint64) []*entity.Tags {
	//find, err := t.dao.Select(t.dao.ID, t.dao.Color, t.dao.Logo, t.dao.Name, t.dao.ALL).Join(t.at, t.at.TagsID.EqCol(t.dao.ID)).Where(t.at.ArticleID.Eq(id)).Find()
	find, err := t.dao.Join(t.at, t.at.TagsID.EqCol(t.dao.ID)).Where(t.at.ArticleID.Eq(id)).Find()
	if err != nil {
		return nil
	}
	return find

}
