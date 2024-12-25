package conf

import (
	"gin-vben-admin/pkg/utils"
	"path/filepath"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	if s.Host == "" {
		s.Host = "."
	}
	return filepath.Join(utils.RelativePath(s.Host), s.Dbname+".db")
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
