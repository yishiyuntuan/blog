// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"blog/model/entity"
)

func newArticle_tags(db *gorm.DB, opts ...gen.DOOption) article_tags {
	_article_tags := article_tags{}

	_article_tags.article_tagsDo.UseDB(db, opts...)
	_article_tags.article_tagsDo.UseModel(&entity.Article_tags{})

	tableName := _article_tags.article_tagsDo.TableName()
	_article_tags.ALL = field.NewAsterisk(tableName)
	_article_tags.TagsID = field.NewUint64(tableName, "tags_id")
	_article_tags.ArticleID = field.NewUint64(tableName, "article_id")

	_article_tags.fillFieldMap()

	return _article_tags
}

type article_tags struct {
	article_tagsDo

	ALL       field.Asterisk
	TagsID    field.Uint64
	ArticleID field.Uint64

	fieldMap map[string]field.Expr
}

func (a article_tags) Table(newTableName string) *article_tags {
	a.article_tagsDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a article_tags) As(alias string) *article_tags {
	a.article_tagsDo.DO = *(a.article_tagsDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *article_tags) updateTableName(table string) *article_tags {
	a.ALL = field.NewAsterisk(table)
	a.TagsID = field.NewUint64(table, "tags_id")
	a.ArticleID = field.NewUint64(table, "article_id")

	a.fillFieldMap()

	return a
}

func (a *article_tags) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *article_tags) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 2)
	a.fieldMap["tags_id"] = a.TagsID
	a.fieldMap["article_id"] = a.ArticleID
}

func (a article_tags) clone(db *gorm.DB) article_tags {
	a.article_tagsDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a article_tags) replaceDB(db *gorm.DB) article_tags {
	a.article_tagsDo.ReplaceDB(db)
	return a
}

type article_tagsDo struct{ gen.DO }

type IArticle_tagsDo interface {
	gen.SubQuery
	Debug() IArticle_tagsDo
	WithContext(ctx context.Context) IArticle_tagsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticle_tagsDo
	WriteDB() IArticle_tagsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticle_tagsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticle_tagsDo
	Not(conds ...gen.Condition) IArticle_tagsDo
	Or(conds ...gen.Condition) IArticle_tagsDo
	Select(conds ...field.Expr) IArticle_tagsDo
	Where(conds ...gen.Condition) IArticle_tagsDo
	Order(conds ...field.Expr) IArticle_tagsDo
	Distinct(cols ...field.Expr) IArticle_tagsDo
	Omit(cols ...field.Expr) IArticle_tagsDo
	Join(table schema.Tabler, on ...field.Expr) IArticle_tagsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticle_tagsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticle_tagsDo
	Group(cols ...field.Expr) IArticle_tagsDo
	Having(conds ...gen.Condition) IArticle_tagsDo
	Limit(limit int) IArticle_tagsDo
	Offset(offset int) IArticle_tagsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticle_tagsDo
	Unscoped() IArticle_tagsDo
	Create(values ...*entity.Article_tags) error
	CreateInBatches(values []*entity.Article_tags, batchSize int) error
	Save(values ...*entity.Article_tags) error
	First() (*entity.Article_tags, error)
	Take() (*entity.Article_tags, error)
	Last() (*entity.Article_tags, error)
	Find() ([]*entity.Article_tags, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Article_tags, err error)
	FindInBatches(result *[]*entity.Article_tags, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Article_tags) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticle_tagsDo
	Assign(attrs ...field.AssignExpr) IArticle_tagsDo
	Joins(fields ...field.RelationField) IArticle_tagsDo
	Preload(fields ...field.RelationField) IArticle_tagsDo
	FirstOrInit() (*entity.Article_tags, error)
	FirstOrCreate() (*entity.Article_tags, error)
	FindByPage(offset int, limit int) (result []*entity.Article_tags, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticle_tagsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a article_tagsDo) Debug() IArticle_tagsDo {
	return a.withDO(a.DO.Debug())
}

func (a article_tagsDo) WithContext(ctx context.Context) IArticle_tagsDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a article_tagsDo) ReadDB() IArticle_tagsDo {
	return a.Clauses(dbresolver.Read)
}

func (a article_tagsDo) WriteDB() IArticle_tagsDo {
	return a.Clauses(dbresolver.Write)
}

func (a article_tagsDo) Session(config *gorm.Session) IArticle_tagsDo {
	return a.withDO(a.DO.Session(config))
}

func (a article_tagsDo) Clauses(conds ...clause.Expression) IArticle_tagsDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a article_tagsDo) Returning(value interface{}, columns ...string) IArticle_tagsDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a article_tagsDo) Not(conds ...gen.Condition) IArticle_tagsDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a article_tagsDo) Or(conds ...gen.Condition) IArticle_tagsDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a article_tagsDo) Select(conds ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a article_tagsDo) Where(conds ...gen.Condition) IArticle_tagsDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a article_tagsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IArticle_tagsDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a article_tagsDo) Order(conds ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a article_tagsDo) Distinct(cols ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a article_tagsDo) Omit(cols ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a article_tagsDo) Join(table schema.Tabler, on ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a article_tagsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a article_tagsDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a article_tagsDo) Group(cols ...field.Expr) IArticle_tagsDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a article_tagsDo) Having(conds ...gen.Condition) IArticle_tagsDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a article_tagsDo) Limit(limit int) IArticle_tagsDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a article_tagsDo) Offset(offset int) IArticle_tagsDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a article_tagsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticle_tagsDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a article_tagsDo) Unscoped() IArticle_tagsDo {
	return a.withDO(a.DO.Unscoped())
}

func (a article_tagsDo) Create(values ...*entity.Article_tags) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a article_tagsDo) CreateInBatches(values []*entity.Article_tags, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a article_tagsDo) Save(values ...*entity.Article_tags) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a article_tagsDo) First() (*entity.Article_tags, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Article_tags), nil
	}
}

func (a article_tagsDo) Take() (*entity.Article_tags, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Article_tags), nil
	}
}

func (a article_tagsDo) Last() (*entity.Article_tags, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Article_tags), nil
	}
}

func (a article_tagsDo) Find() ([]*entity.Article_tags, error) {
	result, err := a.DO.Find()
	return result.([]*entity.Article_tags), err
}

func (a article_tagsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Article_tags, err error) {
	buf := make([]*entity.Article_tags, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a article_tagsDo) FindInBatches(result *[]*entity.Article_tags, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a article_tagsDo) Attrs(attrs ...field.AssignExpr) IArticle_tagsDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a article_tagsDo) Assign(attrs ...field.AssignExpr) IArticle_tagsDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a article_tagsDo) Joins(fields ...field.RelationField) IArticle_tagsDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a article_tagsDo) Preload(fields ...field.RelationField) IArticle_tagsDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a article_tagsDo) FirstOrInit() (*entity.Article_tags, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Article_tags), nil
	}
}

func (a article_tagsDo) FirstOrCreate() (*entity.Article_tags, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Article_tags), nil
	}
}

func (a article_tagsDo) FindByPage(offset int, limit int) (result []*entity.Article_tags, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a article_tagsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a article_tagsDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a article_tagsDo) Delete(models ...*entity.Article_tags) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *article_tagsDo) withDO(do gen.Dao) *article_tagsDo {
	a.DO = *do.(*gen.DO)
	return a
}
