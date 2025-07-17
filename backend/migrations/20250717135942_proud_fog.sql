-- 🌱 Esquema Inicial - Sistema Sintropia
-- PostgreSQL Schema para Agricultura Sintrópica
-- Backend Migration 001

-- ============================================================================
-- CONFIGURACIÓN INICIAL
-- ============================================================================

-- Configurar zona horaria
SET timezone = 'UTC';

-- Habilitar extensiones necesarias
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================================================
-- TABLA: locations (Ubicaciones/Zonas del terreno)
-- ============================================================================
CREATE TABLE IF NOT EXISTS locations (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para locations
CREATE INDEX IF NOT EXISTS idx_locations_name ON locations(name);
CREATE INDEX IF NOT EXISTS idx_locations_deleted_at ON locations(deleted_at);

-- Comentarios
COMMENT ON TABLE locations IS 'Ubicaciones o zonas del terreno para organización territorial';
COMMENT ON COLUMN locations.name IS 'Nombre único de la ubicación (ej: "Zona Norte", "Huerta Principal")';
COMMENT ON COLUMN locations.notes IS 'Notas adicionales sobre la ubicación';

-- ============================================================================
-- TABLA: plants (Catálogo de plantas)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plants (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    scientific VARCHAR(255),
    stratum VARCHAR(50) CHECK (stratum IN ('emergente', 'alto', 'medio', 'bajo', 'rastrero', 'trepador', 'raiz')),
    function VARCHAR(100) CHECK (function IN (
        'fijador_nitrogeno', 'acumulador_dinamico', 'cobertura_suelo', 'cortaviento',
        'polinizador', 'control_plagas', 'aireacion_suelo', 'regulacion_agua',
        'produccion_biomasa', 'alimentario', 'medicinal', 'maderable', 'fibra', 'ornamental'
    )),
    succession_stage VARCHAR(50) CHECK (succession_stage IN ('placenta', 'pionera', 'secundaria', 'climax')),
    external_id VARCHAR(100) UNIQUE, -- Referencia a API externa (ej: Permapeople)
    desired BOOLEAN DEFAULT FALSE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plants
CREATE INDEX IF NOT EXISTS idx_plants_name ON plants(name);
CREATE INDEX IF NOT EXISTS idx_plants_scientific ON plants(scientific);
CREATE INDEX IF NOT EXISTS idx_plants_stratum ON plants(stratum);
CREATE INDEX IF NOT EXISTS idx_plants_function ON plants(function);
CREATE INDEX IF NOT EXISTS idx_plants_succession_stage ON plants(succession_stage);
CREATE INDEX IF NOT EXISTS idx_plants_external_id ON plants(external_id);
CREATE INDEX IF NOT EXISTS idx_plants_desired ON plants(desired);
CREATE INDEX IF NOT EXISTS idx_plants_deleted_at ON plants(deleted_at);

-- Índice compuesto para búsquedas frecuentes
CREATE INDEX IF NOT EXISTS idx_plants_search ON plants(name, scientific, stratum, desired);

-- Comentarios
COMMENT ON TABLE plants IS 'Catálogo de especies vegetales para agricultura sintrópica';
COMMENT ON COLUMN plants.name IS 'Nombre común de la planta';
COMMENT ON COLUMN plants.scientific IS 'Nombre científico de la planta';
COMMENT ON COLUMN plants.stratum IS 'Estrato de vegetación según Ernst Götsch';
COMMENT ON COLUMN plants.function IS 'Función ecológica principal de la planta';
COMMENT ON COLUMN plants.succession_stage IS 'Etapa sucesional según Ernst Götsch';
COMMENT ON COLUMN plants.external_id IS 'ID de referencia en API externa (ej: Permapeople)';
COMMENT ON COLUMN plants.desired IS 'Indica si está en la lista de plantas deseadas';

-- ============================================================================
-- TABLA: arrangements (Lechos de cultivo)
-- ============================================================================
CREATE TABLE IF NOT EXISTS arrangements (
    id BIGSERIAL PRIMARY KEY,
    location_id BIGINT NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('linea', 'isla', 'gremio')),
    length DECIMAL(10,2), -- Solo para líneas
    width DECIMAL(10,2),  -- Solo para líneas
    diameter DECIMAL(10,2), -- Solo para islas
    soil_type VARCHAR(50) CHECK (soil_type IN ('argiloso', 'arenoso', 'franco', 'humifero', 'pedregoso', 'anegadizo')),
    planting_mode VARCHAR(50) CHECK (planting_mode IN ('semilla', 'esqueje', 'estaca', 'planta', 'arbol')),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para arrangements
CREATE INDEX IF NOT EXISTS idx_arrangements_location_id ON arrangements(location_id);
CREATE INDEX IF NOT EXISTS idx_arrangements_type ON arrangements(type);
CREATE INDEX IF NOT EXISTS idx_arrangements_soil_type ON arrangements(soil_type);
CREATE INDEX IF NOT EXISTS idx_arrangements_deleted_at ON arrangements(deleted_at);

-- Comentarios
COMMENT ON TABLE arrangements IS 'Lechos de cultivo con diferentes disposiciones (líneas, islas, gremios)';
COMMENT ON COLUMN arrangements.type IS 'Tipo de disposición: linea, isla, gremio';
COMMENT ON COLUMN arrangements.length IS 'Longitud en metros (solo para líneas)';
COMMENT ON COLUMN arrangements.width IS 'Ancho en metros (solo para líneas)';
COMMENT ON COLUMN arrangements.diameter IS 'Diámetro en metros (solo para islas)';
COMMENT ON COLUMN arrangements.soil_type IS 'Tipo de suelo del lecho';
COMMENT ON COLUMN arrangements.planting_mode IS 'Modalidad de plantación utilizada';

-- ============================================================================
-- TABLA: plantings (Plantaciones específicas)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plantings (
    id BIGSERIAL PRIMARY KEY,
    arrangement_id BIGINT NOT NULL REFERENCES arrangements(id) ON DELETE CASCADE,
    plant_id BIGINT NOT NULL REFERENCES plants(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    status VARCHAR(50) NOT NULL CHECK (status IN (
        'planeada', 'germinacion', 'plantula', 'plantada', 
        'establecida', 'productiva', 'dormante', 'muerta'
    )),
    position VARCHAR(255), -- Descripción de posición en el lecho
    notes TEXT,
    planted_at TIMESTAMP WITH TIME ZONE, -- Fecha de plantación
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plantings
CREATE INDEX IF NOT EXISTS idx_plantings_arrangement_id ON plantings(arrangement_id);
CREATE INDEX IF NOT EXISTS idx_plantings_plant_id ON plantings(plant_id);
CREATE INDEX IF NOT EXISTS idx_plantings_status ON plantings(status);
CREATE INDEX IF NOT EXISTS idx_plantings_planted_at ON plantings(planted_at);
CREATE INDEX IF NOT EXISTS idx_plantings_deleted_at ON plantings(deleted_at);

-- Índice compuesto para consultas frecuentes
CREATE INDEX IF NOT EXISTS idx_plantings_arrangement_plant ON plantings(arrangement_id, plant_id);

-- Comentarios
COMMENT ON TABLE plantings IS 'Instancias específicas de plantas en lechos de cultivo';
COMMENT ON COLUMN plantings.quantity IS 'Cantidad de plantas de esta especie en el lecho';
COMMENT ON COLUMN plantings.status IS 'Estado actual de la plantación en su ciclo de vida';
COMMENT ON COLUMN plantings.position IS 'Descripción textual de la posición en el lecho';
COMMENT ON COLUMN plantings.planted_at IS 'Fecha y hora de plantación en campo';

-- ============================================================================
-- VISTAS ÚTILES
-- ============================================================================

-- Vista: Plantaciones con información completa
CREATE OR REPLACE VIEW v_plantings_full AS
SELECT 
    p.id,
    p.quantity,
    p.status,
    p.position,
    p.planted_at,
    p.created_at,
    
    -- Información de la planta
    pl.name as plant_name,
    pl.scientific as plant_scientific,
    pl.stratum as plant_stratum,
    pl.function as plant_function,
    pl.succession_stage as plant_succession,
    
    -- Información del lecho
    a.name as arrangement_name,
    a.type as arrangement_type,
    a.length,
    a.width,
    a.diameter,
    
    -- Información de la ubicación
    l.name as location_name,
    
    -- Cálculos
    CASE 
        WHEN a.type = 'linea' AND a.length IS NOT NULL AND a.width IS NOT NULL 
        THEN a.length * a.width
        WHEN a.type = 'isla' AND a.diameter IS NOT NULL 
        THEN PI() * POWER(a.diameter / 2, 2)
        ELSE NULL
    END as area_m2,
    
    CASE 
        WHEN a.type = 'linea' AND a.length > 0 AND a.width > 0 
        THEN p.quantity / (a.length * a.width)
        WHEN a.type = 'isla' AND a.diameter > 0 
        THEN p.quantity / (PI() * POWER(a.diameter / 2, 2))
        ELSE NULL
    END as density_per_m2

FROM plantings p
JOIN plants pl ON p.plant_id = pl.id
JOIN arrangements a ON p.arrangement_id = a.id
JOIN locations l ON a.location_id = l.id
WHERE p.deleted_at IS NULL 
  AND pl.deleted_at IS NULL 
  AND a.deleted_at IS NULL 
  AND l.deleted_at IS NULL;

COMMENT ON VIEW v_plantings_full IS 'Vista completa de plantaciones con información de plantas, lechos y ubicaciones';

-- Vista: Resumen por ubicación
CREATE OR REPLACE VIEW v_location_summary AS
SELECT 
    l.id,
    l.name,
    COUNT(DISTINCT a.id) as total_arrangements,
    COUNT(DISTINCT p.id) as total_plantings,
    COUNT(DISTINCT p.plant_id) as unique_plants,
    COALESCE(SUM(p.quantity), 0) as total_plant_count
FROM locations l
LEFT JOIN arrangements a ON l.id = a.location_id AND a.deleted_at IS NULL
LEFT JOIN plantings p ON a.id = p.arrangement_id AND p.deleted_at IS NULL
WHERE l.deleted_at IS NULL
GROUP BY l.id, l.name;

COMMENT ON VIEW v_location_summary IS 'Resumen estadístico por ubicación';

-- ============================================================================
-- FUNCIONES ÚTILES
-- ============================================================================

-- Función para calcular área de un lecho
CREATE OR REPLACE FUNCTION calculate_arrangement_area(
    p_type VARCHAR(50),
    p_length DECIMAL(10,2),
    p_width DECIMAL(10,2),
    p_diameter DECIMAL(10,2)
) RETURNS DECIMAL(10,2) AS $$
BEGIN
    CASE p_type
        WHEN 'linea' THEN
            IF p_length IS NOT NULL AND p_width IS NOT NULL AND p_length > 0 AND p_width > 0 THEN
                RETURN p_length * p_width;
            END IF;
        WHEN 'isla' THEN
            IF p_diameter IS NOT NULL AND p_diameter > 0 THEN
                RETURN PI() * POWER(p_diameter / 2, 2);
            END IF;
        ELSE
            RETURN NULL;
    END CASE;
    
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION calculate_arrangement_area IS 'Calcula el área de un lecho según su tipo y dimensiones';

-- Función para calcular densidad de plantación
CREATE OR REPLACE FUNCTION calculate_planting_density(
    p_arrangement_id BIGINT,
    p_quantity INTEGER
) RETURNS DECIMAL(10,4) AS $$
DECLARE
    arrangement_area DECIMAL(10,2);
BEGIN
    SELECT calculate_arrangement_area(type, length, width, diameter)
    INTO arrangement_area
    FROM arrangements 
    WHERE id = p_arrangement_id AND deleted_at IS NULL;
    
    IF arrangement_area IS NULL OR arrangement_area <= 0 THEN
        RETURN NULL;
    END IF;
    
    RETURN p_quantity / arrangement_area;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION calculate_planting_density IS 'Calcula la densidad de plantación (plantas por m²)';

-- ============================================================================
-- TRIGGERS PARA UPDATED_AT
-- ============================================================================

-- Función para actualizar updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggers para cada tabla
DROP TRIGGER IF EXISTS update_locations_updated_at ON locations;
CREATE TRIGGER update_locations_updated_at 
    BEFORE UPDATE ON locations 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_arrangements_updated_at ON arrangements;
CREATE TRIGGER update_arrangements_updated_at 
    BEFORE UPDATE ON arrangements 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plants_updated_at ON plants;
CREATE TRIGGER update_plants_updated_at 
    BEFORE UPDATE ON plants 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plantings_updated_at ON plantings;
CREATE TRIGGER update_plantings_updated_at 
    BEFORE UPDATE ON plantings 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================================
-- DATOS DE EJEMPLO (OPCIONAL)
-- ============================================================================

-- Insertar ubicaciones de ejemplo
INSERT INTO locations (name, notes) VALUES 
('Zona Norte', 'Área principal con mejor exposición solar'),
('Huerta Central', 'Zona de hortalizas y plantas medicinales'),
('Ladera Este', 'Pendiente con sistema de terrazas')
ON CONFLICT (name) DO NOTHING;

-- Insertar plantas de ejemplo
INSERT INTO plants (name, scientific, stratum, function, succession_stage, desired) VALUES 
('Aguacate Hass', 'Persea americana', 'alto', 'alimentario', 'secundaria', true),
('Frijol Caupí', 'Vigna unguiculata', 'bajo', 'fijador_nitrogeno', 'pionera', true),
('Bambú Guadua', 'Guadua angustifolia', 'alto', 'maderable', 'pionera', false),
('Plátano Dominico', 'Musa acuminata', 'medio', 'alimentario', 'pionera', true),
('Moringa', 'Moringa oleifera', 'medio', 'medicinal', 'pionera', true),
('Leucaena', 'Leucaena leucocephala', 'medio', 'fijador_nitrogeno', 'pionera', false)
ON CONFLICT (external_id) DO NOTHING;

-- Insertar lechos de ejemplo
INSERT INTO arrangements (location_id, name, type, length, width, soil_type) 
SELECT 
    l.id,
    'Línea de Frutales Norte',
    'linea',
    50.0,
    3.0,
    'franco'
FROM locations l 
WHERE l.name = 'Zona Norte'
ON CONFLICT DO NOTHING;

INSERT INTO arrangements (location_id, name, type, diameter, soil_type) 
SELECT 
    l.id,
    'Isla de Leguminosas',
    'isla',
    8.0,
    'humifero'
FROM locations l 
WHERE l.name = 'Huerta Central'
ON CONFLICT DO NOTHING;

-- ============================================================================
-- MENSAJE DE CONFIRMACIÓN
-- ============================================================================

DO $$
BEGIN
    RAISE NOTICE '🌱 Backend Migration 001 completada!';
    RAISE NOTICE '📊 Tablas: locations, plants, arrangements, plantings';
    RAISE NOTICE '🔍 Vistas: v_plantings_full, v_location_summary';
    RAISE NOTICE '⚙️ Funciones: calculate_arrangement_area, calculate_planting_density';
    RAISE NOTICE '🔄 Triggers: updated_at automático en todas las tablas';
    RAISE NOTICE '🌱 Datos de ejemplo insertados';
END $$;