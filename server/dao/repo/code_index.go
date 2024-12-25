package repo

import (
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/pkg/lock"
	"gorm.io/gorm/clause"
	"time"
)

// GenCodeByDay 生成CODE 序号为每天
func GenCodeByDay(prefix string) (string, error) {
	date := time.Now().Format("20060102")
	return GenCode(prefix, date, "-")
}

func GenCodeByMonth(prefix string) (string, error) {
	date := time.Now().Format("200601")
	return GenCode(prefix, date, "-")
}

func GenCodeByYear(prefix string) (string, error) {
	date := time.Now().Format("2006")
	return GenCode(prefix, date, "")
}

func GenCode(prefix string, mid string, sep string) (string, error) {
	ciq := query.CodeIndex
	key := fmt.Sprintf("lock-%s-%s", prefix, mid)
	lo, ctx, err := lock.LockOpt(key, 3*time.Second, 10*time.Millisecond, 300, false)
	if err != nil {
		return "", err
	}
	defer lo.Release(ctx)
	codeStr := ""
	err = query.Q.Transaction(func(tx *query.Query) error {
		idx := time.Now().Unix() % 86400
		code, _ := tx.CodeIndex.Attrs(ciq.Index.Value(idx)).Where(ciq.Type.Eq(prefix), ciq.Date.Eq(mid)).Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).FirstOrCreate()
		if code == nil {
			code = &models.CodeIndex{
				Index: idx,
				Type:  prefix,
				Date:  mid,
			}
			err := tx.CodeIndex.Save(code)
			if err != nil {
				return err
			}
		} else {
			code.Index = code.Index + 1
			_, err := tx.CodeIndex.Where(ciq.ID.Eq(code.ID.Int64())).UpdateSimple(ciq.Index.Value(code.Index))
			if err != nil {
				return err
			}
		}
		codeStr = fmt.Sprintf("%s%s%s%s%05d", prefix, sep, mid, sep, code.Index)
		return nil
	})
	return codeStr, err
}
