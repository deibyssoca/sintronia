package models

// Tipos de parcela
const (
	PlotTypeLine   = "line"   // Líneas de plantación
	PlotTypeIsland = "island" // Islas circulares
	PlotTypeGuild  = "guild"  // Gremios de plantas
)

// Roles de plantas en el sistema sintrópico
const (
	PlantRoleObjetivo    = "objetivo"    // Planta objetivo (producción principal)
	PlantRoleServicio    = "servicio"    // Planta de servicio (apoyo ecológico)
	PlantRoleAcompañante = "acompañante" // Planta acompañante
)

// Estados de las instancias de plantas (simplificados y en inglés)
const (
	PlantStatusPlanned     = "planned"     // Planificada
	PlantStatusGerminated  = "germinated"  // Germinada
	PlantStatusPlanted     = "planted"     // Plantada
	PlantStatusEstablished = "established" // Establecida
	PlantStatusProductive  = "productive"  // Productiva
	PlantStatusDormant     = "dormant"     // Dormante
	PlantStatusDead        = "dead"        // Muerta
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

// Tipos de suelo
const (
	SoilTypeArgiloso  = "argiloso"  // Arcilloso
	SoilTypeArenoso   = "arenoso"   // Arenoso
	SoilTypeFranco    = "franco"    // Franco
	SoilTypeHumifero  = "humifero"  // Rico en humus
	SoilTypePedregoso = "pedregoso" // Pedregoso
	SoilTypeAnegadizo = "anegadizo" // Propenso a encharcamiento
)

// Funciones de validación para el nuevo modelo

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

func IsValidPlotType(plotType string) bool {
	validTypes := []string{PlotTypeLine, PlotTypeIsland, PlotTypeGuild}
	for _, v := range validTypes {
		if v == plotType {
			return true
		}
	}
	return false
}

func IsValidPlantRole(role string) bool {
	validRoles := []string{PlantRoleObjetivo, PlantRoleServicio, PlantRoleAcompañante}
	for _, v := range validRoles {
		if v == role {
			return true
		}
	}
	return false
}

func IsValidPlantStatus(status string) bool {
	validStatuses := []string{
		PlantStatusPlanned, PlantStatusGerminated, PlantStatusPlanted,
		PlantStatusEstablished, PlantStatusProductive, PlantStatusDormant, PlantStatusDead,
	}
	for _, v := range validStatuses {
		if v == status {
			return true
		}
	}
	return false
}

// Mantener funciones de validación existentes para compatibilidad
// (las funciones IsValidStratum, IsValidFunction, etc. se mantienen igual)
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
