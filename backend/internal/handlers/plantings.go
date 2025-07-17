package handlers

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/deibys/sintronia/pkg/models"
// 	"github.com/gin-gonic/gin"
// )

// // Simulamos una base de datos en memoria por ahora
// var plantings []models.Planting
// var plantingIDCounter int64 = 1

// // CreatePlantingHandler maneja la creación de plantaciones
// func CreatePlantingHandler(c *gin.Context) {
// 	var req models.CreatePlantingRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "JSON inválido: " + err.Error(),
// 		})
// 		return
// 	}

// 	// Verificar que el lecho existe
// 	var targetArrangement *models.Arrangement
// 	for _, arrangement := range arrangements {
// 		if arrangement.ID == req.ArrangementID {
// 			targetArrangement = &arrangement
// 			break
// 		}
// 	}

// 	if targetArrangement == nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "El lecho especificado no existe",
// 		})
// 		return
// 	}

// 	// Verificar que la planta existe
// 	// var targetPlant *models.Plant
// 	// for _, plant := range plants {
// 	// 	if plant.ID == req.PlantID {
// 	// 		targetPlant = &plant
// 	// 		break
// 	// 	}
// 	// }

// 	// if targetPlant == nil {
// 	// 	c.JSON(http.StatusBadRequest, models.APIResponse{
// 	// 		Success: false,
// 	// 		Error:   "La planta especificada no existe",
// 	// 	})
// 	// 	return
// 	// }

// 	// Crear nueva plantación
// 	planting := models.Planting{
// 		// ID:            plantingIDCounter,
// 		ArrangementID: req.ArrangementID,
// 		PlantID:       req.PlantID,
// 		Quantity:      req.Quantity,
// 		Status:        req.Status,
// 		Position:      req.Position,
// 		Notes:         req.Notes,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}

// 	// Validar
// 	if err := planting.Validate(); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   err.Error(),
// 		})
// 		return
// 	}

// 	// Guardar en "base de datos"
// 	plantings = append(plantings, planting)
// 	plantingIDCounter++

// 	// Respuesta exitosa con información adicional
// 	plantingResponse := struct {
// 		models.Planting
// 		Plant       models.Plant       `json:"plant"`
// 		Arrangement models.Arrangement `json:"arrangement"`
// 		Density     float64            `json:"density_per_m2"`
// 	}{
// 		Planting:    planting,
// 		Plant:       *targetPlant,
// 		Arrangement: *targetArrangement,
// 		Density:     planting.CalculateDensity(targetArrangement.CalculateArea()),
// 	}

// 	c.JSON(http.StatusCreated, models.APIResponse{
// 		Success: true,
// 		Data:    plantingResponse,
// 		Message: "Plantación creada exitosamente",
// 	})
// }

// // GetPlantingsHandler maneja la obtención de todas las plantaciones
// func GetPlantingsHandler(c *gin.Context) {
// 	// Parámetros de consulta
// 	arrangementIDStr := c.Query("arrangement_id")
// 	plantIDStr := c.Query("plant_id")
// 	status := c.Query("status")

// 	// Filtrar plantaciones
// 	var filteredPlantings []models.Planting
// 	for _, planting := range plantings {
// 		// Filtro por lecho
// 		if arrangementIDStr != "" {
// 			arrangementID, err := strconv.ParseInt(arrangementIDStr, 10, 64)
// 			if err == nil && planting.ArrangementID != arrangementID {
// 				continue
// 			}
// 		}

// 		// Filtro por planta
// 		if plantIDStr != "" {
// 			plantID, err := strconv.ParseInt(plantIDStr, 10, 64)
// 			if err == nil && planting.PlantID != plantID {
// 				continue
// 			}
// 		}

// 		// Filtro por estado
// 		if status != "" && planting.Status != status {
// 			continue
// 		}

// 		filteredPlantings = append(filteredPlantings, planting)
// 	}

// 	// Enriquecer con información de plantas y lechos
// 	var enrichedPlantings []interface{}
// 	for _, planting := range filteredPlantings {
// 		// Buscar planta
// 		var plant models.Plant
// 		for _, p := range plants {
// 			if p.ID == planting.PlantID {
// 				plant = p
// 				break
// 			}
// 		}

// 		// Buscar lecho
// 		var arrangement models.Arrangement
// 		for _, b := range arrangements {
// 			if b.ID == planting.ArrangementID {
// 				arrangement = b
// 				break
// 			}
// 		}

// 		enrichedPlanting := struct {
// 			models.Planting
// 			Plant       models.Plant       `json:"plant"`
// 			Arrangement models.Arrangement `json:"arrangement"`
// 			Density     float64            `json:"density_per_m2"`
// 		}{
// 			Planting:    planting,
// 			Plant:       plant,
// 			Arrangement: arrangement,
// 			Density:     planting.CalculateDensity(arrangement.CalculateArea()),
// 		}

// 		enrichedPlantings = append(enrichedPlantings, enrichedPlanting)
// 	}

// 	c.JSON(http.StatusOK, models.APIResponse{
// 		Success: true,
// 		Data:    enrichedPlantings,
// 	})
// }

// // UpdatePlantingStatusHandler actualiza el estado de una plantación
// func UpdatePlantingStatusHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	var req struct {
// 		Status string `json:"status" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "JSON inválido: " + err.Error(),
// 		})
// 		return
// 	}

// 	if !models.IsValidStatus(req.Status) {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "Estado inválido",
// 		})
// 		return
// 	}

// 	// Buscar y actualizar plantación
// 	for i, planting := range plantings {
// 		if planting.ID == id {
// 			plantings[i].Status = req.Status
// 			plantings[i].UpdatedAt = time.Now()

// 			// Si se marca como plantada, registrar fecha
// 			if req.Status == models.StatusPlanted && plantings[i].PlantedAt == nil {
// 				now := time.Now()
// 				plantings[i].PlantedAt = &now
// 			}

// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Data:    plantings[i],
// 				Message: "Estado actualizado exitosamente",
// 			})
// 			return
// 		}
// 	}

// 	// No encontrada
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Plantación no encontrada",
// 	})
// }

// // DeletePlantingHandler maneja la eliminación de una plantación
// func DeletePlantingHandler(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.APIResponse{
// 			Success: false,
// 			Error:   "ID inválido",
// 		})
// 		return
// 	}

// 	// Buscar y eliminar plantación
// 	for i, planting := range plantings {
// 		if planting.ID == id {
// 			plantings = append(plantings[:i], plantings[i+1:]...)

// 			c.JSON(http.StatusOK, models.APIResponse{
// 				Success: true,
// 				Message: "Plantación eliminada exitosamente",
// 			})
// 			return
// 		}
// 	}

// 	// No encontrada
// 	c.JSON(http.StatusNotFound, models.APIResponse{
// 		Success: false,
// 		Error:   "Plantación no encontrada",
// 	})
// }
