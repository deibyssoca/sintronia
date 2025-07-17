package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware configura CORS para permitir requests del frontend
func CORSMiddleware() gin.HandlerFunc {
	// Obtener orígenes permitidos desde variable de entorno
	allowedOrigins := []string{
		"http://localhost:5173",
		"http://localhost:3000",
		"http://127.0.0.1:5173",
		"http://127.0.0.1:3000",
		"https://stackblitz.com",
	}

	// Agregar orígenes desde variable de entorno si existe
	if envOrigins := os.Getenv("CORS_ALLOWED_ORIGINS"); envOrigins != "" {
		origins := strings.Split(envOrigins, ",")
		for _, origin := range origins {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	}

	// Configuración CORS más permisiva para desarrollo
	config := cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
			"X-Requested-With", "X-Permapeople-Key-Id", "X-Permapeople-Key-Secret",
		},
		ExposeHeaders: []string{
			"Content-Length", "X-User-ID", "X-User-Role",
		},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 horas
	}

	// En desarrollo, permitir todos los orígenes si está configurado
	if os.Getenv("GIN_MODE") == "debug" || os.Getenv("CORS_ALLOW_ALL") == "true" {
		config.AllowAllOrigins = true
		config.AllowCredentials = false // No se puede usar con AllowAllOrigins
	}

	return cors.New(config)
}
