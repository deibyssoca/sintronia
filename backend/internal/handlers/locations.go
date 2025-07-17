package handlers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/deibys/sintronia/pkg/models"
// 	"github.com/gin-gonic/gin"
// )

// // Simulamos una base de datos en memoria por ahora
// var locations []models.Location
// var locationIDCounter int64 = 1

// // CreateLocationHandler maneja la creación de ubicaciones
// func CreateLocationHandler(c *gin.Context) {
// 	var req models.CreateLocationRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "JSON inválido: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Crear nueva ubicación
// 	location := models.Location{
// 		ID:        locationIDCounter,
// 		Name:      req.Name,
// 		Notes:     req.Notes,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Validar
// 	if err := location.Validate(); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   err.Error(),
// 		})
// 		return
// 	}

// 	// Guardar en "base de datos"
// 	locations = append(locations, location)
// 	locationIDCounter++

// 	// Respuesta exitosa
// 	c.JSON(http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Data:    location,
// 		Message: "Ubicación creada exitosamente",
// 	})
// }

// // GetLocationsHandler maneja la obtención de todas las ubicaciones
// func GetLocationsHandler(c *gin.Context) {
// 	// Parámetros de paginación
// 	page := 1
// 	limit := 10

// 	if p := c.Query("page"); p != "" {
// 		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
// 			page = parsed
// 		}
// 	}

// 	if l := c.Query("limit"); l != "" {
// 		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
// 			limit = parsed
// 		}
// 	}

// 	// Calcular offset
// 	offset := (page - 1) * limit
// 	total := int64(len(locations))

// 	// Aplicar paginación
// 	var paginatedLocations []models.Location
// 	if offset < len(locations) {
// 		end := offset + limit
// 		if end > len(locations) {
// 			end = len(locations)
// 		}
// 		paginatedLocations = locations[offset:end]
// 	}

// 	// Calcular total de páginas
// 	totalPages := int(total) / limit
// 	if int(total)%limit != 0 {
// 		totalPages++
// 	}

// 	pagination := models.Pagination{
// 		Page:       page,
// 		Limit:      limit,
// 		Total:      total,
// 		TotalPages: totalPages,
// 	}

// 	c.JSON(http.StatusOK, models.PaginatedResponse{
// 		Success:    true,
// 		Data:       paginatedLocations,
// 		Pagination: pagination,
// 	})
// }

// // GetLocationHandler maneja la obtención de una ubicación específica
// func GetLocationHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	// Buscar ubicación
// 	for _, location := range locations {
// 		if location.ID == id {
// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Data:    location,
// 			})
// 			return
// 		}
// 	}

// 	// No encontrada
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Ubicación no encontrada",
// 	})
// }

// // DeleteLocationHandler maneja la eliminación de una ubicación
// func DeleteLocationHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	// Buscar y eliminar ubicación
// 	for i, location := range locations {
// 		if location.ID == id {
// 			// Eliminar de la slice
// 			locations = append(locations[:i], locations[i+1:]...)

// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Message: "Ubicación eliminada exitosamente",
// 			})
// 			return
// 		}
// 	}

// 	// No encontrada
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Ubicación no encontrada",
// 	})
// }
