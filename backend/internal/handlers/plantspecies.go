package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/deibys/sintronia/internal/db"
	"github.com/deibys/sintronia/internal/repositories"
	"github.com/deibys/sintronia/pkg/models"
	"github.com/gin-gonic/gin"
)

var plantRepo *repositories.PlantRepository

// Inicializar repository solo cuando DB esté disponible
func init() {
	// Este init se ejecutará después de que main() inicialice la DB
}

// getPlantRepo obtiene el repository, inicializándolo si es necesario
func getPlantRepo() *repositories.PlantRepository {
	if plantRepo == nil {
		if db.DB == nil {
			return nil // DB no disponible
		}
		plantRepo = repositories.NewPlantRepository()
	}
	return plantRepo
}

// CreatePlantSpeciesHandler maneja la creación de especies de plantas
func CreatePlantSpeciesHandler(c *gin.Context) {
	// Obtener información del usuario autenticado con verificación
	userID, exists := c.Get("user_id")
	if exists && userID != nil {
		c.Header("X-User-ID", strconv.FormatUint(uint64(userID.(int64)), 10))
		log.Printf("Usuario %v creando especie", userID)
	}

	userRole, exists := c.Get("user_role")
	if exists && userRole != nil {
		c.Header("X-User-Role", userRole.(string))
	}

	var req models.CreatePlantSpeciesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "JSON inválido: " + err.Error(),
		})
		return
	}

	// Verificar si external_id ya existe (si se proporciona)
	if req.ExternalRef != "" {
		exists, err := plantRepo.ExistsByExternalRef(req.ExternalRef)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   "Error verificando external_ref: " + err.Error(),
			})
			return
		}
		if exists {
			c.JSON(http.StatusConflict, models.APIResponse{
				Success: false,
				Error:   "Ya existe una planta con ese external_ref",
			})
			return
		}
	}

	// Crear nueva planta especie
	plant := models.PlantSpecies{
		CommonName:      req.CommonName,
		ScientificName:  req.ScientificName,
		Stratum:         req.Stratum,
		FunctionEcol:    req.FunctionEcol,
		SuccessionStage: req.SuccessionStage,
		ExternalRef:     req.ExternalRef,
		Notes:           req.Notes,
	}

	// Validar
	if err := plant.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Guardar en base de datos
	if err := plantRepo.Create(&plant); err != nil {
		log.Printf("Error creando planta: %v", err)
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Error guardando planta en base de datos",
		})
		return
	}

	log.Printf("Planta creada exitosamente: %s (ID: %d)", plant.CommonName, plant.ID)

	// Respuesta exitosa
	c.JSON(http.StatusCreated, models.APIResponse{
		Success: true,
		Data:    plant,
		Message: "Planta creada exitosamente",
	})
}

// GetPlantsSpeciesHandler maneja la obtención de todas las plantas
func GetPlantsSpeciesHandler(c *gin.Context) {
	// Verificar que la base de datos esté disponible
	repo := getPlantRepo()
	if repo == nil {
		c.JSON(http.StatusServiceUnavailable, models.APIResponse{
			Success: false,
			Error:   "Base de datos no disponible",
			Message: "El servicio está funcionando en modo limitado",
		})
		return
	}

	// Parámetros de paginación mejorados
	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed

			// Límite flexible basado en rol de usuario o configuración
			maxLimit := getMaxPaginationLimit(c)
			if limit > maxLimit {
				limit = maxLimit
			}
		}
	}

	// Construir filtros
	filters := repositories.PlantFilters{
		Search:          c.Query("search"),
		Stratum:         c.Query("stratum"),
		FunctionEcol:    c.Query("function_ecol"),
		SuccessionStage: c.Query("succession_stage"),
		Limit:           limit,
		Offset:          (page - 1) * limit,
	}

	// Obtener plantas de la base de datos
	plants, total, err := repo.GetAll(filters)
	if err != nil {
		log.Printf("Error obteniendo plantas: %v", err)
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Error obteniendo plantas de la base de datos",
		})
		return
	}

	// Calcular total de páginas
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	pagination := models.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, models.PaginatedResponse{
		Success:    true,
		Data:       plants,
		Pagination: pagination,
	})
}

// GetPlantSpeciesHandler maneja la obtención de una planta específica
func GetPlantSpeciesHandler(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	// Buscar planta en base de datos
	plant, err := plantRepo.GetByID(uint(id))
	if err != nil {
		if err.Error() == "planta no encontrada" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Error:   "Planta no encontrada",
			})
		} else {
			log.Printf("Error obteniendo planta: %v", err)
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   "Error obteniendo planta de la base de datos",
			})
		}
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    plant,
	})
}

// UpdatePlantSpeciesHandler maneja la actualización de una planta
func UpdatePlantSpeciesHandler(c *gin.Context) {
	// Obtener información del usuario autenticado con verificación
	userID, exists := c.Get("user_id")
	if exists && userID != nil {
		c.Header("X-User-ID", strconv.FormatUint(uint64(userID.(int64)), 10))
	}

	userRole, exists := c.Get("user_role")
	if exists && userRole != nil {
		c.Header("X-User-Role", userRole.(string))
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	var req models.UpdatePlantSpeciesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "JSON inválido: " + err.Error(),
		})
		return
	}

	// Construir mapa de actualizaciones
	updates := make(map[string]interface{})

	if req.CommonName != nil {
		updates["common_name"] = *req.CommonName
	}
	if req.ScientificName != nil {
		updates["scientific_name"] = *req.ScientificName
	}
	if req.Stratum != nil {
		updates["stratum"] = *req.Stratum
	}
	if req.FunctionEcol != nil {
		updates["function_ecol"] = *req.FunctionEcol
	}
	if req.SuccessionStage != nil {
		updates["succession_stage"] = *req.SuccessionStage
	}
	if req.ExternalRef != nil {
		updates["external_ref"] = *req.ExternalRef
	}
	if req.Notes != nil {
		updates["notes"] = *req.Notes
	}

	// Log de la actualización
	if userID != nil {
		log.Printf("Usuario %v actualizando planta ID: %d", userID, id)
	}

	// Actualizar en base de datos
	plant, err := plantRepo.Update(uint(id), updates)
	if err != nil {
		if err.Error() == "planta no encontrada" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Error:   "Planta no encontrada",
			})
		} else {
			log.Printf("Error actualizando planta: %v", err)
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   "Error actualizando planta en la base de datos",
			})
		}
		return
	}

	// Validar después de la actualización
	if err := plant.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    plant,
		Message: "Planta actualizada exitosamente",
	})
}

// DeletePlantSpeciesHandler maneja la eliminación de una planta
func DeletePlantSpeciesHandler(c *gin.Context) {
	// Obtener información del usuario autenticado con verificación
	userID, exists := c.Get("user_id")
	if exists && userID != nil {
		c.Header("X-User-ID", strconv.FormatUint(uint64(userID.(int64)), 10))
	}

	userRole, exists := c.Get("user_role")
	if exists && userRole != nil {
		c.Header("X-User-Role", userRole.(string))
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "ID inválido",
		})
		return
	}

	// Log de la eliminación
	if userID != nil {
		log.Printf("Usuario %v eliminando planta ID: %d", userID, id)
	}

	// Eliminar de base de datos
	if err := plantRepo.Delete(uint(id)); err != nil {
		if err.Error() == "planta no encontrada" {
			c.JSON(http.StatusNotFound, models.APIResponse{
				Success: false,
				Error:   "Planta no encontrada",
			})
		} else if err.Error() != "" && err.Error()[:50] == "no se puede eliminar la planta porque está siendo" {
			c.JSON(http.StatusConflict, models.APIResponse{
				Success: false,
				Error:   err.Error(),
			})
		} else {
			log.Printf("Error eliminando planta: %v", err)
			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   "Error eliminando planta de la base de datos",
			})
		}
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Planta eliminada exitosamente",
	})
}

// getMaxPaginationLimit determina el límite máximo de paginación basado en el rol del usuario
func getMaxPaginationLimit(c *gin.Context) int {
	// Obtener rol del usuario
	userRole, exists := c.Get("user_role")
	if exists && userRole != nil {
		role := userRole.(string)

		// Los administradores pueden tener límites más altos
		if role == "admin" {
			// Verificar si hay un límite configurado por variable de entorno
			if envLimit := os.Getenv("ADMIN_MAX_PAGINATION_LIMIT"); envLimit != "" {
				if limit, err := strconv.Atoi(envLimit); err == nil && limit > 0 {
					return limit
				}
			}
			return 1000 // Límite alto para admins
		}
	}

	// Límite por defecto para usuarios normales
	if envLimit := os.Getenv("DEFAULT_MAX_PAGINATION_LIMIT"); envLimit != "" {
		if limit, err := strconv.Atoi(envLimit); err == nil && limit > 0 {
			return limit
		}
	}

	return 100 // Límite por defecto
}
