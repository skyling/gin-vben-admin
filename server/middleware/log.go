package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var bodyStr string
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				logrus.Error("read body from request error:", err)
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			bodyStr = "[文件]"
		} else {
			if len(body) > bufferSize {
				bodyStr = "[超出记录长度]"
			} else {
				bodyStr = string(body)
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		var respBody string
		latency := time.Since(now)
		respBody = writer.body.String()

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(respBody) > bufferSize {
				// 截断
				respBody = "超出记录长度"
			}
		}
		statusCode := c.Writer.Status()
		entry := logger.WithFields(logrus.Fields{
			"statusCode":    statusCode,
			"error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"latency":       latency, // time to process
			"client_ip":     c.ClientIP(),
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"user_agent":    c.Request.UserAgent(),
			"body":          bodyStr,
			//"resp":          respBody,
			"user_id": userId,
		})
		if c.Writer.Status() >= http.StatusInternalServerError {
			entry.Error()
		} else if statusCode >= http.StatusBadRequest {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
