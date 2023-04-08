package mapper

import (
	"blog/dao/gen"

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
func NewArticleDao(dao *gen.ArticleExec, q *gen.Query) ArticleDao {
	return &articleDao{
		dao: dao,
		q:   q,
		sfg: new(singleflight.Group)}
}

type userDao struct {
	dao *gen.UserExec
	// cache cache.ArticleCache
	sfg *singleflight.Group
}

// NewUserDao creating the dao interface
func NewUserDao(dao *gen.UserExec) UserDao {
	return &userDao{
		dao: dao,
		sfg: new(singleflight.Group)}
}
