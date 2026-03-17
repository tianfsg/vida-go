package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Port       string
	ImageDir   string
	PhotoLimit int
	StaticDir  string
	ViewsDir   string
}

// exeDir returns the directory where the compiled binary lives.
// This ensures paths resolve correctly in cPanel where the working
// directory may differ from the binary location.
func exeDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	base := exeDir()

	imageDir := os.Getenv("IMAGE_DIR")
	if imageDir == "" {
		imageDir = filepath.Join(base, "static", "images", "data")
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = filepath.Join(base, "static")
	}

	viewsDir := os.Getenv("VIEWS_DIR")
	if viewsDir == "" {
		viewsDir = filepath.Join(base, "views")
	}

	return &Config{
		Port:       port,
		ImageDir:   imageDir,
		PhotoLimit: 24,
		StaticDir:  staticDir,
		ViewsDir:   viewsDir,
	}
}
