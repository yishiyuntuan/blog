package mapper_impl

import (
	"blog/dao/gen"
	"blog/dao/mapper"
	"blog/logger"
	"blog/model/entity"

	"github.com/tdewolff/parse/v2/strconv"
	"golang.org/x/sync/singleflight"
)

// 检测Member是否实现了IMember接口
// 1） _为了避免变量未使用编译的时候报错
// 2）_的类型为articleDao，接口的值为ArticleDao的地址，(nil)表示该地址为nil。
// var _ ArticleDao = (*articleDao)(nil)

type articleDao struct {
	dao *gen.ArticleExec
	q   *gen.Query
	// cache cache.ArticleCache
	sfg *singleflight.Group
}

// NewArticleDao creating the dao interface
func NewArticleDao(dao *gen.ArticleExec, q *gen.Query) mapper.ArticleDao {
	return &articleDao{
		dao: dao,
		q:   q,
		sfg: new(singleflight.Group)}
}

func (a articleDao) GetListByCid(cid string, size, num int) ([]*entity.Article, int64) {
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
func (a articleDao) GetListByTid(tid string, size, num int) ([]*entity.Article, int64) {
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

func (a articleDao) GetArticleByID(id uint64) *entity.Article {
	article, err := a.dao.Where(a.dao.ID.Eq(id)).First()
	if err != nil {
		logger.Log.Errorf("获取文章失败，id is %d, %v", id, err)
		return nil
	}
	return article
}
