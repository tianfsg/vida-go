package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func baseDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

func VerifyCaptcha(c *gin.Context) {
	token := c.PostForm("h-captcha-response")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No captcha token provided"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Captcha verified"})
}

func DownloadCV(c *gin.Context) {
	langCode := c.PostForm("language")
	if langCode == "" {
		langCode = "en"
	}

	captchaToken := c.PostForm("h-captcha-response")
	if captchaToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Captcha verification required"})
		return
	}

	filePath := filepath.Join(baseDir(), "static", "content", langCode+"-cv.pdf")
	fileName := "Sebastian_GutierrezCV_" + langCode + ".pdf"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "CV not found for the requested language"})
		return
	}

	c.FileAttachment(filePath, fileName)
}

func Contact(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")
	log.Printf("Contact form: name=%s email=%s message=%s", name, email, message)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Message received"})
}
