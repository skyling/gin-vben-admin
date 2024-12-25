package middleware

import (
	"gin-vben-admin/global"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

// NewRateLimiter IP地址流量限制
func NewRateLimiter(key func(*gin.Context) string, createLimiter func(*gin.Context) (*rate.Limiter, int),
	abort func(*gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		k := key(c)
		limiter, err := global.BigCache.Get(c, k, &rate.Limiter{})
		if err != nil {
			var expire int
			limiter, expire = createLimiter(c)
			global.BigCache.Set(c, k, store.WithExpiration(time.Duration(expire)*time.Second))
		}
		ok := limiter.(*rate.Limiter).Allow()
		if !ok {
			abort(c)
			return
		}
		c.Next()
	}
}
