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

func newMessage(db *gorm.DB, opts ...gen.DOOption) message {
	_message := message{}

	_message.messageDo.UseDB(db, opts...)
	_message.messageDo.UseModel(&entity.Message{})

	tableName := _message.messageDo.TableName()
	_message.ALL = field.NewAsterisk(tableName)
	_message.ID = field.NewUint64(tableName, "id")
	_message.CreatedAt = field.NewTime(tableName, "created_at")
	_message.UpdatedAt = field.NewTime(tableName, "updated_at")
	_message.Name = field.NewString(tableName, "name")
	_message.Qq = field.NewString(tableName, "qq")
	_message.Email = field.NewString(tableName, "email")
	_message.Content = field.NewString(tableName, "content")
	_message.Like = field.NewInt64(tableName, "like")
	_message.Check = field.NewBool(tableName, "check")
	_message.Show = field.NewBool(tableName, "show")

	_message.fillFieldMap()

	return _message
}

type message struct {
	messageDo

	ALL       field.Asterisk
	ID        field.Uint64
	CreatedAt field.Time
	UpdatedAt field.Time
	Name      field.String // 昵称
	Qq        field.String // QQ
	Email     field.String // 邮箱
	Content   field.String // 内容
	Like      field.Int64  // 点赞
	Check     field.Bool   // 审核状态
	Show      field.Bool   // 显示

	fieldMap map[string]field.Expr
}

func (m message) Table(newTableName string) *message {
	m.messageDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m message) As(alias string) *message {
	m.messageDo.DO = *(m.messageDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *message) updateTableName(table string) *message {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint64(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.Name = field.NewString(table, "name")
	m.Qq = field.NewString(table, "qq")
	m.Email = field.NewString(table, "email")
	m.Content = field.NewString(table, "content")
	m.Like = field.NewInt64(table, "like")
	m.Check = field.NewBool(table, "check")
	m.Show = field.NewBool(table, "show")

	m.fillFieldMap()

	return m
}

func (m *message) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *message) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["name"] = m.Name
	m.fieldMap["qq"] = m.Qq
	m.fieldMap["email"] = m.Email
	m.fieldMap["content"] = m.Content
	m.fieldMap["like"] = m.Like
	m.fieldMap["check"] = m.Check
	m.fieldMap["show"] = m.Show
}

func (m message) clone(db *gorm.DB) message {
	m.messageDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m message) replaceDB(db *gorm.DB) message {
	m.messageDo.ReplaceDB(db)
	return m
}

type messageDo struct{ gen.DO }

type IMessageDo interface {
	gen.SubQuery
	Debug() IMessageDo
	WithContext(ctx context.Context) IMessageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessageDo
	WriteDB() IMessageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessageDo
	Not(conds ...gen.Condition) IMessageDo
	Or(conds ...gen.Condition) IMessageDo
	Select(conds ...field.Expr) IMessageDo
	Where(conds ...gen.Condition) IMessageDo
	Order(conds ...field.Expr) IMessageDo
	Distinct(cols ...field.Expr) IMessageDo
	Omit(cols ...field.Expr) IMessageDo
	Join(table schema.Tabler, on ...field.Expr) IMessageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessageDo
	Group(cols ...field.Expr) IMessageDo
	Having(conds ...gen.Condition) IMessageDo
	Limit(limit int) IMessageDo
	Offset(offset int) IMessageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageDo
	Unscoped() IMessageDo
	Create(values ...*entity.Message) error
	CreateInBatches(values []*entity.Message, batchSize int) error
	Save(values ...*entity.Message) error
	First() (*entity.Message, error)
	Take() (*entity.Message, error)
	Last() (*entity.Message, error)
	Find() ([]*entity.Message, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Message, err error)
	FindInBatches(result *[]*entity.Message, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Message) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessageDo
	Assign(attrs ...field.AssignExpr) IMessageDo
	Joins(fields ...field.RelationField) IMessageDo
	Preload(fields ...field.RelationField) IMessageDo
	FirstOrInit() (*entity.Message, error)
	FirstOrCreate() (*entity.Message, error)
	FindByPage(offset int, limit int) (result []*entity.Message, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messageDo) Debug() IMessageDo {
	return m.withDO(m.DO.Debug())
}

func (m messageDo) WithContext(ctx context.Context) IMessageDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messageDo) ReadDB() IMessageDo {
	return m.Clauses(dbresolver.Read)
}

func (m messageDo) WriteDB() IMessageDo {
	return m.Clauses(dbresolver.Write)
}

func (m messageDo) Session(config *gorm.Session) IMessageDo {
	return m.withDO(m.DO.Session(config))
}

func (m messageDo) Clauses(conds ...clause.Expression) IMessageDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messageDo) Returning(value interface{}, columns ...string) IMessageDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messageDo) Not(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messageDo) Or(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messageDo) Select(conds ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messageDo) Where(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMessageDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m messageDo) Order(conds ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messageDo) Distinct(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messageDo) Omit(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messageDo) Join(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messageDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessageDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messageDo) Group(cols ...field.Expr) IMessageDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messageDo) Having(conds ...gen.Condition) IMessageDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messageDo) Limit(limit int) IMessageDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messageDo) Offset(offset int) IMessageDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessageDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messageDo) Unscoped() IMessageDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messageDo) Create(values ...*entity.Message) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messageDo) CreateInBatches(values []*entity.Message, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messageDo) Save(values ...*entity.Message) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messageDo) First() (*entity.Message, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Message), nil
	}
}

func (m messageDo) Take() (*entity.Message, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Message), nil
	}
}

func (m messageDo) Last() (*entity.Message, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Message), nil
	}
}

func (m messageDo) Find() ([]*entity.Message, error) {
	result, err := m.DO.Find()
	return result.([]*entity.Message), err
}

func (m messageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Message, err error) {
	buf := make([]*entity.Message, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messageDo) FindInBatches(result *[]*entity.Message, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messageDo) Attrs(attrs ...field.AssignExpr) IMessageDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messageDo) Assign(attrs ...field.AssignExpr) IMessageDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messageDo) Joins(fields ...field.RelationField) IMessageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messageDo) Preload(fields ...field.RelationField) IMessageDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messageDo) FirstOrInit() (*entity.Message, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Message), nil
	}
}

func (m messageDo) FirstOrCreate() (*entity.Message, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Message), nil
	}
}

func (m messageDo) FindByPage(offset int, limit int) (result []*entity.Message, count int64, err error) {
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

func (m messageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messageDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messageDo) Delete(models ...*entity.Message) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messageDo) withDO(do gen.Dao) *messageDo {
	m.DO = *do.(*gen.DO)
	return m
}
