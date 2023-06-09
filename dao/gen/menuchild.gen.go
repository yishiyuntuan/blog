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

func newMenuchild(db *gorm.DB, opts ...gen.DOOption) menuchild {
	_menuchild := menuchild{}

	_menuchild.menuchildDo.UseDB(db, opts...)
	_menuchild.menuchildDo.UseModel(&entity.Menuchild{})

	tableName := _menuchild.menuchildDo.TableName()
	_menuchild.ALL = field.NewAsterisk(tableName)
	_menuchild.ID = field.NewUint64(tableName, "id")
	_menuchild.Sort = field.NewUint64(tableName, "sort")
	_menuchild.Name = field.NewString(tableName, "name")
	_menuchild.Ename = field.NewString(tableName, "ename")
	_menuchild.Icon = field.NewString(tableName, "icon")
	_menuchild.Link = field.NewString(tableName, "link")
	_menuchild.ParentID = field.NewUint64(tableName, "parent_id")

	_menuchild.fillFieldMap()

	return _menuchild
}

type menuchild struct {
	menuchildDo

	ALL      field.Asterisk
	ID       field.Uint64
	Sort     field.Uint64 // 排序字段
	Name     field.String // 菜单名
	Ename    field.String // 英文名
	Icon     field.String // 图标svg格式
	Link     field.String // 路由名
	ParentID field.Uint64 // 父级id

	fieldMap map[string]field.Expr
}

func (m menuchild) Table(newTableName string) *menuchild {
	m.menuchildDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m menuchild) As(alias string) *menuchild {
	m.menuchildDo.DO = *(m.menuchildDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *menuchild) updateTableName(table string) *menuchild {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint64(table, "id")
	m.Sort = field.NewUint64(table, "sort")
	m.Name = field.NewString(table, "name")
	m.Ename = field.NewString(table, "ename")
	m.Icon = field.NewString(table, "icon")
	m.Link = field.NewString(table, "link")
	m.ParentID = field.NewUint64(table, "parent_id")

	m.fillFieldMap()

	return m
}

func (m *menuchild) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *menuchild) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 7)
	m.fieldMap["id"] = m.ID
	m.fieldMap["sort"] = m.Sort
	m.fieldMap["name"] = m.Name
	m.fieldMap["ename"] = m.Ename
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["link"] = m.Link
	m.fieldMap["parent_id"] = m.ParentID
}

func (m menuchild) clone(db *gorm.DB) menuchild {
	m.menuchildDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m menuchild) replaceDB(db *gorm.DB) menuchild {
	m.menuchildDo.ReplaceDB(db)
	return m
}

type menuchildDo struct{ gen.DO }

type IMenuchildDo interface {
	gen.SubQuery
	Debug() IMenuchildDo
	WithContext(ctx context.Context) IMenuchildDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMenuchildDo
	WriteDB() IMenuchildDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMenuchildDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMenuchildDo
	Not(conds ...gen.Condition) IMenuchildDo
	Or(conds ...gen.Condition) IMenuchildDo
	Select(conds ...field.Expr) IMenuchildDo
	Where(conds ...gen.Condition) IMenuchildDo
	Order(conds ...field.Expr) IMenuchildDo
	Distinct(cols ...field.Expr) IMenuchildDo
	Omit(cols ...field.Expr) IMenuchildDo
	Join(table schema.Tabler, on ...field.Expr) IMenuchildDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMenuchildDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMenuchildDo
	Group(cols ...field.Expr) IMenuchildDo
	Having(conds ...gen.Condition) IMenuchildDo
	Limit(limit int) IMenuchildDo
	Offset(offset int) IMenuchildDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuchildDo
	Unscoped() IMenuchildDo
	Create(values ...*entity.Menuchild) error
	CreateInBatches(values []*entity.Menuchild, batchSize int) error
	Save(values ...*entity.Menuchild) error
	First() (*entity.Menuchild, error)
	Take() (*entity.Menuchild, error)
	Last() (*entity.Menuchild, error)
	Find() ([]*entity.Menuchild, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Menuchild, err error)
	FindInBatches(result *[]*entity.Menuchild, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Menuchild) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMenuchildDo
	Assign(attrs ...field.AssignExpr) IMenuchildDo
	Joins(fields ...field.RelationField) IMenuchildDo
	Preload(fields ...field.RelationField) IMenuchildDo
	FirstOrInit() (*entity.Menuchild, error)
	FirstOrCreate() (*entity.Menuchild, error)
	FindByPage(offset int, limit int) (result []*entity.Menuchild, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMenuchildDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m menuchildDo) Debug() IMenuchildDo {
	return m.withDO(m.DO.Debug())
}

func (m menuchildDo) WithContext(ctx context.Context) IMenuchildDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m menuchildDo) ReadDB() IMenuchildDo {
	return m.Clauses(dbresolver.Read)
}

func (m menuchildDo) WriteDB() IMenuchildDo {
	return m.Clauses(dbresolver.Write)
}

func (m menuchildDo) Session(config *gorm.Session) IMenuchildDo {
	return m.withDO(m.DO.Session(config))
}

func (m menuchildDo) Clauses(conds ...clause.Expression) IMenuchildDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m menuchildDo) Returning(value interface{}, columns ...string) IMenuchildDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m menuchildDo) Not(conds ...gen.Condition) IMenuchildDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m menuchildDo) Or(conds ...gen.Condition) IMenuchildDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m menuchildDo) Select(conds ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m menuchildDo) Where(conds ...gen.Condition) IMenuchildDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m menuchildDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMenuchildDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m menuchildDo) Order(conds ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m menuchildDo) Distinct(cols ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m menuchildDo) Omit(cols ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m menuchildDo) Join(table schema.Tabler, on ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m menuchildDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m menuchildDo) RightJoin(table schema.Tabler, on ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m menuchildDo) Group(cols ...field.Expr) IMenuchildDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m menuchildDo) Having(conds ...gen.Condition) IMenuchildDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m menuchildDo) Limit(limit int) IMenuchildDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m menuchildDo) Offset(offset int) IMenuchildDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m menuchildDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuchildDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m menuchildDo) Unscoped() IMenuchildDo {
	return m.withDO(m.DO.Unscoped())
}

func (m menuchildDo) Create(values ...*entity.Menuchild) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m menuchildDo) CreateInBatches(values []*entity.Menuchild, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m menuchildDo) Save(values ...*entity.Menuchild) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m menuchildDo) First() (*entity.Menuchild, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menuchild), nil
	}
}

func (m menuchildDo) Take() (*entity.Menuchild, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menuchild), nil
	}
}

func (m menuchildDo) Last() (*entity.Menuchild, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menuchild), nil
	}
}

func (m menuchildDo) Find() ([]*entity.Menuchild, error) {
	result, err := m.DO.Find()
	return result.([]*entity.Menuchild), err
}

func (m menuchildDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Menuchild, err error) {
	buf := make([]*entity.Menuchild, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m menuchildDo) FindInBatches(result *[]*entity.Menuchild, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m menuchildDo) Attrs(attrs ...field.AssignExpr) IMenuchildDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m menuchildDo) Assign(attrs ...field.AssignExpr) IMenuchildDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m menuchildDo) Joins(fields ...field.RelationField) IMenuchildDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m menuchildDo) Preload(fields ...field.RelationField) IMenuchildDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m menuchildDo) FirstOrInit() (*entity.Menuchild, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menuchild), nil
	}
}

func (m menuchildDo) FirstOrCreate() (*entity.Menuchild, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Menuchild), nil
	}
}

func (m menuchildDo) FindByPage(offset int, limit int) (result []*entity.Menuchild, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m menuchildDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m menuchildDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m menuchildDo) Delete(models ...*entity.Menuchild) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *menuchildDo) withDO(do gen.Dao) *menuchildDo {
	m.DO = *do.(*gen.DO)
	return m
}
