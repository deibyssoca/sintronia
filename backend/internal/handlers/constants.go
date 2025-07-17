package handlers

import (
	"net/http"

	"github.com/deibys/sintronia/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetConstantsHandler devuelve todas las constantes disponibles
func GetConstantsHandler(c *gin.Context) {
	constants := map[string]interface{}{
		"estratos": []string{
			models.StratumEmergent,
			models.StratumHigh,
			models.StratumMedium,
			models.StratumLow,
			models.StratumGround,
			models.StratumClimber,
			models.StratumRoot,
		},
		"etapas_sucesionales": []string{
			models.SuccessionPlacenta,
			models.SuccessionPioneer,
			models.SuccessionSecondary,
			models.SuccessionClimax,
		},
		"funciones": []string{
			models.FunctionNitrogenFixer,
			models.FunctionDynamicAccumulator,
			models.FunctionGroundCover,
			models.FunctionWindbreak,
			models.FunctionPollinator,
			models.FunctionPestControl,
			models.FunctionSoilAeration,
			models.FunctionWaterRegulation,
			models.FunctionBiomassProduction,
			models.FunctionFood,
			models.FunctionMedicinal,
			models.FunctionTimber,
			models.FunctionFiber,
			models.FunctionOrnamental,
		},
		"modalidades_plantacion": []string{
			models.PlantingModeSeed,
			models.PlantingModeCutting,
			models.PlantingModeStake,
			models.PlantingModeSeedling,
			models.PlantingModeTree,
		},
		"estados": []string{
			models.StatusPlanned,
			models.StatusGerminating,
			models.StatusSeedling,
			models.StatusPlanted,
			models.StatusEstablished,
			models.StatusProductive,
			models.StatusDormant,
			models.StatusDead,
		},
		"tipos_lecho": []string{
			models.ArrangementTypeLine,
			models.ArrangementTypeIsland,
			models.ArrangementTypeGuild,
		},
		"tipos_suelo": []string{
			models.SoilTypeArgiloso,
			models.SoilTypeArenoso,
			models.SoilTypeFranco,
			models.SoilTypeHumifero,
			models.SoilTypePedregoso,
			models.SoilTypeAnegadizo,
		},
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    constants,
		Message: "Constantes del sistema",
	})
}
