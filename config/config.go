package config

import (
	"os"
)

type Config struct {
	Port       string
	ImageDir   string
	PhotoLimit int
	StaticDir  string
	ViewsDir   string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	imageDir := os.Getenv("IMAGE_DIR")
	if imageDir == "" {
		imageDir = "./static/images/data"
	}

	return &Config{
		Port:       port,
		ImageDir:   imageDir,
		PhotoLimit: 24,
		StaticDir:  "./static",
		ViewsDir:   "./views",
	}
}
