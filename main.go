package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tianfsg/vida-go/static/lang"
)

// Función para obtener el idioma de la solicitud
func getLanguage(c *gin.Context) string {
	langCode := c.Query("lang")
	if langCode == "" {
		langCode = "EN"
	}
	return langCode
}

// Custom function for addition
func add(a, b int) int {
	return a + b
}

// Custom function for subtraction
func sub(a, b int) int {
	return a - b
}

// Función para validar que el archivo es una imagen
func isValidImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif"
}

func main() {
	// Verificar si la carpeta de templates existe
	if _, err := os.Stat("./views"); os.IsNotExist(err) {
		log.Fatalf("Error: El directorio './views' no existe\n")
	}

	// Verificar si el archivo index.html existe
	if _, err := os.Stat("./views/index.html"); os.IsNotExist(err) {
		log.Fatalf("Error: El archivo 'index.html' no se encuentra en './views'\n")
	} else {
		log.Println("Archivo 'index.html' encontrado correctamente.")
	}

	// Crear un router Gin
	r := gin.Default()

	// Configurar la carpeta de archivos estáticos
	r.Static("/static", "./static")

	// Configurar el motor de plantillas HTML
	r.SetFuncMap(template.FuncMap{
		"t":   func(key string) string { return key },
		"add": add,
		"sub": sub,
	})

	// Cargar los templates
	r.LoadHTMLGlob("./views/*.html")

	// Middleware para manejar formularios y JSON
	r.Use(func(c *gin.Context) {
		c.Request.ParseForm()
		c.Next()
	})

	// Ruta principal
	r.GET("/portfolio", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)
		// fmt.Println(langCode, languageData)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para la página de privacidad
	r.GET("/privacy", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusOK, "privacy.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para la página de cookies
	r.GET("/cookies", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusOK, "cookies.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para la página de términos
	r.GET("/terms", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusOK, "terms.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para la página de aviso legal
	r.GET("/legal-notice", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusOK, "legal.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para la página de charla
	r.GET("/talk", func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusOK, "talk.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Ruta para verificar hCaptcha
	r.POST("/verify-captcha", func(c *gin.Context) {
		token := c.PostForm("h-captcha-response")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "No captcha token provided",
			})
			return
		}

		// Aquí iría la verificación real del token de hCaptcha
		// Por ahora, simulamos una verificación exitosa
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Captcha verified successfully",
		})
	})

	// Ruta para descargar el CV
	r.POST("/download-cv", func(c *gin.Context) {
		lang := c.PostForm("language")
		if lang == "" {
			lang = "en"
		}

		// Verificar token hCaptcha aquí si es necesario
		captchaToken := c.PostForm("h-captcha-response")
		if captchaToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Captcha verification required",
			})
			return
		}

		filePath := filepath.Join("static", "content", "cv", lang+"-cv.pdf")
		fileName := "Sebastian_GutierrezCV_" + lang + ".pdf"

		// Verificar si el archivo existe
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "CV not found for specified language",
			})
			return
		}

		c.FileAttachment(filePath, fileName)
	})

	// Ruta para manejar el formulario de contacto
	r.POST("/contact", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		message := c.PostForm("message")

		// Aquí irían las validaciones y el envío del email
		// Por ahora solo registramos en el log
		log.Printf("Nuevo mensaje de contacto: Nombre=%s, Email=%s, Mensaje=%s",
			name, email, message)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Message sent successfully",
		})
	})

	// Ruta para la galería de fotos
	r.GET("/gallery", func(c *gin.Context) {
		// Obtener el idioma de la consulta (por ejemplo, "?lang=ES")
		langCode := getLanguage(c)

		// Obtener los datos de idioma
		languageData := lang.GetLanguage(langCode)

		// Directorio de fotos
		dir := "./static/images/data"

		// Leer el directorio y obtener los nombres de los archivos
		files, err := os.ReadDir(dir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read photos directory"})
			return
		}

		// Crear una lista de URLs de fotos
		var photos []string
		for _, file := range files {
			if !file.IsDir() && isValidImage(file.Name()) {
				photoURL := "/static/images/data/" + file.Name()
				photos = append(photos, photoURL)
			}
		}

		// Pasar las fotos y los datos de idioma al frontend
		c.HTML(http.StatusOK, "gallery.html", gin.H{
			"language": langCode,
			"data":     languageData,
			"photos":   photos,
		})
	})

	// Middleware para manejar errores 404
	r.NoRoute(func(c *gin.Context) {
		langCode := getLanguage(c)
		languageData := lang.GetLanguage(langCode)

		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"language": langCode,
			"data":     languageData,
		})
	})

	// Puerto de escucha (local)
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8060"
	// }

	//Puerto de escucha (web)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Usa un puerto no privilegiado
	}

	// Iniciar el servidor
	log.Printf("Servidor iniciado en http://localhost:%s", port)
	r.Run("0.0.0.0:" + port)
}
