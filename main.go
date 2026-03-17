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

	// Release mode — elimina todos los [GIN-debug] del log
	gin.SetMode(gin.ReleaseMode)

	if _, err := os.Stat(cfg.ViewsDir); os.IsNotExist(err) {
		log.Fatalf("views/ directory not found at %s", cfg.ViewsDir)
	}

	r := gin.Default()

	// Solo confiar en localhost (proxy de cPanel)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Static files
	r.Static("/static", cfg.StaticDir)

	// Templates ANTES de registrar rutas (thread-safe)
	funcMap := template.FuncMap{
		"add":       func(a, b int) int { return a + b },
		"sub":       func(a, b int) int { return a - b },
		"hasSuffix": strings.HasSuffix,
	}
	tmpl := template.New("").Funcs(funcMap)
	template.Must(tmpl.ParseGlob(cfg.ViewsDir + "/partials/*.html"))
	template.Must(tmpl.ParseGlob(cfg.ViewsDir + "/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Middleware
	r.Use(middleware.Lang())
	r.Use(func(c *gin.Context) {
		c.Request.ParseForm()
		c.Next()
	})

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

	log.Printf("Server running at http://localhost:%s", cfg.Port)
	log.Printf("ImageDir: %s", cfg.ImageDir)
	log.Printf("StaticDir: %s", cfg.StaticDir)
	log.Printf("ViewsDir: %s", cfg.ViewsDir)

	r.Run("0.0.0.0:" + cfg.Port)
}
