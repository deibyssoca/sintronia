package models

// Estratos de vegetación en agricultura sintrópica
const (
	StratumEmergent = "emergente" // Árboles de gran porte (>25m)
	StratumHigh     = "alto"      // Árboles medianos (15-25m)
	StratumMedium   = "medio"     // Árboles pequeños y arbustos (5-15m)
	StratumLow      = "bajo"      // Arbustos y herbáceas (1-5m)
	StratumGround   = "rastrero"  // Cobertura del suelo (<1m)
	StratumClimber  = "trepador"  // Plantas trepadoras
	StratumRoot     = "raiz"      // Sistema radicular/tubérculos
)

// Etapas sucesionales según Ernst Götsch
const (
	SuccessionPlacenta  = "placenta"   // Preparación del suelo
	SuccessionPioneer   = "pionera"    // Colonización inicial
	SuccessionSecondary = "secundaria" // Consolidación
	SuccessionClimax    = "climax"     // Clímax/madurez
)

// Funciones ecológicas en sistemas sintrópicos
const (
	FunctionNitrogenFixer      = "fijador_nitrogeno"   // Leguminosas
	FunctionDynamicAccumulator = "acumulador_dinamico" // Acumulan minerales
	FunctionGroundCover        = "cobertura_suelo"     // Protección del suelo
	FunctionWindbreak          = "cortaviento"         // Protección contra viento
	FunctionPollinator         = "polinizador"         // Atrae polinizadores
	FunctionPestControl        = "control_plagas"      // Control biológico
	FunctionSoilAeration       = "aireacion_suelo"     // Mejora estructura del suelo
	FunctionWaterRegulation    = "regulacion_agua"     // Manejo hídrico
	FunctionBiomassProduction  = "produccion_biomasa"  // Generación de materia orgánica
	FunctionFood               = "alimentario"         // Producción de alimentos
	FunctionMedicinal          = "medicinal"           // Propiedades medicinales
	FunctionTimber             = "maderable"           // Producción de madera
	FunctionFiber              = "fibra"               // Producción de fibras
	FunctionOrnamental         = "ornamental"          // Valor estético
)

// Modalidades de plantación
const (
	PlantingModeSeed     = "semilla" // Siembra directa
	PlantingModeCutting  = "esqueje" // Propagación vegetativa
	PlantingModeStake    = "estaca"  // Estacas leñosas
	PlantingModeSeedling = "planta"  // Plantines/mudas
	PlantingModeTree     = "arbol"   // Árboles desarrollados
)

// Estados del ciclo de plantación
const (
	StatusPlanned     = "planeada"    // En planificación
	StatusGerminating = "germinacion" // En proceso de germinación
	StatusSeedling    = "plantula"    // Estado de plántula
	StatusPlanted     = "plantada"    // Plantada en campo
	StatusEstablished = "establecida" // Establecida y creciendo
	StatusProductive  = "productiva"  // En etapa productiva
	StatusDormant     = "dormante"    // En dormancia
	StatusDead        = "muerta"      // No viable
)

// Tipos de lecho de cultivo
const (
	ArrangementTypeLine   = "linea"  // Líneas de plantación
	ArrangementTypeIsland = "isla"   // Islas circulares
	ArrangementTypeGuild  = "gremio" // Gremios de plantas
)

// Tipos de suelo
const (
	SoilTypeArgiloso  = "argiloso"  // Arcilloso
	SoilTypeArenoso   = "arenoso"   // Arenoso
	SoilTypeFranco    = "franco"    // Franco
	SoilTypeHumifero  = "humifero"  // Rico en humus
	SoilTypePedregoso = "pedregoso" // Pedregoso
	SoilTypeAnegadizo = "anegadizo" // Propenso a encharcamiento
)

// Funciones de validación
func IsValidStratum(stratum string) bool {
	validStrata := []string{
		StratumEmergent, StratumHigh, StratumMedium,
		StratumLow, StratumGround, StratumClimber, StratumRoot,
	}
	for _, v := range validStrata {
		if v == stratum {
			return true
		}
	}
	return false
}

// Etapas sucesionales según Ernst Götsch
func IsValidSuccessionStage(stage string) bool {
	validStages := []string{
		SuccessionPlacenta, SuccessionPioneer,
		SuccessionSecondary, SuccessionClimax,
	}
	for _, v := range validStages {
		if v == stage {
			return true
		}
	}
	return false
}

func IsValidFunction(function string) bool {
	validFunctions := []string{
		FunctionNitrogenFixer, FunctionDynamicAccumulator, FunctionGroundCover,
		FunctionWindbreak, FunctionPollinator, FunctionPestControl,
		FunctionSoilAeration, FunctionWaterRegulation, FunctionBiomassProduction,
		FunctionFood, FunctionMedicinal, FunctionTimber, FunctionFiber, FunctionOrnamental,
	}
	for _, v := range validFunctions {
		if v == function {
			return true
		}
	}
	return false
}

func IsValidPlantingMode(mode string) bool {
	validModes := []string{
		PlantingModeSeed, PlantingModeCutting,
		PlantingModeStake, PlantingModeSeedling, PlantingModeTree,
	}
	for _, v := range validModes {
		if v == mode {
			return true
		}
	}
	return false
}

func IsValidStatus(status string) bool {
	validStatuses := []string{
		StatusPlanned, StatusGerminating, StatusSeedling,
		StatusPlanted, StatusEstablished, StatusProductive,
		StatusDormant, StatusDead,
	}
	for _, v := range validStatuses {
		if v == status {
			return true
		}
	}
	return false
}

func IsValidArrangementType(arrangementType string) bool {
	validTypes := []string{ArrangementTypeLine, ArrangementTypeIsland, ArrangementTypeGuild}
	for _, v := range validTypes {
		if v == arrangementType {
			return true
		}
	}
	return false
}

func IsValidSoilType(soilType string) bool {
	validTypes := []string{
		SoilTypeArgiloso, SoilTypeArenoso, SoilTypeFranco,
		SoilTypeHumifero, SoilTypePedregoso, SoilTypeAnegadizo,
	}
	for _, v := range validTypes {
		if v == soilType {
			return true
		}
	}
	return false
}
