package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mylukin/easy-i18n/i18n"
	"golang.org/x/text/message"
	"strings"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang, _ := c.Cookie("lang")
		if lang == "" {
			lang, _ = c.GetQuery("lang")
		}
		accept := c.GetHeader("Accept-Language")
		accept = strings.Split(accept, "-")[0]
		fallback := "zh"
		tag := message.MatchLanguage(lang, accept, fallback)
		i18n.SetLang(tag)
		c.Writer.Header().Set("Cache-Control", "no-store,no-cache,must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Next()
	}
}
