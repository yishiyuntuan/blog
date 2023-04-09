package mapper

import (
	"blog/dao/gen"
	"blog/middleware/logger"
	"blog/model/entity"

	"github.com/tdewolff/parse/v2/strconv"
	"golang.org/x/sync/singleflight"
)

type ArticleDaoImpl struct {
	dao *gen.ArticleExec
	q   *gen.Query
	// cache cache.ArticleCache
	sfg *singleflight.Group
}

// NewArticleDao creating the dao interface
func NewArticleDao() ArticleDao {
	return &ArticleDaoImpl{
		dao: gen.Article,
		q:   gen.Q,
		sfg: new(singleflight.Group)}
}
func (a ArticleDaoImpl) GetListByCid(cid string, size, num int) ([]*entity.Article, int64) {
	parseUint, _ := strconv.ParseUint([]byte(cid))
	page, n, err := a.dao.Select(
		a.dao.ID, a.dao.CreatedAt, a.dao.UpdatedAt,
		a.dao.DeletedAt, a.dao.Title, a.dao.Img,
		a.dao.DeletedAt, a.dao.Cid).
		Where(a.dao.Cid.Eq(parseUint)).FindByPage(size*(num-1), size)
	if err != nil {
		logger.Log.Error(err)
		return nil, 0
	}
	return page, n
}

func (a ArticleDaoImpl) GetListByTid(tid string, size, num int) ([]*entity.Article, int64) {
	parseUint, _ := strconv.ParseUint([]byte(tid))
	result, count, err := a.dao.Select(
		a.dao.ID, a.dao.CreatedAt, a.dao.UpdatedAt, a.dao.DeletedAt,
		a.dao.Title, a.dao.Img, a.dao.DeletedAt, a.dao.Cid).
		Join(a.q.Article_tags, a.dao.ID.EqCol(a.q.Article_tags.ArticleID)).
		Where(a.q.Article_tags.TagsID.Eq(parseUint)).
		FindByPage(size*(num-1), size)
	if err != nil {
		logger.Log.Error(err)
		return nil, 0
	}
	return result, count
}

func (a ArticleDaoImpl) GetListByMid(mid string, size, num int) ([]*entity.Article, int64) {
	parseUint, _ := strconv.ParseUint([]byte(mid))
	m := a.q.Menuchild.As("m")
	c := a.q.Category.As("c")

	// result, count, err :=
	result, count, err := a.dao.Select(
		a.dao.ID, a.dao.CreatedAt, a.dao.UpdatedAt, a.dao.DeletedAt,
		a.dao.Title, a.dao.Img, a.dao.DeletedAt, a.dao.Cid).
		Join(m).
		Join(c, m.ID.EqCol(c.Mid), c.ID.EqCol(a.dao.Cid)).
		Where(a.q.Menuchild.ID.Eq(parseUint)).FindByPage(size*(num-1), size)
	if err != nil {
		logger.Log.Error(err)
		return nil, 0
	}
	return result, count
}

func (a ArticleDaoImpl) GetListAll(size, num int) ([]*entity.Article, int64) {
	result, count, err := a.dao.Select(
		a.dao.ID, a.dao.CreatedAt, a.dao.UpdatedAt, a.dao.DeletedAt,
		a.dao.Title, a.dao.Img, a.dao.DeletedAt, a.dao.Cid).FindByPage(size*(num-1), size)
	if err != nil {
		logger.Log.Error(err)
		return nil, 0
	}
	return result, count
}

func (a ArticleDaoImpl) GetArticlePathByID(id uint64) string {
	first, err := a.q.Article_path.Where(a.q.Article_path.ID.Eq(int64(id))).First()
	if err != nil {
		logger.Log.Error(err)
		return ""
	}
	return first.LocalPath
}

func (a ArticleDaoImpl) GetArticleInfoByID(id uint64) *entity.Article {
	first, err := a.dao.Where(a.dao.ID.Eq(id)).First()
	if err != nil {
		logger.Log.Error(err)
		return nil
	}
	return first
}
