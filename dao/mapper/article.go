package mapper

import (
	"blog/dao/gen"
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
	// cache cache.ArticleCache
	sfg *singleflight.Group
}

// NewArticleDao creating the dao interface
func NewArticleDao(dao *gen.ArticleExec) ArticleDao {
	return &articleDao{
		dao: dao,
		sfg: new(singleflight.Group)}
}

func (a articleDao) GetList(cid, mid, tid string, size, num int) ([]*entity.Article, int64) {
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
