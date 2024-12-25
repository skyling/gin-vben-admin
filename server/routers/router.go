package routers

import (
	"gin-vben-admin/global"
	"gin-vben-admin/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"time"
)

var (
	swagHandler gin.HandlerFunc
)

// ApiLimit 接口请求次数限制
func ApiLimit(r *gin.Engine) {
	r.Use(middleware.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // 客户端IP
	}, func(c *gin.Context) (*rate.Limiter, int) {
		return rate.NewLimiter(rate.Every(100*time.Millisecond), 50), 5 // 0.1s内只能有10个请求
	}, func(c *gin.Context) {
		c.AbortWithStatus(429) // 超出次数后返回
	}))
}

// InitCORS 初始化跨域配置
func InitCORS(router *gin.Engine) {
	conf := global.Conf
	if conf.Cors.AllowOrigins != nil && conf.Cors.AllowOrigins[0] != "" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     conf.Cors.AllowOrigins,
			AllowMethods:     conf.Cors.AllowMethods,
			AllowHeaders:     conf.Cors.AllowHeaders,
			AllowCredentials: conf.Cors.AllowCredentials,
			ExposeHeaders:    conf.Cors.ExposeHeaders,
			AllowWildcard:    true,
			MaxAge:           12 * time.Hour,
		}))
		return
	}
}

func initGin() *gin.Engine {
	engine := gin.New()
	engine.Use(middleware.TraceID(), middleware.OperationRecord(logrus.StandardLogger()), gin.Recovery())
	if global.Conf.System.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		if swagHandler != nil {
			engine.GET("swagger/*any", swagHandler)
		}
	}
	return engine
}

func InitRouter() *gin.Engine {
	r := initGin()

	InitCORS(r)
	ApiLimit(r)

	AdminRouter(r)
	return r
}
