package handlers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/deibys/sintronia/pkg/models"
// 	"github.com/gin-gonic/gin"
// )

// // Simulamos una base de datos en memoria por ahora
// var arrangements []models.Arrangement
// var arrangementIDCounter int64 = 1

// // CreatearrAngementHandler maneja la creación de la disposición (linea, isla etc)
// func CreateArrangementHandler(c *gin.Context) {
// 	var req models.CreateArrangementRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "JSON inválido: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Verificar que la ubicación existe
// 	locationExists := false
// 	for _, location := range locations {
// 		if location.ID == req.LocationID {
// 			locationExists = true
// 			break
// 		}
// 	}

// 	if !locationExists {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "La ubicación especificada no existe",
// 		})
// 		return
// 	}

// 	// Crear nuevo lecho
// 	arrangement := models.Arrangement{
// 		ID:           arrangementIDCounter,
// 		Name:         req.Name,
// 		LocationID:   req.LocationID,
// 		Type:         req.Type,
// 		Length:       req.Length,
// 		Width:        req.Width,
// 		Diameter:     req.Diameter,
// 		SoilType:     req.SoilType,
// 		PlantingMode: req.PlantingMode,
// 		Notes:        req.Notes,
// 		CreatedAt:    time.Now(),
// 		UpdatedAt:    time.Now(),
// 	}

// 	// Validar
// 	if err := arrangement.Validate(); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   err.Error(),
// 		})
// 		return
// 	}

// 	// Guardar en "base de datos"
// 	arrangements = append(arrangements, arrangement)
// 	arrangementIDCounter++

// 	// Respuesta exitosa con área calculada
// 	arrangementResponse := struct {
// 		models.Arrangement
// 		Area float64 `json:"area_m2"`
// 	}{
// 		Arrangement: arrangement,
// 		Area:        arrangement.CalculateArea(),
// 	}

// 	c.JSON(http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Data:    arrangementResponse,
// 		Message: "Lecho creado exitosamente",
// 	})
// }

// // GetArrangementsHandler maneja la obtención de todos los lechos
// func GetArrangementsHandler(c *gin.Context) {
// 	// Parámetros de consulta
// 	locationIDStr := c.Query("location_id")
// 	arrangementType := c.Query("type")

// 	// Filtrar lechos
// 	var filteredArrangements []models.Arrangement
// 	for _, arrangement := range arrangements {
// 		// Filtro por ubicación
// 		if locationIDStr != "" {
// 			locationID, err := strconv.ParseInt(locationIDStr, 10, 64)
// 			if err == nil && arrangement.LocationID != locationID {
// 				continue
// 			}
// 		}

// 		// Filtro por tipo
// 		if arrangementType != "" && arrangement.Type != arrangementType {
// 			continue
// 		}

// 		filteredArrangements = append(filteredArrangements, arrangement)
// 	}

// 	// Agregar área calculada a cada lecho
// 	var arrangementsWithArea []interface{}
// 	for _, arrangement := range filteredArrangements {
// 		arrangementWithArea := struct {
// 			models.Arrangement
// 			Area float64 `json:"area_m2"`
// 		}{
// 			Arrangement: arrangement,
// 			Area:        arrangement.CalculateArea(),
// 		}
// 		arrangementsWithArea = append(arrangementsWithArea, arrangementWithArea)
// 	}

// 	c.JSON(http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Data:    arrangementsWithArea,
// 	})
// }

// // GetArrangementHandler maneja la obtención de un lecho específico
// func GetArrangementHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	// Buscar lecho
// 	for _, arrangement := range arrangements {
// 		if arrangement.ID == id {
// 			arrangementWithArea := struct {
// 				models.Arrangement
// 				Area float64 `json:"area_m2"`
// 			}{
// 				Arrangement: arrangement,
// 				Area:        arrangement.CalculateArea(),
// 			}

// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Data:    arrangementWithArea,
// 			})
// 			return
// 		}
// 	}

// 	// No encontrado
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Lecho no encontrado",
// 	})
// }

// // DeleteArrangementHandler maneja la eliminación de un lecho
// func DeleteArrangementHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	// Verificar si hay plantaciones asociadas
// 	hasPlantings := false
// 	for _, planting := range plantings {
// 		if planting.ArrangementID == id {
// 			hasPlantings = true
// 			break
// 		}
// 	}

// 	if hasPlantings {
// 		c.JSON(http.StatusConflict, models.APIResponse{
// 			Success: false,
// 			Error:   "No se puede eliminar el lecho porque tiene plantaciones asociadas",
// 		})
// 		return
// 	}

// 	// Buscar y eliminar lecho
// 	for i, arrangement := range arrangements {
// 		if arrangement.ID == id {
// 			arrangements = append(arrangements[:i], arrangements[i+1:]...)

// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Message: "Lecho eliminado exitosamente",
// 			})
// 			return
// 		}
// 	}

// 	// No encontrado
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Lecho no encontrado",
// 	})
// }
