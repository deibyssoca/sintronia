package repositories

import (
	"errors"
	"fmt"

	"github.com/deibys/sintronia/internal/db"
	"github.com/deibys/sintronia/pkg/models"
	"gorm.io/gorm"
)

type PlantRepository struct {
	db *gorm.DB
}

func NewPlantRepository() *PlantRepository {

	// Verificar que la conexión DB esté inicializada
	if db.DB == nil {
		panic("Base de datos no inicializada. Asegúrate de llamar db.InitDatabase() antes de crear repositorios")
	}

	return &PlantRepository{
		db: db.DB,
	}
}

// Create crea una nueva planta
func (r *PlantRepository) Create(plant *models.PlantSpecies) error {

	if plant == nil {
		panic("*************    PROBANDO ERROR VIENE NILL plant    ************")
	}

	if r.db == nil {
		panic("*************    PROBANDO ERROR VIENE NILL r.db    ************")
	}

	if err := r.db.Create(plant).Error; err != nil {
		return fmt.Errorf("error creando planta: %w", err)
	}
	return nil
}

// GetAll obtiene todas las plantas con filtros opcionales
func (r *PlantRepository) GetAll(filters PlantFilters) ([]models.PlantSpecies, int64, error) {
	var plants []models.PlantSpecies
	var total int64

	query := r.db.Model(&models.PlantSpecies{})

	// Aplicar filtros
	if filters.Search != "" {
		searchTerm := "%" + filters.Search + "%"
		query = query.Where("common_name ILIKE ? OR scientific_name ILIKE ?", searchTerm, searchTerm)
	}

	if filters.Stratum != "" {
		query = query.Where("stratum = ?", filters.Stratum)
	}

	if filters.FunctionEcol != "" {
		query = query.Where("function_ecol = ?", filters.FunctionEcol)
	}

	if filters.SuccessionStage != "" {
		query = query.Where("succession_stage = ?", filters.SuccessionStage)
	}

	// Contar total antes de paginación
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error contando plantas: %w", err)
	}

	// Aplicar paginación
	if filters.Limit > 0 {
		query = query.Limit(filters.Limit)
	}

	if filters.Offset > 0 {
		query = query.Offset(filters.Offset)
	}

	// Ordenar por fecha de creación (más recientes primero)
	query = query.Order("created_at DESC")

	// Ejecutar consulta
	if err := query.Find(&plants).Error; err != nil {
		return nil, 0, fmt.Errorf("error obteniendo plantas: %w", err)
	}

	return plants, total, nil
}

// GetByID obtiene una planta por ID
func (r *PlantRepository) GetByID(id uint) (*models.PlantSpecies, error) {
	var plant models.PlantSpecies

	if err := r.db.First(&plant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("planta no encontrada")
		}
		return nil, fmt.Errorf("error obteniendo planta: %w", err)
	}

	return &plant, nil
}

// Update actualiza una planta
func (r *PlantRepository) Update(id uint, updates map[string]interface{}) (*models.PlantSpecies, error) {
	var plant models.PlantSpecies

	// Verificar que la planta existe
	if err := r.db.First(&plant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("planta no encontrada")
		}
		return nil, fmt.Errorf("error obteniendo planta: %w", err)
	}

	// Actualizar campos
	if err := r.db.Model(&plant).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("error actualizando planta: %w", err)
	}

	// Recargar la planta actualizada
	if err := r.db.First(&plant, id).Error; err != nil {
		return nil, fmt.Errorf("error recargando planta: %w", err)
	}

	return &plant, nil
}

// Delete elimina una planta (soft delete)
func (r *PlantRepository) Delete(id uint) error {
	// Verificar que la planta existe
	var plant models.PlantSpecies
	if err := r.db.First(&plant, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("planta no encontrada")
		}
		return fmt.Errorf("error obteniendo planta: %w", err)
	}

	// Verificar que no esté siendo usada en plantaciones
	// var plantingCount int64
	// if err := r.db.Model(&models.Planting{}).Where("plant_id = ?", id).Count(&plantingCount).Error; err != nil {
	// 	return fmt.Errorf("error verificando plantaciones: %w", err)
	// }

	// if plantingCount > 0 {
	// 	return fmt.Errorf("no se puede eliminar la planta porque está siendo usada en %d plantaciones", plantingCount)
	// }

	// Soft delete
	if err := r.db.Delete(&plant).Error; err != nil {
		return fmt.Errorf("error eliminando planta: %w", err)
	}

	return nil
}

// ExistsByExternalRef verifica si existe una planta con el external_id dado
func (r *PlantRepository) ExistsByExternalRef(externalRef string) (bool, error) {
	if externalRef == "" {
		return false, nil
	}

	var count int64
	if err := r.db.Model(&models.PlantSpecies{}).Where("external_ref = ?", externalRef).Count(&count).Error; err != nil {
		return false, fmt.Errorf("error verificando external_ref: %w", err)
	}

	return count > 0, nil
}

// PlantFilters estructura para filtros de búsqueda
type PlantFilters struct {
	Search          string
	Stratum         string
	FunctionEcol    string
	SuccessionStage string
	Limit           int
	Offset          int
}
