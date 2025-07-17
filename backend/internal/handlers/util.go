package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SaludoHandler responde con un mensaje de saludo.
func SaludoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "!Hola desde Sintronia con Gin!",
	})
}
