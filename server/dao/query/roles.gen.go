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

func newRole(db *gorm.DB, opts ...gen.DOOption) role {
	_role := role{}

	_role.roleDo.UseDB(db, opts...)
	_role.roleDo.UseModel(&models.Role{})

	tableName := _role.roleDo.TableName()
	_role.ALL = field.NewAsterisk(tableName)
	_role.ID = field.NewInt64(tableName, "id")
	_role.CreatedAt = field.NewTime(tableName, "created_at")
	_role.UpdatedAt = field.NewTime(tableName, "updated_at")
	_role.DeletedAt = field.NewField(tableName, "deleted_at")
	_role.TenantID = field.NewInt64(tableName, "tenant_id")
	_role.Name = field.NewString(tableName, "name")
	_role.Code = field.NewString(tableName, "code")
	_role.Sort = field.NewInt64(tableName, "sort")
	_role.Status = field.NewInt64(tableName, "status")
	_role.Remark = field.NewString(tableName, "remark")
	_role.Permissions = roleManyToManyPermissions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Permissions", "models.Permission"),
		Permissions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Permissions.Permissions", "models.Permission"),
		},
		Roles: struct {
			field.RelationField
			Permissions struct {
				field.RelationField
			}
			Users struct {
				field.RelationField
				Dept struct {
					field.RelationField
					Depts struct {
						field.RelationField
					}
				}
				Roles struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Permissions.Roles", "models.Role"),
			Permissions: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Permissions.Roles.Permissions", "models.Permission"),
			},
			Users: struct {
				field.RelationField
				Dept struct {
					field.RelationField
					Depts struct {
						field.RelationField
					}
				}
				Roles struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Permissions.Roles.Users", "models.User"),
				Dept: struct {
					field.RelationField
					Depts struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Permissions.Roles.Users.Dept", "models.Dept"),
					Depts: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Permissions.Roles.Users.Dept.Depts", "models.Dept"),
					},
				},
				Roles: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Permissions.Roles.Users.Roles", "models.Role"),
				},
			},
		},
	}

	_role.Users = roleManyToManyUsers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Users", "models.User"),
	}

	_role.fillFieldMap()

	return _role
}

type role struct {
	roleDo

	ALL         field.Asterisk
	ID          field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	TenantID    field.Int64
	Name        field.String
	Code        field.String
	Sort        field.Int64
	Status      field.Int64
	Remark      field.String
	Permissions roleManyToManyPermissions

	Users roleManyToManyUsers

	fieldMap map[string]field.Expr
}

func (r role) Table(newTableName string) *role {
	r.roleDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r role) As(alias string) *role {
	r.roleDo.DO = *(r.roleDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *role) updateTableName(table string) *role {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.TenantID = field.NewInt64(table, "tenant_id")
	r.Name = field.NewString(table, "name")
	r.Code = field.NewString(table, "code")
	r.Sort = field.NewInt64(table, "sort")
	r.Status = field.NewInt64(table, "status")
	r.Remark = field.NewString(table, "remark")

	r.fillFieldMap()

	return r
}

func (r *role) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *role) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 12)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["tenant_id"] = r.TenantID
	r.fieldMap["name"] = r.Name
	r.fieldMap["code"] = r.Code
	r.fieldMap["sort"] = r.Sort
	r.fieldMap["status"] = r.Status
	r.fieldMap["remark"] = r.Remark

}

func (r role) clone(db *gorm.DB) role {
	r.roleDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r role) replaceDB(db *gorm.DB) role {
	r.roleDo.ReplaceDB(db)
	return r
}

type roleManyToManyPermissions struct {
	db *gorm.DB

	field.RelationField

	Permissions struct {
		field.RelationField
	}
	Roles struct {
		field.RelationField
		Permissions struct {
			field.RelationField
		}
		Users struct {
			field.RelationField
			Dept struct {
				field.RelationField
				Depts struct {
					field.RelationField
				}
			}
			Roles struct {
				field.RelationField
			}
		}
	}
}

func (a roleManyToManyPermissions) Where(conds ...field.Expr) *roleManyToManyPermissions {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a roleManyToManyPermissions) WithContext(ctx context.Context) *roleManyToManyPermissions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a roleManyToManyPermissions) Session(session *gorm.Session) *roleManyToManyPermissions {
	a.db = a.db.Session(session)
	return &a
}

func (a roleManyToManyPermissions) Model(m *models.Role) *roleManyToManyPermissionsTx {
	return &roleManyToManyPermissionsTx{a.db.Model(m).Association(a.Name())}
}

type roleManyToManyPermissionsTx struct{ tx *gorm.Association }

func (a roleManyToManyPermissionsTx) Find() (result []*models.Permission, err error) {
	return result, a.tx.Find(&result)
}

func (a roleManyToManyPermissionsTx) Append(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a roleManyToManyPermissionsTx) Replace(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a roleManyToManyPermissionsTx) Delete(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a roleManyToManyPermissionsTx) Clear() error {
	return a.tx.Clear()
}

func (a roleManyToManyPermissionsTx) Count() int64 {
	return a.tx.Count()
}

type roleManyToManyUsers struct {
	db *gorm.DB

	field.RelationField
}

func (a roleManyToManyUsers) Where(conds ...field.Expr) *roleManyToManyUsers {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a roleManyToManyUsers) WithContext(ctx context.Context) *roleManyToManyUsers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a roleManyToManyUsers) Session(session *gorm.Session) *roleManyToManyUsers {
	a.db = a.db.Session(session)
	return &a
}

func (a roleManyToManyUsers) Model(m *models.Role) *roleManyToManyUsersTx {
	return &roleManyToManyUsersTx{a.db.Model(m).Association(a.Name())}
}

type roleManyToManyUsersTx struct{ tx *gorm.Association }

func (a roleManyToManyUsersTx) Find() (result []*models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a roleManyToManyUsersTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a roleManyToManyUsersTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a roleManyToManyUsersTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a roleManyToManyUsersTx) Clear() error {
	return a.tx.Clear()
}

func (a roleManyToManyUsersTx) Count() int64 {
	return a.tx.Count()
}

type roleDo struct{ gen.DO }

type IRoleDo interface {
	gen.SubQuery
	Debug() IRoleDo
	WithContext(ctx context.Context) IRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoleDo
	WriteDB() IRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoleDo
	Not(conds ...gen.Condition) IRoleDo
	Or(conds ...gen.Condition) IRoleDo
	Select(conds ...field.Expr) IRoleDo
	Where(conds ...gen.Condition) IRoleDo
	Order(conds ...field.Expr) IRoleDo
	Distinct(cols ...field.Expr) IRoleDo
	Omit(cols ...field.Expr) IRoleDo
	Join(table schema.Tabler, on ...field.Expr) IRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoleDo
	Group(cols ...field.Expr) IRoleDo
	Having(conds ...gen.Condition) IRoleDo
	Limit(limit int) IRoleDo
	Offset(offset int) IRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleDo
	Unscoped() IRoleDo
	Create(values ...*models.Role) error
	CreateInBatches(values []*models.Role, batchSize int) error
	Save(values ...*models.Role) error
	First() (*models.Role, error)
	Take() (*models.Role, error)
	Last() (*models.Role, error)
	Find() ([]*models.Role, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Role, err error)
	FindInBatches(result *[]*models.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Role) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoleDo
	Assign(attrs ...field.AssignExpr) IRoleDo
	Joins(fields ...field.RelationField) IRoleDo
	Preload(fields ...field.RelationField) IRoleDo
	FirstOrInit() (*models.Role, error)
	FirstOrCreate() (*models.Role, error)
	FindByPage(offset int, limit int) (result []*models.Role, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int64) (result *models.Role, err error)
	GetByIds(ids []int64) (result []*models.Role, err error)
	GetByIDAndTID(tid int64, id int64) (result *models.Role, err error)
	GetByIdsAndTID(tid int64, ids []int64) (result []*models.Role, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (r roleDo) GetByID(id int64) (result *models.Role, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM roles WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id in (@ids)
func (r roleDo) GetByIds(ids []int64) (result []*models.Role, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM roles WHERE id in (?) ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id=@id and tenant_id=@tid
func (r roleDo) GetByIDAndTID(tid int64, id int64) (result *models.Role, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, tid)
	generateSQL.WriteString("SELECT * FROM roles WHERE id=? and tenant_id=? ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id in (@ids) and tenant_id=@tid
func (r roleDo) GetByIdsAndTID(tid int64, ids []int64) (result []*models.Role, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	params = append(params, tid)
	generateSQL.WriteString("SELECT * FROM roles WHERE id in (?) and tenant_id=? ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (r roleDo) Debug() IRoleDo {
	return r.withDO(r.DO.Debug())
}

func (r roleDo) WithContext(ctx context.Context) IRoleDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleDo) ReadDB() IRoleDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleDo) WriteDB() IRoleDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleDo) Session(config *gorm.Session) IRoleDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleDo) Clauses(conds ...clause.Expression) IRoleDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleDo) Returning(value interface{}, columns ...string) IRoleDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleDo) Not(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleDo) Or(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleDo) Select(conds ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleDo) Where(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleDo) Order(conds ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleDo) Distinct(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleDo) Omit(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleDo) Join(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleDo) Group(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleDo) Having(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleDo) Limit(limit int) IRoleDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleDo) Offset(offset int) IRoleDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleDo) Unscoped() IRoleDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleDo) Create(values ...*models.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleDo) CreateInBatches(values []*models.Role, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleDo) Save(values ...*models.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleDo) First() (*models.Role, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Take() (*models.Role, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Last() (*models.Role, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Find() ([]*models.Role, error) {
	result, err := r.DO.Find()
	return result.([]*models.Role), err
}

func (r roleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Role, err error) {
	buf := make([]*models.Role, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleDo) FindInBatches(result *[]*models.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleDo) Attrs(attrs ...field.AssignExpr) IRoleDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleDo) Assign(attrs ...field.AssignExpr) IRoleDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleDo) Joins(fields ...field.RelationField) IRoleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleDo) Preload(fields ...field.RelationField) IRoleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleDo) FirstOrInit() (*models.Role, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) FirstOrCreate() (*models.Role, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) FindByPage(offset int, limit int) (result []*models.Role, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleDo) Delete(models ...*models.Role) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleDo) withDO(do gen.Dao) *roleDo {
	r.DO = *do.(*gen.DO)
	return r
}
