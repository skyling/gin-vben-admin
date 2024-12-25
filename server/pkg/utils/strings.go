package utils

import (
	"github.com/duke-git/lancet/v2/slice"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

func SplitBySize(str string, length int) []string {
	all := strings.Split(str, "")
	allChuck := slice.Chunk(all, length)
	ret := []string{}
	for _, v := range allChuck {
		ret = append(ret, strings.Join(v, ""))
	}
	return ret
}

func JSON(v interface{}) string {
	str, _ := jsoniter.ConfigDefault.MarshalToString(v)
	return str
}

func UnJSON(str string, v interface{}) error {
	return jsoniter.ConfigDefault.UnmarshalFromString(str, v)
}
