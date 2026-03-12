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

	if _, err := os.Stat("./views"); os.IsNotExist(err) {
		log.Fatal("views/ directory not found")
	}

	r := gin.Default()

	// Static files
	r.Static("/static", cfg.StaticDir)

	// Build a single template set: partials first, then pages.
	// Both globs must be parsed into the SAME template.Template so that
	// page templates can call {{template "footer" .}} etc.
	funcMap := template.FuncMap{
		"add":       func(a, b int) int { return a + b },
		"sub":       func(a, b int) int { return a - b },
		"hasSuffix": strings.HasSuffix,
	}

	tmpl := template.New("").Funcs(funcMap)
	template.Must(tmpl.ParseGlob("./views/partials/*.html"))
	template.Must(tmpl.ParseGlob("./views/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Language middleware
	r.Use(middleware.Lang())

	// Parse form data for POST routes
	r.Use(func(c *gin.Context) {
		c.Request.ParseForm()
		c.Next()
	})

	// ── Pages ──────────────────────────────────────────────
	r.GET("/", handlers.Hub)
	r.GET("/portfolio", handlers.Index)
	r.GET("/privacy", handlers.Privacy)
	r.GET("/cookies", handlers.Cookies)
	r.GET("/terms", handlers.Terms)
	r.GET("/legal-notice", handlers.LegalNotice)
	r.GET("/talk", handlers.Talk)

	// ── Gallery ────────────────────────────────────────────
	r.GET("/gallery", handlers.GalleryPage(cfg))
	r.GET("/photos", handlers.PhotosAPI(cfg))

	// ── Forms & Downloads ──────────────────────────────────
	r.POST("/verify-captcha", handlers.VerifyCaptcha)
	r.POST("/download-cv", handlers.DownloadCV)
	r.POST("/contact", handlers.Contact)

	// ── 404 ────────────────────────────────────────────────
	r.NoRoute(handlers.NotFound)

	log.Printf("Server running at http://localhost:%s", cfg.Port)
	r.Run("0.0.0.0:" + cfg.Port)
}
