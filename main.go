package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tianfsg/vida-go/config"
	"github.com/tianfsg/vida-go/handlers"
	"github.com/tianfsg/vida-go/middleware"
)

func main() {
	cfg := config.Load()

	gin.SetMode(gin.ReleaseMode)

	if _, err := os.Stat(cfg.ViewsDir); os.IsNotExist(err) {
		log.Fatalf("[FATAL] views/ not found at: %s", cfg.ViewsDir)
	}

	log.Printf("[STARTUP] Port:      %s", cfg.Port)
	log.Printf("[STARTUP] StaticDir: %s", cfg.StaticDir)
	log.Printf("[STARTUP] ViewsDir:  %s", cfg.ViewsDir)
	log.Printf("[STARTUP] ImageDir:  %s", cfg.ImageDir)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Static("/static", cfg.StaticDir)

	funcMap := template.FuncMap{
		"add":       func(a, b int) int { return a + b },
		"sub":       func(a, b int) int { return a - b },
		"hasSuffix": strings.HasSuffix,
	}
	tmpl := template.New("").Funcs(funcMap)
	template.Must(tmpl.ParseGlob(cfg.ViewsDir + "/partials/*.html"))
	template.Must(tmpl.ParseGlob(cfg.ViewsDir + "/*.html"))
	r.SetHTMLTemplate(tmpl)

	r.Use(middleware.Lang())
	r.Use(func(c *gin.Context) {
		c.Request.ParseForm()
		c.Next()
	})

	// Health — usado por el watchdog
	r.GET("/health", handlers.Health)

	// Pages
	r.GET("/", handlers.Hub)
	r.GET("/portfolio", handlers.Index)
	r.GET("/privacy", handlers.Privacy)
	r.GET("/cookies", handlers.Cookies)
	r.GET("/terms", handlers.Terms)
	r.GET("/legal-notice", handlers.LegalNotice)
	r.GET("/talk", handlers.Talk)

	// Gallery
	r.GET("/gallery", handlers.GalleryPage(cfg))
	r.GET("/photos", handlers.PhotosAPI(cfg))

	// Forms
	r.POST("/verify-captcha", handlers.VerifyCaptcha)
	r.POST("/download-cv", handlers.DownloadCV)
	r.POST("/contact", handlers.Contact)

	// 404
	r.NoRoute(handlers.NotFound)

	log.Printf("[STARTUP] Listening on 0.0.0.0:%s", cfg.Port)
	if err := r.Run("0.0.0.0:" + cfg.Port); err != nil {
		log.Fatalf("[FATAL] Server failed to start: %v", err)
	}
}
