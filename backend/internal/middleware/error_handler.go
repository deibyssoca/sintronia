package middleware

import (
	"net/http"

	"github.com/deibys/sintronia/pkg/models"
	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware para manejo centralizado de errores
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Si hay errores, manejarlos
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			response := models.APIResponse{
				Success: false,
				Error:   err.Error(),
			}

			// Determinar código de estado basado en el tipo de error
			statusCode := http.StatusInternalServerError
			if c.Writer.Status() != http.StatusOK {
				statusCode = c.Writer.Status()
			}

			c.JSON(statusCode, response)
		}
	}
}
