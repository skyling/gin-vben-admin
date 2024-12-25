package middleware

import (
	"fmt"
	"gin-vben-admin/pkg/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"time"
)

var timeFormat = "02/Jan/2006:15:04:05 +0800"

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.NewV4().String()
		logrus.AddHook(logger.NewTraceIdHook(traceId))
		c.Next()
	}
}

// Logger is the logrus logger handler
func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":    hostname,
			"statusCode":  statusCode,
			"latency":     latency, // time to process
			"client_ip":   clientIP,
			"method":      c.Request.Method,
			"path":        path,
			"referer":     referer,
			"data_length": dataLength,
			"user_agent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := ""
			switch logrus.StandardLogger().Formatter.(type) {
			case *logrus.JSONFormatter:
				break
			default:
				msg = fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			}
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}

		}
	}
}
