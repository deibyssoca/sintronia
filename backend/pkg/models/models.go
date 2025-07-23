package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Site representa un sitio o terreno principal
type Site struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"-:migration"`
	AreaM2    float64        `json:"area_m2" gorm:"type:decimal(12,2)"` // Área total calculada
	LengthM   float64        `json:"length_m" gorm:"type:decimal(10,2)"`
	WidthM    float64        `json:"width_m" gorm:"type:decimal(10,2)"`
	Notes     string         `json:"notes" gorm:"type:text"`
	Climate   string         `json:"climate" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Plantations []Plantation `json:"plantations,omitempty" gorm:"foreignKey:SiteID"`
}

// Validate valida los datos de un sitio
func (s *Site) Validate() error {
	if strings.TrimSpace(s.Name) == "" {
		return errors.New("el nombre del sitio es requerido")
	}

	if s.AreaM2 < 0 {
		return errors.New("el área no puede ser negativa")
	}

	if s.LengthM < 0 || s.WidthM < 0 {
		return errors.New("las dimensiones no pueden ser negativas")
	}

	return nil
}

// CalculateArea calcula el área basada en longitud y ancho si no está definida
func (s *Site) CalculateArea() float64 {
	if s.AreaM2 > 0 {
		return s.AreaM2
	}
	if s.LengthM > 0 && s.WidthM > 0 {
		return s.LengthM * s.WidthM
	}
	return 0
}

// Plantation representa una plantación o zona de cultivo dentro de un sitio
type Plantation struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	SiteID    uint           `json:"site_id" gorm:"not null;index"`
	Name      string         `json:"name" gorm:"not null;-:migration"`
	AreaM2    float64        `json:"area_m2" gorm:"type:decimal(12,2)"` // Área definida o calculada
	Notes     string         `json:"notes" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Site                Site                 `json:"site,omitempty" gorm:"foreignKey:SiteID"`
	Plots               []Plot               `json:"plots,omitempty" gorm:"foreignKey:PlantationID"`
	SuggestionTemplates []SuggestionTemplate `json:"suggestion_templates,omitempty" gorm:"foreignKey:PlantationID"`
}

// Validate valida los datos de una plantación
func (p *Plantation) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("el nombre de la plantación es requerido")
	}

	if p.SiteID <= 0 {
		return errors.New("el sitio es requerido")
	}

	if p.AreaM2 < 0 {
		return errors.New("el área no puede ser negativa")
	}

	return nil
}

// PlantSpecies representa una especie de planta en el catálogo
type PlantSpecies struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CommonName      string         `json:"common_name" gorm:"not null;index;-:migration"`
	ScientificName  string         `json:"scientific_name" gorm:"index;-:migration"`
	Stratum         string         `json:"stratum" gorm:"type:varchar(50);index;-:migration"`             // Ej: "bajo", "medio", "alto"
	FunctionEcol    string         `json:"function_ecol" gorm:"type:varchar(100);index;-:migration"`      // "objetivo" o "servicio"
	SuccessionStage string         `json:"succession_stage" gorm:"type:varchar(50);index;-:migration"`    // Ej: "pionera", "secundaria", "climax"
	ExternalRef     string         `json:"external_ref" gorm:"type:varchar(100);uniqueIndex;-:migration"` // Referencia a la API externa
	Notes           string         `json:"notes" gorm:"type:text"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	PlantInstances []PlantInstance `json:"plant_instances,omitempty" gorm:"foreignKey:SpeciesID"`
}

// Validate valida los datos de una especie de planta
func (ps *PlantSpecies) Validate() error {
	if strings.TrimSpace(ps.CommonName) == "" {
		return errors.New("el nombre común de la especie es requerido")
	}

	if ps.Stratum != "" && !IsValidStratum(ps.Stratum) {
		return errors.New("estrato inválido")
	}

	if ps.FunctionEcol != "" && !IsValidFunction(ps.FunctionEcol) {
		return errors.New("función ecológica inválida")
	}

	if ps.SuccessionStage != "" && !IsValidSuccessionStage(ps.SuccessionStage) {
		return errors.New("etapa sucesional inválida")
	}

	return nil
}

// Plot representa una parcela o lecho de cultivo
type Plot struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	PlantationID uint           `json:"plantation_id" gorm:"not null;index"`
	PlotType     string         `json:"plot_type" gorm:"type:varchar(50);not null;index;-:migration"` // "line", "island", "guild"
	LengthM      float64        `json:"length_m" gorm:"type:decimal(10,2)"`                           // Solo para líneas
	WidthM       float64        `json:"width_m" gorm:"type:decimal(10,2)"`                            // Solo para líneas
	DiameterM    float64        `json:"diameter_m" gorm:"type:decimal(10,2)"`                         // Solo para islas
	Geometry     string         `json:"geometry" gorm:"type:text"`                                    // GeoJSON opcional
	Notes        string         `json:"notes" gorm:"type:text"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Plantation     Plantation      `json:"plantation,omitempty" gorm:"foreignKey:PlantationID"`
	PlantInstances []PlantInstance `json:"plant_instances,omitempty" gorm:"foreignKey:PlotID"`
}

// Validate valida los datos de una parcela
func (p *Plot) Validate() error {
	if p.PlantationID <= 0 {
		return errors.New("la plantación es requerida")
	}

	if !IsValidPlotType(p.PlotType) {
		return errors.New("tipo de parcela inválido")
	}

	// Validaciones específicas por tipo
	switch p.PlotType {
	case PlotTypeLine:
		if p.LengthM <= 0 || p.WidthM <= 0 {
			return errors.New("las líneas requieren longitud y ancho válidos")
		}
	case PlotTypeIsland:
		if p.DiameterM <= 0 {
			return errors.New("las islas requieren un diámetro válido")
		}
	}

	return nil
}

// CalculateArea calcula el área de la parcela en metros cuadrados
func (p *Plot) CalculateArea() float64 {
	switch p.PlotType {
	case PlotTypeLine:
		return p.LengthM * p.WidthM
	case PlotTypeIsland:
		radius := p.DiameterM / 2
		return 3.14159 * radius * radius
	default:
		return 0
	}
}

// PlantInstance representa una instancia específica de plantas en una parcela
type PlantInstance struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	PlotID    uint           `json:"plot_id" gorm:"not null;index"`
	SpeciesID uint           `json:"species_id" gorm:"not null;index"`
	Quantity  int            `json:"quantity" gorm:"not null;check:quantity > 0;-:migration"`
	Role      string         `json:"role" gorm:"type:varchar(50);index;-:migration"`            // "objetivo", "servicio", "acompañante"
	Status    string         `json:"status" gorm:"type:varchar(50);not null;index;-:migration"` // "planned", "germinated", "planted", etc.
	Position  string         `json:"position" gorm:"type:text;-:migration"`                     // GeoJSON o descripción textual
	Order     int            `json:"order" gorm:"not null;-:migration"`
	PlantedAt *time.Time     `json:"planted_at" gorm:"type:date"` // Fecha de plantación
	Notes     string         `json:"notes" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Plot    Plot         `json:"plot,omitempty" gorm:"foreignKey:PlotID"`
	Species PlantSpecies `json:"species,omitempty" gorm:"foreignKey:SpeciesID"`
}

// Validate valida los datos de una instancia de planta
func (pi *PlantInstance) Validate() error {
	if pi.PlotID <= 0 {
		return errors.New("la parcela es requerida")
	}

	if pi.SpeciesID <= 0 {
		return errors.New("la especie es requerida")
	}

	if pi.Quantity <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}

	if pi.Role != "" && !IsValidPlantRole(pi.Role) {
		return errors.New("rol de planta inválido")
	}

	if !IsValidPlantStatus(pi.Status) {
		return errors.New("estado inválido")
	}

	return nil
}

// CalculateDensity calcula la densidad de plantación (plantas por m²)
func (pi *PlantInstance) CalculateDensity(plotArea float64) float64 {
	if plotArea <= 0 {
		return 0
	}
	return float64(pi.Quantity) / plotArea
}

// SuggestionTemplate representa una plantilla de sugerencias para plantaciones
type SuggestionTemplate struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	PlantationID uint           `json:"plantation_id" gorm:"not null;index"`
	Name         string         `json:"name" gorm:"not null"`
	Description  string         `json:"description" gorm:"type:text"`
	Rules        string         `json:"rules" gorm:"type:jsonb"` // Reglas JSON para densidad, estrato, sucesión
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete

	// Relaciones
	Plantation Plantation `json:"plantation,omitempty" gorm:"foreignKey:PlantationID"`
}

// Validate valida los datos de una plantilla de sugerencias
func (st *SuggestionTemplate) Validate() error {
	if strings.TrimSpace(st.Name) == "" {
		return errors.New("el nombre de la plantilla es requerido")
	}

	if st.PlantationID <= 0 {
		return errors.New("la plantación es requerida")
	}

	return nil
}

// Estructuras para respuestas de API (mantenemos compatibilidad)
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

// Estructuras para requests del nuevo modelo
type CreateSiteRequest struct {
	Name    string  `json:"name" binding:"required"`
	AreaM2  float64 `json:"area_m2"`
	LengthM float64 `json:"length_m"`
	WidthM  float64 `json:"width_m"`
	Notes   string  `json:"notes"`
}

type UpdateSiteRequest struct {
	Name    *string  `json:"name"`
	AreaM2  *float64 `json:"area_m2"`
	LengthM *float64 `json:"length_m"`
	WidthM  *float64 `json:"width_m"`
	Notes   *string  `json:"notes"`
}

type CreatePlantationRequest struct {
	SiteID uint    `json:"site_id" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	AreaM2 float64 `json:"area_m2"`
	Notes  string  `json:"notes"`
}

type CreatePlantSpeciesRequest struct {
	CommonName      string `json:"common_name" binding:"required"`
	ScientificName  string `json:"scientific_name"`
	Stratum         string `json:"stratum"`
	FunctionEcol    string `json:"function_ecol"`
	SuccessionStage string `json:"succession_stage"`
	ExternalRef     string `json:"external_ref"`
	Notes           string `json:"notes"`
}

type UpdatePlantSpeciesRequest struct {
	CommonName      *string `json:"common_name"`
	ScientificName  *string `json:"scientific_name"`
	Stratum         *string `json:"stratum"`
	FunctionEcol    *string `json:"function_ecol"`
	SuccessionStage *string `json:"succession_stage"`
	ExternalRef     *string `json:"external_ref"`
	Notes           *string `json:"notes"`
}

type CreatePlotRequest struct {
	PlantationID uint    `json:"plantation_id" binding:"required"`
	PlotType     string  `json:"plot_type" binding:"required"`
	LengthM      float64 `json:"length_m"`
	WidthM       float64 `json:"width_m"`
	DiameterM    float64 `json:"diameter_m"`
	Geometry     string  `json:"geometry"`
	Notes        string  `json:"notes"`
}

type CreatePlantInstanceRequest struct {
	PlotID    uint   `json:"plot_id" binding:"required"`
	SpeciesID uint   `json:"species_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	Role      string `json:"role"`
	Status    string `json:"status" binding:"required"`
	Position  string `json:"position"`
	Notes     string `json:"notes"`
}

type CreateSuggestionTemplateRequest struct {
	PlantationID uint   `json:"plantation_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Rules        string `json:"rules"` // JSON string
}
