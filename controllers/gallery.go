package controllers

import (
	"html/template"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tianfsg/vida-go/static/lang"
)

// Photo represents an image with its URL and orientation
type Photo struct {
	URL    string
	IsWide bool // Indicates if the image is wider than tall
	IsTall bool // Indicates if the image is taller than wide
}

// LoadPhotosFromDirectory loads photos from a directory based on offset and limit (for infinite scrolling)
func LoadPhotosFromDirectory(directory string, offset int, limit int) ([]Photo, error) {
	var photos []Photo
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var validPhotos []Photo
	for _, file := range files {
		if !file.IsDir() && isValidImage(file.Name()) {
			isWide, err := isImageWide(filepath.Join(directory, file.Name()))
			if err != nil {
				isWide = false
			}
			validPhotos = append(validPhotos, Photo{
				URL:    filepath.Join("/static/images/data", file.Name()),
				IsWide: isWide,
				IsTall: !isWide, // If it's not wide, it is considered tall
			})
		}
	}

	// Calculate the batch of photos based on offset and limit
	startIndex := offset
	endIndex := startIndex + limit
	if endIndex > len(validPhotos) {
		endIndex = len(validPhotos)
	}

	// Return the subset of photos for this batch
	photos = validPhotos[startIndex:endIndex]

	return photos, nil
}

// isValidImage checks if a file has a valid image extension
func isValidImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

// isImageWide checks if the image is wider than tall
func isImageWide(imagePath string) (bool, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false, err
	}

	bounds := img.Bounds()
	return bounds.Dx() > bounds.Dy(), nil
}

// GalleryController handles the gallery page requests
type GalleryController struct{}

// Gallery renders the gallery page
func (gc *GalleryController) Gallery(w http.ResponseWriter, r *http.Request) {
	// Obtener el idioma de la query o por defecto "en"
	langCode := r.URL.Query().Get("lang")
	if langCode == "" {
		langCode = "EN"
	}

	// Obtener los datos de idioma
	languageData := lang.GetLanguage(langCode)

	// Obtener las fotos de la galería
	dir := "./static/images/data"
	photos, err := LoadPhotosFromDirectory(dir, 0, 24) // Limite de 24 fotos, por ejemplo
	if err != nil {
		http.Error(w, "Unable to load gallery photos", http.StatusInternalServerError)
		return
	}

	// Parsear el template de la galería
	tmpl := template.Must(template.ParseFiles("views/gallery.html"))

	// Renderizar el template con los datos de idioma y las fotos
	tmpl.Execute(w, map[string]interface{}{
		"language": langCode,
		"data":     languageData,
		"photos":   photos,
	})
}
