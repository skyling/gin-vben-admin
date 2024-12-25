package utils

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path/filepath"
)

// RelativePath 获取相对可执行文件的路径
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}

func CreateLocalPath(pathStr string) error {
	if fileutil.IsExist(pathStr) {
		return nil
	}
	return fileutil.CreateDir(pathStr)
}
