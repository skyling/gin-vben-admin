package main

import (
	"gin-vben-admin/dao/models"
	"gorm.io/gen"
)

// Querier Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int64) (*gen.T, error)
	// SELECT * FROM @@table WHERE id in (@ids)
	GetByIds(ids []int64) ([]*gen.T, error)
}

type TenantQuerier interface {
	// SELECT * FROM @@table WHERE id=@id and tenant_id=@tid
	GetByIDAndTID(tid, id int64) (*gen.T, error)
	// SELECT * FROM @@table WHERE id in (@ids) and tenant_id=@tid
	GetByIdsAndTID(tid int64, ids []int64) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "dao/query",
		Mode:          gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.ApplyBasic(models.ModelsList...)
	g.ApplyInterface(func(Querier) {}, models.ModelsList...)
	g.ApplyInterface(func(TenantQuerier) {}, models.ModelsTenantList...)
	g.Execute()
}
