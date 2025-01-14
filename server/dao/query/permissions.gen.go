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

func newPermission(db *gorm.DB, opts ...gen.DOOption) permission {
	_permission := permission{}

	_permission.permissionDo.UseDB(db, opts...)
	_permission.permissionDo.UseModel(&models.Permission{})

	tableName := _permission.permissionDo.TableName()
	_permission.ALL = field.NewAsterisk(tableName)
	_permission.ID = field.NewInt64(tableName, "id")
	_permission.CreatedAt = field.NewTime(tableName, "created_at")
	_permission.UpdatedAt = field.NewTime(tableName, "updated_at")
	_permission.DeletedAt = field.NewField(tableName, "deleted_at")
	_permission.Type = field.NewString(tableName, "type")
	_permission.ParentID = field.NewInt64(tableName, "parent_id")
	_permission.Sort = field.NewInt64(tableName, "sort")
	_permission.Name = field.NewString(tableName, "name")
	_permission.Code = field.NewString(tableName, "code")
	_permission.Description = field.NewString(tableName, "description")
	_permission.Permissions = permissionHasManyPermissions{
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

	_permission.Roles = permissionManyToManyRoles{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Roles", "models.Role"),
	}

	_permission.fillFieldMap()

	return _permission
}

type permission struct {
	permissionDo

	ALL         field.Asterisk
	ID          field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Type        field.String
	ParentID    field.Int64
	Sort        field.Int64
	Name        field.String
	Code        field.String
	Description field.String
	Permissions permissionHasManyPermissions

	Roles permissionManyToManyRoles

	fieldMap map[string]field.Expr
}

func (p permission) Table(newTableName string) *permission {
	p.permissionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p permission) As(alias string) *permission {
	p.permissionDo.DO = *(p.permissionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *permission) updateTableName(table string) *permission {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Type = field.NewString(table, "type")
	p.ParentID = field.NewInt64(table, "parent_id")
	p.Sort = field.NewInt64(table, "sort")
	p.Name = field.NewString(table, "name")
	p.Code = field.NewString(table, "code")
	p.Description = field.NewString(table, "description")

	p.fillFieldMap()

	return p
}

func (p *permission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *permission) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["type"] = p.Type
	p.fieldMap["parent_id"] = p.ParentID
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["name"] = p.Name
	p.fieldMap["code"] = p.Code
	p.fieldMap["description"] = p.Description

}

func (p permission) clone(db *gorm.DB) permission {
	p.permissionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p permission) replaceDB(db *gorm.DB) permission {
	p.permissionDo.ReplaceDB(db)
	return p
}

type permissionHasManyPermissions struct {
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

func (a permissionHasManyPermissions) Where(conds ...field.Expr) *permissionHasManyPermissions {
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

func (a permissionHasManyPermissions) WithContext(ctx context.Context) *permissionHasManyPermissions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a permissionHasManyPermissions) Session(session *gorm.Session) *permissionHasManyPermissions {
	a.db = a.db.Session(session)
	return &a
}

func (a permissionHasManyPermissions) Model(m *models.Permission) *permissionHasManyPermissionsTx {
	return &permissionHasManyPermissionsTx{a.db.Model(m).Association(a.Name())}
}

type permissionHasManyPermissionsTx struct{ tx *gorm.Association }

func (a permissionHasManyPermissionsTx) Find() (result []*models.Permission, err error) {
	return result, a.tx.Find(&result)
}

func (a permissionHasManyPermissionsTx) Append(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a permissionHasManyPermissionsTx) Replace(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a permissionHasManyPermissionsTx) Delete(values ...*models.Permission) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a permissionHasManyPermissionsTx) Clear() error {
	return a.tx.Clear()
}

func (a permissionHasManyPermissionsTx) Count() int64 {
	return a.tx.Count()
}

type permissionManyToManyRoles struct {
	db *gorm.DB

	field.RelationField
}

func (a permissionManyToManyRoles) Where(conds ...field.Expr) *permissionManyToManyRoles {
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

func (a permissionManyToManyRoles) WithContext(ctx context.Context) *permissionManyToManyRoles {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a permissionManyToManyRoles) Session(session *gorm.Session) *permissionManyToManyRoles {
	a.db = a.db.Session(session)
	return &a
}

func (a permissionManyToManyRoles) Model(m *models.Permission) *permissionManyToManyRolesTx {
	return &permissionManyToManyRolesTx{a.db.Model(m).Association(a.Name())}
}

type permissionManyToManyRolesTx struct{ tx *gorm.Association }

func (a permissionManyToManyRolesTx) Find() (result []*models.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a permissionManyToManyRolesTx) Append(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a permissionManyToManyRolesTx) Replace(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a permissionManyToManyRolesTx) Delete(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a permissionManyToManyRolesTx) Clear() error {
	return a.tx.Clear()
}

func (a permissionManyToManyRolesTx) Count() int64 {
	return a.tx.Count()
}

type permissionDo struct{ gen.DO }

type IPermissionDo interface {
	gen.SubQuery
	Debug() IPermissionDo
	WithContext(ctx context.Context) IPermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPermissionDo
	WriteDB() IPermissionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPermissionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPermissionDo
	Not(conds ...gen.Condition) IPermissionDo
	Or(conds ...gen.Condition) IPermissionDo
	Select(conds ...field.Expr) IPermissionDo
	Where(conds ...gen.Condition) IPermissionDo
	Order(conds ...field.Expr) IPermissionDo
	Distinct(cols ...field.Expr) IPermissionDo
	Omit(cols ...field.Expr) IPermissionDo
	Join(table schema.Tabler, on ...field.Expr) IPermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPermissionDo
	Group(cols ...field.Expr) IPermissionDo
	Having(conds ...gen.Condition) IPermissionDo
	Limit(limit int) IPermissionDo
	Offset(offset int) IPermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPermissionDo
	Unscoped() IPermissionDo
	Create(values ...*models.Permission) error
	CreateInBatches(values []*models.Permission, batchSize int) error
	Save(values ...*models.Permission) error
	First() (*models.Permission, error)
	Take() (*models.Permission, error)
	Last() (*models.Permission, error)
	Find() ([]*models.Permission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Permission, err error)
	FindInBatches(result *[]*models.Permission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Permission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPermissionDo
	Assign(attrs ...field.AssignExpr) IPermissionDo
	Joins(fields ...field.RelationField) IPermissionDo
	Preload(fields ...field.RelationField) IPermissionDo
	FirstOrInit() (*models.Permission, error)
	FirstOrCreate() (*models.Permission, error)
	FindByPage(offset int, limit int) (result []*models.Permission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int64) (result *models.Permission, err error)
	GetByIds(ids []int64) (result []*models.Permission, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (p permissionDo) GetByID(id int64) (result *models.Permission, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM permissions WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE id in (@ids)
func (p permissionDo) GetByIds(ids []int64) (result []*models.Permission, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM permissions WHERE id in (?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p permissionDo) Debug() IPermissionDo {
	return p.withDO(p.DO.Debug())
}

func (p permissionDo) WithContext(ctx context.Context) IPermissionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p permissionDo) ReadDB() IPermissionDo {
	return p.Clauses(dbresolver.Read)
}

func (p permissionDo) WriteDB() IPermissionDo {
	return p.Clauses(dbresolver.Write)
}

func (p permissionDo) Session(config *gorm.Session) IPermissionDo {
	return p.withDO(p.DO.Session(config))
}

func (p permissionDo) Clauses(conds ...clause.Expression) IPermissionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p permissionDo) Returning(value interface{}, columns ...string) IPermissionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p permissionDo) Not(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p permissionDo) Or(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p permissionDo) Select(conds ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p permissionDo) Where(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p permissionDo) Order(conds ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p permissionDo) Distinct(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p permissionDo) Omit(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p permissionDo) Join(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p permissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p permissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p permissionDo) Group(cols ...field.Expr) IPermissionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p permissionDo) Having(conds ...gen.Condition) IPermissionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p permissionDo) Limit(limit int) IPermissionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p permissionDo) Offset(offset int) IPermissionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p permissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPermissionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p permissionDo) Unscoped() IPermissionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p permissionDo) Create(values ...*models.Permission) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p permissionDo) CreateInBatches(values []*models.Permission, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p permissionDo) Save(values ...*models.Permission) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p permissionDo) First() (*models.Permission, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Permission), nil
	}
}

func (p permissionDo) Take() (*models.Permission, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Permission), nil
	}
}

func (p permissionDo) Last() (*models.Permission, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Permission), nil
	}
}

func (p permissionDo) Find() ([]*models.Permission, error) {
	result, err := p.DO.Find()
	return result.([]*models.Permission), err
}

func (p permissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Permission, err error) {
	buf := make([]*models.Permission, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p permissionDo) FindInBatches(result *[]*models.Permission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p permissionDo) Attrs(attrs ...field.AssignExpr) IPermissionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p permissionDo) Assign(attrs ...field.AssignExpr) IPermissionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p permissionDo) Joins(fields ...field.RelationField) IPermissionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p permissionDo) Preload(fields ...field.RelationField) IPermissionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p permissionDo) FirstOrInit() (*models.Permission, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Permission), nil
	}
}

func (p permissionDo) FirstOrCreate() (*models.Permission, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Permission), nil
	}
}

func (p permissionDo) FindByPage(offset int, limit int) (result []*models.Permission, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p permissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p permissionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p permissionDo) Delete(models ...*models.Permission) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *permissionDo) withDO(do gen.Dao) *permissionDo {
	p.DO = *do.(*gen.DO)
	return p
}
