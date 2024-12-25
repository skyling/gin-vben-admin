package utils

import "github.com/bwmarrin/snowflake"

func ParseIDstoInt64(ids []snowflake.ID) []int64 {
	ret := []int64{}
	for _, id := range ids {
		ret = append(ret, id.Int64())
	}
	return ret
}
