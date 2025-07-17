package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Plant representa el inventario de plantas.
// Los datos pueden provenir de la API de Permapeople o ingresarse manualmente.
type Plant struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"not null;index"`
	Scientific      string         `json:"scientific" gorm:"index"`
	Stratum         string         `json:"stratum" gorm:"index"`               // Ej: "bajo", "medio", "alto"
	Function        string         `json:"function" gorm:"index"`              // "objetivo" o "servicio"
	SuccessionStage string         `json:"succession_stage" gorm:"index"`      // Ej: "pionera", "secundaria", "climax"
	ExternalID      string         `json:"external_id" gorm:"uniqueIndex"`     // Referencia a la API externa
	Desired         bool           `json:"desired" gorm:"default:false;index"` // Lista de plantas deseadas
	Notes           string         `json:"notes" gorm:"type:text"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Plantings []Planting `json:"plantings,omitempty" gorm:"foreignKey:PlantID"`
}

// Validate valida los datos de una planta
func (p *Plant) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("el nombre de la planta es requerido")
	}

	if p.Stratum != "" && !IsValidStratum(p.Stratum) {
		return errors.New("estrato inválido")
	}

	if p.Function != "" && !IsValidFunction(p.Function) {
		return errors.New("función inválida")
	}

	if p.SuccessionStage != "" && !IsValidSuccessionStage(p.SuccessionStage) {
		return errors.New("etapa sucesional inválida")
	}

	return nil
}

// Location representa una zona o área para el inventario (sin geolocalización por ahora).
type Location struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;uniqueIndex"`
	Notes     string         `json:"notes,omitempty" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Arrangements []Arrangement `json:"arrangements,omitempty" gorm:"foreignKey:LocationID"`
}

// Validate valida los datos de una ubicación
func (l *Location) Validate() error {
	if strings.TrimSpace(l.Name) == "" {
		return errors.New("el nombre de la ubicación es requerido")
	}
	return nil
}

// Arrangement representa la disposición de cultivo, que puede ser lineal (línea) o circular (isla).
type Arrangement struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	LocationID   uint           `json:"location_id" gorm:"not null;index"`
	Type         string         `json:"type" gorm:"not null;index"` // "linea", "isla", "gremio"
	Length       float64        `json:"length,omitempty"`           // Solo para líneas
	Width        float64        `json:"width,omitempty"`            // Solo para líneas
	Diameter     float64        `json:"diameter,omitempty"`         // Solo para islas
	SoilType     string         `json:"soil_type" gorm:"index"`
	PlantingMode string         `json:"planting_mode" gorm:"index"` // Modalidad de plantación
	Notes        string         `json:"notes,omitempty" gorm:"type:text"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Location  Location   `json:"location,omitempty" gorm:"foreignKey:LocationID"`
	Plantings []Planting `json:"plantings,omitempty" gorm:"foreignKey:ArrangementID"`
}

// Validate valida los datos de un lecho
func (b *Arrangement) Validate() error {
	if strings.TrimSpace(b.Name) == "" {
		return errors.New("el nombre del lecho es requerido")
	}

	if b.LocationID <= 0 {
		return errors.New("la ubicación es requerida")
	}

	if !IsValidArrangementType(b.Type) {
		return errors.New("tipo de lecho inválido")
	}

	// Validaciones específicas por tipo
	switch b.Type {
	case ArrangementTypeLine:
		if b.Length <= 0 || b.Width <= 0 {
			return errors.New("las líneas requieren longitud y ancho válidos")
		}
	case ArrangementTypeIsland:
		if b.Diameter <= 0 {
			return errors.New("las islas requieren un diámetro válido")
		}
	}

	if b.SoilType != "" && !IsValidSoilType(b.SoilType) {
		return errors.New("tipo de suelo inválido")
	}

	if b.PlantingMode != "" && !IsValidPlantingMode(b.PlantingMode) {
		return errors.New("modalidad de plantación inválida")
	}

	return nil
}

// CalculateArea calcula el área del lecho en metros cuadrados
func (b *Arrangement) CalculateArea() float64 {
	switch b.Type {
	case ArrangementTypeLine:
		return b.Length * b.Width
	case ArrangementTypeIsland:
		radius := b.Diameter / 2
		return 3.14159 * radius * radius
	default:
		return 0
	}
}

// Planting vincula una planta a un lecho de cultivo.
type Planting struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	ArrangementID uint           `json:"arrangement_id" gorm:"not null;index"`
	PlantID       uint           `json:"plant_id" gorm:"not null;index"`
	Quantity      int            `json:"quantity" gorm:"not null;check:quantity > 0"` // Puede haber varias instancias de la misma planta
	Status        string         `json:"status" gorm:"not null;index"`                // Ej: "planeada", "en germinacion", "plantada", etc.
	Position      string         `json:"position,omitempty"`                          // Valor o descripción para la sugerencia textual
	Notes         string         `json:"notes,omitempty" gorm:"type:text"`
	PlantedAt     *time.Time     `json:"planted_at,omitempty"` // Fecha de plantación
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Arrangement Arrangement `json:"arrangement,omitempty" gorm:"foreignKey:ArrangementID"`
	Plant       Plant       `json:"plant,omitempty" gorm:"foreignKey:PlantID"`
}

// Validate valida los datos de una plantación
func (p *Planting) Validate() error {
	if p.ArrangementID <= 0 {
		return errors.New("el lecho es requerido")
	}

	if p.PlantID <= 0 {
		return errors.New("la planta es requerida")
	}

	if p.Quantity <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	if !IsValidStatus(p.Status) {
		return errors.New("estado inválido")
	}

	return nil
}

// CalculateDensity calcula la densidad de plantación (plantas por m²)
func (p *Planting) CalculateDensity(arrangementArea float64) float64 {
	if arrangementArea <= 0 {
		return 0
	}
	return float64(p.Quantity) / arrangementArea
}

// Estructuras para respuestas de API
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
	Error      string      `json:"error,omitempty"`
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// Estructuras para requests
type CreatePlantRequest struct {
	Name            string `json:"name" binding:"required"`
	Scientific      string `json:"scientific"`
	Stratum         string `json:"stratum"`
	Function        string `json:"function"`
	SuccessionStage string `json:"succession_stage"`
	ExternalID      string `json:"external_id"`
	Desired         bool   `json:"desired"`
	Notes           string `json:"notes"`
}

type UpdatePlantRequest struct {
	Name            *string `json:"name"`
	Scientific      *string `json:"scientific"`
	Stratum         *string `json:"stratum"`
	Function        *string `json:"function"`
	SuccessionStage *string `json:"succession_stage"`
	ExternalID      *string `json:"external_id"`
	Desired         *bool   `json:"desired"`
	Notes           *string `json:"notes"`
}

type CreateLocationRequest struct {
	Name  string `json:"name" binding:"required"`
	Notes string `json:"notes"`
}

type CreateArrangementRequest struct {
	Name         string  `json:"name" binding:"required"`
	LocationID   uint    `json:"location_id" binding:"required"`
	Type         string  `json:"type" binding:"required"`
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
	Diameter     float64 `json:"diameter"`
	SoilType     string  `json:"soil_type"`
	PlantingMode string  `json:"planting_mode"`
	Notes        string  `json:"notes"`
}

type CreatePlantingRequest struct {
	ArrangementID uint   `json:"arrangement_id" binding:"required"`
	PlantID       uint   `json:"plant_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required,min=1"`
	Status        string `json:"status" binding:"required"`
	Position      string `json:"position"`
	Notes         string `json:"notes"`
}
