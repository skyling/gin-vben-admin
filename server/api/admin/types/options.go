package types

import (
	"gin-vben-admin/pkg/e"
)

type OptionsItem struct {
	Label string      `json:"label"` // 英文名称
	Value interface{} `json:"value"` // 值
}

type OptionsData struct {
	Items []*OptionsItem `json:"items"`
}

func RespOptionsData(data *OptionsData) e.Response {
	return e.RespOK.WithData(data)
}
