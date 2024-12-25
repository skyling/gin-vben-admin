package locales

import "golang.org/x/text/language"

func init() {
	initZh(language.Make("zh-CN"))
	initZh(language.Make("zh-Hans"))
}
