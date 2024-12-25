// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gin-vben-admin/dao/models"
)

func newSetting(db *gorm.DB, opts ...gen.DOOption) setting {
	_setting := setting{}

	_setting.settingDo.UseDB(db, opts...)
	_setting.settingDo.UseModel(&models.Setting{})

	tableName := _setting.settingDo.TableName()
	_setting.ALL = field.NewAsterisk(tableName)
	_setting.ID = field.NewInt64(tableName, "id")
	_setting.CreatedAt = field.NewTime(tableName, "created_at")
	_setting.UpdatedAt = field.NewTime(tableName, "updated_at")
	_setting.DeletedAt = field.NewField(tableName, "deleted_at")
	_setting.TenantID = field.NewInt64(tableName, "tenant_id")
	_setting.Type = field.NewString(tableName, "type")
	_setting.Name = field.NewString(tableName, "name")
	_setting.Value = field.NewString(tableName, "value")

	_setting.fillFieldMap()

	return _setting
}

type setting struct {
	settingDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	TenantID  field.Int64
	Type      field.String
	Name      field.String
	Value     field.String

	fieldMap map[string]field.Expr
}

func (s setting) Table(newTableName string) *setting {
	s.settingDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s setting) As(alias string) *setting {
	s.settingDo.DO = *(s.settingDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *setting) updateTableName(table string) *setting {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.TenantID = field.NewInt64(table, "tenant_id")
	s.Type = field.NewString(table, "type")
	s.Name = field.NewString(table, "name")
	s.Value = field.NewString(table, "value")

	s.fillFieldMap()

	return s
}

func (s *setting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *setting) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["tenant_id"] = s.TenantID
	s.fieldMap["type"] = s.Type
	s.fieldMap["name"] = s.Name
	s.fieldMap["value"] = s.Value
}

func (s setting) clone(db *gorm.DB) setting {
	s.settingDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s setting) replaceDB(db *gorm.DB) setting {
	s.settingDo.ReplaceDB(db)
	return s
}

type settingDo struct{ gen.DO }

type ISettingDo interface {
	gen.SubQuery
	Debug() ISettingDo
	WithContext(ctx context.Context) ISettingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISettingDo
	WriteDB() ISettingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISettingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISettingDo
	Not(conds ...gen.Condition) ISettingDo
	Or(conds ...gen.Condition) ISettingDo
	Select(conds ...field.Expr) ISettingDo
	Where(conds ...gen.Condition) ISettingDo
	Order(conds ...field.Expr) ISettingDo
	Distinct(cols ...field.Expr) ISettingDo
	Omit(cols ...field.Expr) ISettingDo
	Join(table schema.Tabler, on ...field.Expr) ISettingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISettingDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISettingDo
	Group(cols ...field.Expr) ISettingDo
	Having(conds ...gen.Condition) ISettingDo
	Limit(limit int) ISettingDo
	Offset(offset int) ISettingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISettingDo
	Unscoped() ISettingDo
	Create(values ...*models.Setting) error
	CreateInBatches(values []*models.Setting, batchSize int) error
	Save(values ...*models.Setting) error
	First() (*models.Setting, error)
	Take() (*models.Setting, error)
	Last() (*models.Setting, error)
	Find() ([]*models.Setting, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Setting, err error)
	FindInBatches(result *[]*models.Setting, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Setting) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISettingDo
	Assign(attrs ...field.AssignExpr) ISettingDo
	Joins(fields ...field.RelationField) ISettingDo
	Preload(fields ...field.RelationField) ISettingDo
	FirstOrInit() (*models.Setting, error)
	FirstOrCreate() (*models.Setting, error)
	FindByPage(offset int, limit int) (result []*models.Setting, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISettingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int64) (result *models.Setting, err error)
	GetByIds(ids []int64) (result []*models.Setting, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (s settingDo) GetByID(id int64) (result *models.Setting, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM settings WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id in (@ids)
func (s settingDo) GetByIds(ids []int64) (result []*models.Setting, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM settings WHERE id in (?) ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s settingDo) Debug() ISettingDo {
	return s.withDO(s.DO.Debug())
}

func (s settingDo) WithContext(ctx context.Context) ISettingDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s settingDo) ReadDB() ISettingDo {
	return s.Clauses(dbresolver.Read)
}

func (s settingDo) WriteDB() ISettingDo {
	return s.Clauses(dbresolver.Write)
}

func (s settingDo) Session(config *gorm.Session) ISettingDo {
	return s.withDO(s.DO.Session(config))
}

func (s settingDo) Clauses(conds ...clause.Expression) ISettingDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s settingDo) Returning(value interface{}, columns ...string) ISettingDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s settingDo) Not(conds ...gen.Condition) ISettingDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s settingDo) Or(conds ...gen.Condition) ISettingDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s settingDo) Select(conds ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s settingDo) Where(conds ...gen.Condition) ISettingDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s settingDo) Order(conds ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s settingDo) Distinct(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s settingDo) Omit(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s settingDo) Join(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s settingDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s settingDo) RightJoin(table schema.Tabler, on ...field.Expr) ISettingDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s settingDo) Group(cols ...field.Expr) ISettingDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s settingDo) Having(conds ...gen.Condition) ISettingDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s settingDo) Limit(limit int) ISettingDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s settingDo) Offset(offset int) ISettingDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s settingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISettingDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s settingDo) Unscoped() ISettingDo {
	return s.withDO(s.DO.Unscoped())
}

func (s settingDo) Create(values ...*models.Setting) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s settingDo) CreateInBatches(values []*models.Setting, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s settingDo) Save(values ...*models.Setting) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s settingDo) First() (*models.Setting, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Setting), nil
	}
}

func (s settingDo) Take() (*models.Setting, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Setting), nil
	}
}

func (s settingDo) Last() (*models.Setting, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Setting), nil
	}
}

func (s settingDo) Find() ([]*models.Setting, error) {
	result, err := s.DO.Find()
	return result.([]*models.Setting), err
}

func (s settingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Setting, err error) {
	buf := make([]*models.Setting, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s settingDo) FindInBatches(result *[]*models.Setting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s settingDo) Attrs(attrs ...field.AssignExpr) ISettingDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s settingDo) Assign(attrs ...field.AssignExpr) ISettingDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s settingDo) Joins(fields ...field.RelationField) ISettingDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s settingDo) Preload(fields ...field.RelationField) ISettingDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s settingDo) FirstOrInit() (*models.Setting, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Setting), nil
	}
}

func (s settingDo) FirstOrCreate() (*models.Setting, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Setting), nil
	}
}

func (s settingDo) FindByPage(offset int, limit int) (result []*models.Setting, count int64, err error) {
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

func (s settingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s settingDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s settingDo) Delete(models ...*models.Setting) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *settingDo) withDO(do gen.Dao) *settingDo {
	s.DO = *do.(*gen.DO)
	return s
}
