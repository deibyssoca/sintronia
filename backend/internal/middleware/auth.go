package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/deibys/sintronia/pkg/models"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica la existencia y validez de un token de autorización en el header.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Espera que el header Authorization tenga el formato "Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		keyID := c.GetHeader("x-permapeople-key-id")

		log.Printf("authHeader: %s | keyID: %s", authHeader, keyID)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") || keyID == "" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Token de autorización requerido",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" { //|| parts[1] != "your-secret-token" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Formato de token inválido. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		// Si el token es correcto, continúa con la solicitud
		c.Next()
	}
}

// AuthMiddleware middleware de autenticación
func AuthMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Token de autorización requerido",
			})
			c.Abort()
			return
		}

		// Verificar formato Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Formato de token inválido. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Validar el token (por ahora simulamos la validación)
		// En producción aquí validarías JWT, consultarías base de datos, etc.
		if !isValidToken(token) {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Token inválido o expirado",
			})
			c.Abort()
			return
		}

		// Obtener información del usuario (simulado)
		userID, userRole := getUserFromToken(token)

		// Agregar información del usuario al contexto
		c.Set("user_id", userID)
		c.Set("user_role", userRole)

		// Continuar con el siguiente handler
		c.Next()
	}
}

// AdminMiddleware middleware que requiere rol de administrador
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				Success: false,
				Error:   "Usuario no autenticado",
			})
			c.Abort()
			return
		}

		if userRole != "admin" {
			c.JSON(http.StatusForbidden, models.APIResponse{
				Success: false,
				Error:   "Acceso denegado. Se requieren permisos de administrador",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Funciones auxiliares (simuladas - reemplazar con lógica real)
func isValidToken(token string) bool {
	// Por ahora aceptamos cualquier token que no esté vacío
	// En producción: validar JWT, verificar en base de datos, etc.
	validTokens := map[string]bool{
		"test-token":  true,
		"admin-token": true,
		"user-token":  true,
	}

	return validTokens[token]
}

func getUserFromToken(token string) (int64, string) {
	// Simulación - en producción extraer del JWT o consultar base de datos
	tokenUsers := map[string]struct {
		ID   int64
		Role string
	}{
		"test-token":  {ID: 1, Role: "user"},
		"admin-token": {ID: 2, Role: "admin"},
		"user-token":  {ID: 3, Role: "user"},
	}

	if user, exists := tokenUsers[token]; exists {
		return user.ID, user.Role
	}

	return 0, "guest"
}

// OptionalAuthMiddleware middleware de autenticación opcional
// No bloquea si no hay token, pero agrega info del usuario si existe
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				token := tokenParts[1]
				if isValidToken(token) {
					userID, userRole := getUserFromToken(token)
					c.Set("user_id", userID)
					c.Set("user_role", userRole)
				}
			}
		}
		c.Next()
	}
}
