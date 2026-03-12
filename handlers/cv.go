package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// VerifyCaptcha handles POST /verify-captcha.
// Stub for now — replace with real hCaptcha verification when ready.
func VerifyCaptcha(c *gin.Context) {
	token := c.PostForm("h-captcha-response")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No captcha token provided",
		})
		return
	}

	// TODO: verify token against hCaptcha API
	// https://docs.hcaptcha.com/#verify-the-user-response-server-side
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Captcha verified",
	})
}

// DownloadCV handles POST /download-cv.
func DownloadCV(c *gin.Context) {
	langCode := c.PostForm("language")
	if langCode == "" {
		langCode = "en"
	}

	captchaToken := c.PostForm("h-captcha-response")
	if captchaToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Captcha verification required",
		})
		return
	}

	// TODO: verify captchaToken against hCaptcha before serving the file

	filePath := filepath.Join("static", "content", langCode+"-cv.pdf")
	fileName := "Sebastian_GutierrezCV_" + langCode + ".pdf"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "CV not found for the requested language",
		})
		return
	}

	c.FileAttachment(filePath, fileName)
}

// Contact handles POST /contact.
func Contact(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	// TODO: add email sending (e.g. via SendGrid or SMTP)
	log.Printf("Contact form: name=%s email=%s message=%s", name, email, message)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Message received",
	})
}
