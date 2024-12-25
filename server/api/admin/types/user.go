package types

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"time"
)

type LoginLogItem struct {
	ID        snowflake.ID `json:"id" swaggertype:"string"`
	IP        string       `json:"ip"`        // 登录的IP
	Source    string       `json:"source"`    // 来源
	Time      time.Time    `json:"time"`      // 时间
	UserAgent string       `json:"userAgent"` // 请求接口头部信息
}

type LoginLogData struct {
	Items []*LoginLogItem `json:"items"` // 数据列表
	Total int64           `json:"total"` // 总数据条数
}

func RespLoginLogs(logs []*models.LoginLog, total int64) e.Response {
	resp := LoginLogData{
		Total: total,
	}
	for _, log := range logs {
		resp.Items = append(resp.Items, &LoginLogItem{
			ID:        log.ID,
			IP:        log.IP,
			Source:    log.Source,
			Time:      log.CreatedAt,
			UserAgent: log.UserAgent,
		})
	}

	return e.RespOK.WithData(resp)
}
