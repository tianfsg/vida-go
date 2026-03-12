package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tianfsg/vida-go/lang"
	"github.com/tianfsg/vida-go/middleware"
)

// getLang is a helper to pull language data from the Gin context.
func getLang(c *gin.Context) (string, lang.Language) {
	code := c.GetString(middleware.LangCodeKey)
	data, _ := c.Get(middleware.LangKey)
	return code, data.(lang.Language)
}

func Index(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "portfolio.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func Hub(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "hub.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func Privacy(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "privacy.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func Cookies(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "cookies.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func Terms(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "terms.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func LegalNotice(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "legal.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func Talk(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusOK, "talk.html", gin.H{
		"language": code,
		"data":     data,
	})
}

func NotFound(c *gin.Context) {
	code, data := getLang(c)
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"language": code,
		"data":     data,
	})
}
