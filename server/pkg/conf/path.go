package conf

import (
	"gin-vben-admin/pkg/utils"
	"path"
	"strings"
	"time"
)

const (
	PathUpload = "upload"
	PathTemp   = "tmp"
)

func InitLocalPath() error {
	base := C.System.ResourcePath
	utils.CreateLocalPath(path.Join(base, PathUpload))
	utils.CreateLocalPath(path.Join(base, PathTemp))
	return nil
}

func GetPath(dir string, filekey string) (string, string) {
	filepathStr := path.Join(C.System.ResourcePath, dir, filekey)
	utils.CreateLocalPath(path.Dir(filepathStr))
	return filepathStr, "/" + strings.TrimPrefix(strings.TrimPrefix(filepathStr, C.System.ResourcePath), "/")
}

func GetFilePath(filekey string) string {
	return path.Join(C.System.ResourcePath, filekey)
}

func GetDailyPath(dir string, filekey string) (string, string) {
	filepathStr := path.Join(C.System.ResourcePath, dir, time.Now().Format("20060102"), filekey)
	utils.CreateLocalPath(path.Dir(filepathStr))
	return filepathStr, "/" + strings.TrimPrefix(strings.TrimPrefix(filepathStr, C.System.ResourcePath), "/")
}

func BuildUrl(filepath string) string {
	return C.System.BaseURI + filepath
}
