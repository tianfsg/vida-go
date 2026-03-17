package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tianfsg/vida-go/config"
)

// Photo represents a single image with its URLs.
type Photo struct {
	URL  string
	WebP string
}

// isValidImage checks if a filename has a supported image extension.
func isValidImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp"
}

// loadPhotos scans the image directory and returns Photo objects.
// If the file is already a .webp it is served directly as URL (no separate WebP field needed).
// If it's a JPG/PNG, checks for a .webp sibling and sets the WebP field for the <source> element.
func loadPhotos(cfg *config.Config, offset, limit int) ([]Photo, error) {
	files, err := os.ReadDir(cfg.ImageDir)
	if err != nil {
		return nil, err
	}

	var photos []Photo
	count := 0

	for _, file := range files {
		if file.IsDir() || !isValidImage(file.Name()) {
			continue
		}

		if count < offset {
			count++
			continue
		}

		filename := file.Name()
		ext := strings.ToLower(filepath.Ext(filename))
		base := strings.TrimSuffix(filename, filepath.Ext(filename))

		var photo Photo

		if ext == ".webp" {
			// Serve webp directly — no JPG fallback needed
			photo = Photo{
				URL:  "/static/images/data/" + filename,
				WebP: "",
			}
		} else {
			// JPG/PNG — check for webp sibling
			webpPath := filepath.Join(cfg.ImageDir, base+".webp")
			photo = Photo{
				URL: "/static/images/data/" + filename,
			}
			if _, err := os.Stat(webpPath); err == nil {
				photo.WebP = "/static/images/data/" + base + ".webp"
			}
		}

		photos = append(photos, photo)
		count++

		if limit > 0 && len(photos) >= limit {
			break
		}
	}

	return photos, nil
}

// GalleryPage handles GET /gallery — server-side rendered page with initial photos.
func GalleryPage(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, data := getLang(c)

		photos, err := loadPhotos(cfg, 0, cfg.PhotoLimit)
		if err != nil {
			log.Printf("Gallery: failed to load photos: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load gallery photos"})
			return
		}

		c.HTML(http.StatusOK, "gallery.html", gin.H{
			"language": code,
			"data":     data,
			"photos":   photos,
		})
	}
}

// PhotosAPI handles GET /photos?offset=N — JSON endpoint for infinite scroll.
func PhotosAPI(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil || offset < 0 {
			offset = 0
		}

		photos, err := loadPhotos(cfg, offset, cfg.PhotoLimit)
		if err != nil {
			log.Printf("PhotosAPI: failed to load photos: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load photos"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"photos": photos,
			"offset": offset + len(photos),
		})
	}
}
