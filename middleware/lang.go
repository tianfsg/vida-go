package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tianfsg/vida-go/lang"
)

const LangKey = "languageData"
const LangCodeKey = "langCode"

// Lang extracts the ?lang= query param, resolves the language data,
// and attaches both to the Gin context for use in any handler.
func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("lang")
		if code == "" {
			code = "EN"
		}

		data := lang.GetLanguage(code)

		c.Set(LangCodeKey, code)
		c.Set(LangKey, data)

		c.Next()
	}
}
