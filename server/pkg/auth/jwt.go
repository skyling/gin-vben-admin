package auth

import (
	"context"
	"fmt"
	"gin-vben-admin/dao/models"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/constant"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	UserID snowflake.ID `json:"u,omitempty"` // 用户ID
	Source string       `json:"s,omitempty"` // 来源
	Time   int64        `json:"t,omitempty"` // 时间
}

type token struct {
	Token   string
	UserId  snowflake.ID
	Source  string
	TimeOut time.Time
}

// GenToken 获取token
func GenToken(userID snowflake.ID, source string) (string, error) {
	var c = Claims{
		UserID: userID,
		Source: source,
		Time:   time.Now().Unix(),
	}
	sconf := global.Conf.System
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(sconf.TokenExpiredTime) * time.Second).Unix(),
		Issuer:    "gin-vben-admin",
		IssuedAt:  time.Now().Unix(),
	}

	// 使用指定的签名方法创建签名对象
	cl := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenStr, err := cl.SignedString([]byte(sconf.TokenSecret))

	if err != nil {
		return "", err
	}
	ctx := context.Background()
	key := buildCacheKey(userID, source, c.Time, false)
	// 是否允许多点登录
	if !sconf.UseMultipoint {
		c.Time = 0
		key = buildCacheKey(userID, source, 0, false)
		oldKey := buildCacheKey(userID, source, 0, true)
		// 重置旧token,将之前的token 重置为旧token
		if v, exists := global.TokenCache.Get(ctx, key, &token{}); exists == nil {
			vt := v.(*token)
			if time.Now().Sub(vt.TimeOut).Seconds() > float64(sconf.TokenOldExpiredTime) {
				vt.TimeOut = time.Now().Add(time.Duration(sconf.TokenOldExpiredTime) * time.Second)
			}
			global.TokenCache.Set(ctx, oldKey, vt, func(o *store.Options) {
				o.Expiration = time.Duration(sconf.TokenOldExpiredTime) * time.Second
			})
		}
	}

	// 设置新的token
	t := &token{
		Token:   tokenStr,
		TimeOut: time.Now().Add(time.Duration(sconf.TokenExpiredTime) * time.Second),
		UserId:  userID,
		Source:  source,
	}
	err = global.TokenCache.Set(ctx, key, t, func(o *store.Options) {
		o.Expiration = time.Duration(sconf.TokenExpiredTime) * time.Second
	})
	return tokenStr, err
}

// DeleteToken 删除Token
func DeleteToken(userId snowflake.ID, source string, t int64) bool {
	key := buildCacheKey(userId, source, t, false)
	err := global.TokenCache.Delete(context.Background(), key)
	return err == nil
}

// 构建token 过期时间
func buildCacheKey(userID snowflake.ID, source string, t int64, old bool) string {
	if old {
		return fmt.Sprintf("%s_%s_%d_old", userID, source, t)
	}
	return fmt.Sprintf("%d_%s_%d", userID, source, t)
}

// RefreshToken 刷新Token
func RefreshToken(tokenString string, source string, ti int64) (string, error) {
	// 解析token
	sconf := global.Conf.System
	t, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(sconf.TokenSecret), nil
	})
	ctx := context.Background()
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(*Claims); ok && t.Valid { // 校验token
		if source == "" {
			source = claims.Source
		}
		key := buildCacheKey(claims.UserID, source, ti, false)
		// 新key 是否存在
		if _, err := global.TokenCache.Get(ctx, key, &token{}); err == nil || source != "" {
			return GenToken(claims.UserID, source)
		}
	}
	return "", e.ErrUnauthorized
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	sconf := global.Conf.System
	// 解析token
	t, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(sconf.TokenSecret), nil
	})
	ctx := context.Background()
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*Claims); ok && t.Valid { // 校验token
		key := buildCacheKey(claims.UserID, claims.Source, claims.Time, false)
		// 新key 是否存在
		if tk, err := global.TokenCache.Get(ctx, key, &token{}); err == nil && tk.(*token).Token == tokenString {
			return claims, nil
		}
		// 旧key 是否存在
		oldKey := buildCacheKey(claims.UserID, claims.Source, claims.Time, true)
		if tk, err := global.TokenCache.Get(ctx, oldKey, &token{}); err == nil && tk.(*token).Token == tokenString {
			return claims, nil
		}
		return nil, e.ErrUnauthorized
	}
	return nil, e.ErrUnauthorized
}

// ParseUser 从token中解析出user
func ParseUser(c *gin.Context) error {
	token := GetTokenFromRequest(c)
	if token == "" {
		return e.ErrUnauthorized
	}
	claims, err := ParseToken(token)
	if err != nil {
		logrus.Infof("parse user error %v", err)
		return err
	}
	c.Set(constant.ClaimKey, claims)
	if err != nil {
		return err
	}
	u, err := repo.UserSrv.GetUserByID(claims.UserID)
	if err != nil {
		return err
	}
	c.Set(constant.UserKey, u)
	c.Set(constant.UserIDKey, u.ID)
	c.Set(constant.TenantIDKey, u.TenantID)
	return nil
}

// GetTokenFromRequest 从请求中获取Token
func GetTokenFromRequest(c *gin.Context) string {
	key := constant.HeadKeyToken
	token := c.Request.Header.Get(key)
	if token == "" {
		token, _ = c.Cookie(key)
	}

	if token == "" {
		token = c.Request.Header.Get("Authorization")
	}

	if token == "" {
		token = c.Query(key)
	}
	return token
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *models.User {
	if user, _ := c.Get(constant.UserKey); user != nil {
		if u, ok := user.(*models.User); ok {
			return u
		}
	}
	return nil
}

// CurrentUserID 获取当前用户
func CurrentUserID(c *gin.Context) snowflake.ID {
	if userId, _ := c.Get(constant.UserIDKey); userId != nil {
		if uid, ok := userId.(snowflake.ID); ok {
			return uid
		}
	}
	return -1
}

// CurrentTenantID 获取当前租户ID
func CurrentTenantID(c *gin.Context) snowflake.ID {
	if tenantId, _ := c.Get(constant.TenantIDKey); tenantId != nil {
		if uid, ok := tenantId.(snowflake.ID); ok {
			return uid
		}
	}
	return -1
}

// CurrentClaims 获取当前的claims
func CurrentClaims(c *gin.Context) *Claims {
	if c, exists := c.Get(constant.ClaimKey); exists {
		return c.(*Claims)
	}
	return nil
}

// CurrentSource 获取当前登录来源
func CurrentSource(c *gin.Context) string {
	if source, exists := c.Get(constant.SourceKey); exists {
		return source.(string)
	}
	return ""
}
