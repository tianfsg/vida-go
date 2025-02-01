// vida-go/controllers/home.go

package controllers

import (
	"html/template"
	"net/http"

	"github.com/tianfsg/vida-go/static/lang"
)

type HomeController struct{}

func (hc *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	// Get language from query parameter or default to "en"
	langCode := r.URL.Query().Get("lang")
	if langCode == "" {
		langCode = "EN"
	}

	// Get language data
	languageData := lang.GetLanguage(langCode)

	// Parse template
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	// Render template with language data
	tmpl.Execute(w, languageData)
}
