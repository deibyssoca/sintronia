package services

import (
	"errors"

	"github.com/deibys/sintronia/pkg/models"
)

// CreatePlant es un ejemplo de función que se encargaría de
// validar y guardar la planta en la base de datos.
func CreatePlant(plant *models.Plant) error {
	// Aquí insertarías la lógica de validación:
	if plant.Name == "" {
		return errors.New("el nombre de la planta es obligatorio")
	}

	// Lógica de guardado:
	// Por ejemplo, podrías llamar a un repositorio del paquete internal/db:
	// err := db.PlantRepository.Insert(plant)
	// if err != nil {
	//    return err
	// }

	// Por ahora, asumiremos que la operación tuvo éxito:
	return nil
}
