package repo

import (
	"gin-vben-admin/dao/query"
)

var SettingSrv = new(settingSrv)

type settingSrv struct {
}

func (s *settingSrv) GetValue(name string) string {
	qs := query.Setting
	set, _ := qs.Where(qs.Name.Eq(name)).First()
	if set == nil {
		return ""
	}
	return set.Value
}
