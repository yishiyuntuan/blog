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

func newCategory(db *gorm.DB, opts ...gen.DOOption) category {
	_category := category{}

	_category.categoryDo.UseDB(db, opts...)
	_category.categoryDo.UseModel(&entity.Category{})

	tableName := _category.categoryDo.TableName()
	_category.ALL = field.NewAsterisk(tableName)
	_category.ID = field.NewUint64(tableName, "id")
	_category.Name = field.NewString(tableName, "name")
	_category.Mid = field.NewUint64(tableName, "mid")
	_category.Homeshow = field.NewBool(tableName, "homeshow")

	_category.fillFieldMap()

	return _category
}

type category struct {
	categoryDo

	ALL      field.Asterisk
	ID       field.Uint64
	Name     field.String // 分类名
	Mid      field.Uint64 // 菜单子项ID
	Homeshow field.Bool   // 主页是否显示

	fieldMap map[string]field.Expr
}

func (c category) Table(newTableName string) *category {
	c.categoryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c category) As(alias string) *category {
	c.categoryDo.DO = *(c.categoryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *category) updateTableName(table string) *category {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewUint64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Mid = field.NewUint64(table, "mid")
	c.Homeshow = field.NewBool(table, "homeshow")

	c.fillFieldMap()

	return c
}

func (c *category) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *category) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["mid"] = c.Mid
	c.fieldMap["homeshow"] = c.Homeshow
}

func (c category) clone(db *gorm.DB) category {
	c.categoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c category) replaceDB(db *gorm.DB) category {
	c.categoryDo.ReplaceDB(db)
	return c
}

type categoryDo struct{ gen.DO }

type ICategoryDo interface {
	gen.SubQuery
	Debug() ICategoryDo
	WithContext(ctx context.Context) ICategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICategoryDo
	WriteDB() ICategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICategoryDo
	Not(conds ...gen.Condition) ICategoryDo
	Or(conds ...gen.Condition) ICategoryDo
	Select(conds ...field.Expr) ICategoryDo
	Where(conds ...gen.Condition) ICategoryDo
	Order(conds ...field.Expr) ICategoryDo
	Distinct(cols ...field.Expr) ICategoryDo
	Omit(cols ...field.Expr) ICategoryDo
	Join(table schema.Tabler, on ...field.Expr) ICategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICategoryDo
	Group(cols ...field.Expr) ICategoryDo
	Having(conds ...gen.Condition) ICategoryDo
	Limit(limit int) ICategoryDo
	Offset(offset int) ICategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryDo
	Unscoped() ICategoryDo
	Create(values ...*entity.Category) error
	CreateInBatches(values []*entity.Category, batchSize int) error
	Save(values ...*entity.Category) error
	First() (*entity.Category, error)
	Take() (*entity.Category, error)
	Last() (*entity.Category, error)
	Find() ([]*entity.Category, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Category, err error)
	FindInBatches(result *[]*entity.Category, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Category) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICategoryDo
	Assign(attrs ...field.AssignExpr) ICategoryDo
	Joins(fields ...field.RelationField) ICategoryDo
	Preload(fields ...field.RelationField) ICategoryDo
	FirstOrInit() (*entity.Category, error)
	FirstOrCreate() (*entity.Category, error)
	FindByPage(offset int, limit int) (result []*entity.Category, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c categoryDo) Debug() ICategoryDo {
	return c.withDO(c.DO.Debug())
}

func (c categoryDo) WithContext(ctx context.Context) ICategoryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoryDo) ReadDB() ICategoryDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoryDo) WriteDB() ICategoryDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoryDo) Session(config *gorm.Session) ICategoryDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoryDo) Clauses(conds ...clause.Expression) ICategoryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoryDo) Returning(value interface{}, columns ...string) ICategoryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoryDo) Not(conds ...gen.Condition) ICategoryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoryDo) Or(conds ...gen.Condition) ICategoryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoryDo) Select(conds ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoryDo) Where(conds ...gen.Condition) ICategoryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICategoryDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c categoryDo) Order(conds ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoryDo) Distinct(cols ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoryDo) Omit(cols ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoryDo) Join(table schema.Tabler, on ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoryDo) RightJoin(table schema.Tabler, on ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoryDo) Group(cols ...field.Expr) ICategoryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoryDo) Having(conds ...gen.Condition) ICategoryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoryDo) Limit(limit int) ICategoryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoryDo) Offset(offset int) ICategoryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoryDo) Unscoped() ICategoryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoryDo) Create(values ...*entity.Category) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoryDo) CreateInBatches(values []*entity.Category, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoryDo) Save(values ...*entity.Category) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoryDo) First() (*entity.Category, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Take() (*entity.Category, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Last() (*entity.Category, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Find() ([]*entity.Category, error) {
	result, err := c.DO.Find()
	return result.([]*entity.Category), err
}

func (c categoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Category, err error) {
	buf := make([]*entity.Category, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoryDo) FindInBatches(result *[]*entity.Category, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoryDo) Attrs(attrs ...field.AssignExpr) ICategoryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoryDo) Assign(attrs ...field.AssignExpr) ICategoryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoryDo) Joins(fields ...field.RelationField) ICategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoryDo) Preload(fields ...field.RelationField) ICategoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoryDo) FirstOrInit() (*entity.Category, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) FirstOrCreate() (*entity.Category, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) FindByPage(offset int, limit int) (result []*entity.Category, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c categoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoryDo) Delete(models ...*entity.Category) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoryDo) withDO(do gen.Dao) *categoryDo {
	c.DO = *do.(*gen.DO)
	return c
}