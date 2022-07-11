// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newSubjectField(db *gorm.DB) subjectField {
	_subjectField := subjectField{}

	_subjectField.subjectFieldDo.UseDB(db)
	_subjectField.subjectFieldDo.UseModel(&dao.SubjectField{})

	tableName := _subjectField.subjectFieldDo.TableName()
	_subjectField.ALL = field.NewField(tableName, "*")
	_subjectField.Sid = field.NewUint32(tableName, "field_sid")
	_subjectField.Tid = field.NewUint16(tableName, "field_tid")
	_subjectField.Tags = field.NewBytes(tableName, "field_tags")
	_subjectField.Rate1 = field.NewUint32(tableName, "field_rate_1")
	_subjectField.Rate2 = field.NewUint32(tableName, "field_rate_2")
	_subjectField.Rate3 = field.NewUint32(tableName, "field_rate_3")
	_subjectField.Rate4 = field.NewUint32(tableName, "field_rate_4")
	_subjectField.Rate5 = field.NewUint32(tableName, "field_rate_5")
	_subjectField.Rate6 = field.NewUint32(tableName, "field_rate_6")
	_subjectField.Rate7 = field.NewUint32(tableName, "field_rate_7")
	_subjectField.Rate8 = field.NewUint32(tableName, "field_rate_8")
	_subjectField.Rate9 = field.NewUint32(tableName, "field_rate_9")
	_subjectField.Rate10 = field.NewUint32(tableName, "field_rate_10")
	_subjectField.Airtime = field.NewUint8(tableName, "field_airtime")
	_subjectField.Rank = field.NewUint32(tableName, "field_rank")
	_subjectField.Year = field.NewInt32(tableName, "field_year")
	_subjectField.Mon = field.NewInt8(tableName, "field_mon")
	_subjectField.WeekDay = field.NewInt8(tableName, "field_week_day")
	_subjectField.Date = field.NewTime(tableName, "field_date")
	_subjectField.Redirect = field.NewField(tableName, "field_redirect")

	_subjectField.fillFieldMap()

	return _subjectField
}

type subjectField struct {
	subjectFieldDo subjectFieldDo

	ALL      field.Field
	Sid      field.Uint32
	Tid      field.Uint16
	Tags     field.Bytes
	Rate1    field.Uint32
	Rate2    field.Uint32
	Rate3    field.Uint32
	Rate4    field.Uint32
	Rate5    field.Uint32
	Rate6    field.Uint32
	Rate7    field.Uint32
	Rate8    field.Uint32
	Rate9    field.Uint32
	Rate10   field.Uint32
	Airtime  field.Uint8
	Rank     field.Uint32
	Year     field.Int32
	Mon      field.Int8
	WeekDay  field.Int8
	Date     field.Time
	Redirect field.Field

	fieldMap map[string]field.Expr
}

func (s subjectField) Table(newTableName string) *subjectField {
	s.subjectFieldDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectField) As(alias string) *subjectField {
	s.subjectFieldDo.DO = *(s.subjectFieldDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectField) updateTableName(table string) *subjectField {
	s.ALL = field.NewField(table, "*")
	s.Sid = field.NewUint32(table, "field_sid")
	s.Tid = field.NewUint16(table, "field_tid")
	s.Tags = field.NewBytes(table, "field_tags")
	s.Rate1 = field.NewUint32(table, "field_rate_1")
	s.Rate2 = field.NewUint32(table, "field_rate_2")
	s.Rate3 = field.NewUint32(table, "field_rate_3")
	s.Rate4 = field.NewUint32(table, "field_rate_4")
	s.Rate5 = field.NewUint32(table, "field_rate_5")
	s.Rate6 = field.NewUint32(table, "field_rate_6")
	s.Rate7 = field.NewUint32(table, "field_rate_7")
	s.Rate8 = field.NewUint32(table, "field_rate_8")
	s.Rate9 = field.NewUint32(table, "field_rate_9")
	s.Rate10 = field.NewUint32(table, "field_rate_10")
	s.Airtime = field.NewUint8(table, "field_airtime")
	s.Rank = field.NewUint32(table, "field_rank")
	s.Year = field.NewInt32(table, "field_year")
	s.Mon = field.NewInt8(table, "field_mon")
	s.WeekDay = field.NewInt8(table, "field_week_day")
	s.Date = field.NewTime(table, "field_date")
	s.Redirect = field.NewField(table, "field_redirect")

	s.fillFieldMap()

	return s
}

func (s *subjectField) WithContext(ctx context.Context) *subjectFieldDo {
	return s.subjectFieldDo.WithContext(ctx)
}

func (s subjectField) TableName() string { return s.subjectFieldDo.TableName() }

func (s subjectField) Alias() string { return s.subjectFieldDo.Alias() }

func (s *subjectField) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectField) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["field_sid"] = s.Sid
	s.fieldMap["field_tid"] = s.Tid
	s.fieldMap["field_tags"] = s.Tags
	s.fieldMap["field_rate_1"] = s.Rate1
	s.fieldMap["field_rate_2"] = s.Rate2
	s.fieldMap["field_rate_3"] = s.Rate3
	s.fieldMap["field_rate_4"] = s.Rate4
	s.fieldMap["field_rate_5"] = s.Rate5
	s.fieldMap["field_rate_6"] = s.Rate6
	s.fieldMap["field_rate_7"] = s.Rate7
	s.fieldMap["field_rate_8"] = s.Rate8
	s.fieldMap["field_rate_9"] = s.Rate9
	s.fieldMap["field_rate_10"] = s.Rate10
	s.fieldMap["field_airtime"] = s.Airtime
	s.fieldMap["field_rank"] = s.Rank
	s.fieldMap["field_year"] = s.Year
	s.fieldMap["field_mon"] = s.Mon
	s.fieldMap["field_week_day"] = s.WeekDay
	s.fieldMap["field_date"] = s.Date
	s.fieldMap["field_redirect"] = s.Redirect
}

func (s subjectField) clone(db *gorm.DB) subjectField {
	s.subjectFieldDo.ReplaceDB(db)
	return s
}

type subjectFieldDo struct{ gen.DO }

func (s subjectFieldDo) Debug() *subjectFieldDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectFieldDo) WithContext(ctx context.Context) *subjectFieldDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectFieldDo) ReadDB() *subjectFieldDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectFieldDo) WriteDB() *subjectFieldDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectFieldDo) Clauses(conds ...clause.Expression) *subjectFieldDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectFieldDo) Returning(value interface{}, columns ...string) *subjectFieldDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectFieldDo) Not(conds ...gen.Condition) *subjectFieldDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectFieldDo) Or(conds ...gen.Condition) *subjectFieldDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectFieldDo) Select(conds ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectFieldDo) Where(conds ...gen.Condition) *subjectFieldDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectFieldDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectFieldDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectFieldDo) Order(conds ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectFieldDo) Distinct(cols ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectFieldDo) Omit(cols ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectFieldDo) Join(table schema.Tabler, on ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectFieldDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectFieldDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectFieldDo) Group(cols ...field.Expr) *subjectFieldDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectFieldDo) Having(conds ...gen.Condition) *subjectFieldDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectFieldDo) Limit(limit int) *subjectFieldDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectFieldDo) Offset(offset int) *subjectFieldDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectFieldDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectFieldDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectFieldDo) Unscoped() *subjectFieldDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectFieldDo) Create(values ...*dao.SubjectField) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectFieldDo) CreateInBatches(values []*dao.SubjectField, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectFieldDo) Save(values ...*dao.SubjectField) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectFieldDo) First() (*dao.SubjectField, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectField), nil
	}
}

func (s subjectFieldDo) Take() (*dao.SubjectField, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectField), nil
	}
}

func (s subjectFieldDo) Last() (*dao.SubjectField, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectField), nil
	}
}

func (s subjectFieldDo) Find() ([]*dao.SubjectField, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectField), err
}

func (s subjectFieldDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectField, err error) {
	buf := make([]*dao.SubjectField, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectFieldDo) FindInBatches(result *[]*dao.SubjectField, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectFieldDo) Attrs(attrs ...field.AssignExpr) *subjectFieldDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectFieldDo) Assign(attrs ...field.AssignExpr) *subjectFieldDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectFieldDo) Joins(fields ...field.RelationField) *subjectFieldDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectFieldDo) Preload(fields ...field.RelationField) *subjectFieldDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectFieldDo) FirstOrInit() (*dao.SubjectField, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectField), nil
	}
}

func (s subjectFieldDo) FirstOrCreate() (*dao.SubjectField, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectField), nil
	}
}

func (s subjectFieldDo) FindByPage(offset int, limit int) (result []*dao.SubjectField, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectFieldDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectFieldDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s *subjectFieldDo) withDO(do gen.Dao) *subjectFieldDo {
	s.DO = *do.(*gen.DO)
	return s
}
