-- 🌱 Nueva Migración - Modelo Actualizado Sintropia
-- PostgreSQL Schema v2.0 para Agricultura Sintrópica
-- Migración 002: Nuevo modelo jerárquico

-- ============================================================================
-- CONFIGURACIÓN INICIAL
-- ============================================================================

-- Configurar zona horaria
SET timezone = 'UTC';

-- Habilitar extensiones necesarias
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
--CREATE EXTENSION IF NOT EXISTS "postgis"; -- Para geometrías GeoJSON

-- ============================================================================
-- TABLA: sites (Sitios/Terrenos)
-- ============================================================================
CREATE TABLE IF NOT EXISTS sites (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    area_m2 DECIMAL(12,2), -- Área total calculada
    length_m DECIMAL(10,2),
    width_m DECIMAL(10,2),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para sites
CREATE INDEX IF NOT EXISTS idx_sites_name ON sites(name);
CREATE INDEX IF NOT EXISTS idx_sites_deleted_at ON sites(deleted_at);

-- Comentarios
COMMENT ON TABLE sites IS 'Sitios o terrenos principales del sistema';
COMMENT ON COLUMN sites.name IS 'Nombre del sitio (ej: "Finca La Esperanza")';
COMMENT ON COLUMN sites.area_m2 IS 'Área total del sitio en metros cuadrados';
COMMENT ON COLUMN sites.length_m IS 'Longitud del sitio en metros';
COMMENT ON COLUMN sites.width_m IS 'Ancho del sitio en metros';

-- ============================================================================
-- TABLA: plantations (Plantaciones/Zonas de cultivo)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plantations (
    id BIGSERIAL PRIMARY KEY,
    site_id BIGINT NOT NULL REFERENCES sites(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    area_m2 DECIMAL(12,2), -- Área definida o calculada
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plantations
CREATE INDEX IF NOT EXISTS idx_plantations_site_id ON plantations(site_id);
CREATE INDEX IF NOT EXISTS idx_plantations_name ON plantations(name);
CREATE INDEX IF NOT EXISTS idx_plantations_deleted_at ON plantations(deleted_at);

-- Comentarios
COMMENT ON TABLE plantations IS 'Plantaciones o zonas de cultivo dentro de un sitio';
COMMENT ON COLUMN plantations.name IS 'Nombre de la plantación (ej: "Zona Norte", "Huerta Central")';
COMMENT ON COLUMN plantations.area_m2 IS 'Área de la plantación en metros cuadrados';

-- ============================================================================
-- TABLA: plant_species (Especies de plantas)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plant_species (
    id BIGSERIAL PRIMARY KEY,
    common_name VARCHAR(255) NOT NULL,
    scientific_name VARCHAR(255),
    stratum VARCHAR(50) CHECK (stratum IN ('emergente', 'alto', 'medio', 'bajo', 'rastrero', 'trepador', 'raiz')),
    function_ecol VARCHAR(100) CHECK (function_ecol IN (
        'fijador_nitrogeno', 'acumulador_dinamico', 'cobertura_suelo', 'cortaviento',
        'polinizador', 'control_plagas', 'aireacion_suelo', 'regulacion_agua',
        'produccion_biomasa', 'alimentario', 'medicinal', 'maderable', 'fibra', 'ornamental'
    )),
    succession_stage VARCHAR(50) CHECK (succession_stage IN ('placenta', 'pionera', 'secundaria', 'climax')),
    external_ref VARCHAR(100) UNIQUE, -- Referencia a API externa (ej: Permapeople)
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plant_species
CREATE INDEX IF NOT EXISTS idx_plant_species_common_name ON plant_species(common_name);
CREATE INDEX IF NOT EXISTS idx_plant_species_scientific_name ON plant_species(scientific_name);
CREATE INDEX IF NOT EXISTS idx_plant_species_stratum ON plant_species(stratum);
CREATE INDEX IF NOT EXISTS idx_plant_species_function_ecol ON plant_species(function_ecol);
CREATE INDEX IF NOT EXISTS idx_plant_species_succession_stage ON plant_species(succession_stage);
CREATE INDEX IF NOT EXISTS idx_plant_species_external_ref ON plant_species(external_ref);
CREATE INDEX IF NOT EXISTS idx_plant_species_deleted_at ON plant_species(deleted_at);

-- Índice compuesto para búsquedas frecuentes
CREATE INDEX IF NOT EXISTS idx_plant_species_search ON plant_species(common_name, scientific_name, stratum);

-- Comentarios
COMMENT ON TABLE plant_species IS 'Catálogo de especies vegetales para agricultura sintrópica';
COMMENT ON COLUMN plant_species.common_name IS 'Nombre común de la especie';
COMMENT ON COLUMN plant_species.scientific_name IS 'Nombre científico de la especie';
COMMENT ON COLUMN plant_species.stratum IS 'Estrato de vegetación según Ernst Götsch';
COMMENT ON COLUMN plant_species.function_ecol IS 'Función ecológica principal de la especie';
COMMENT ON COLUMN plant_species.succession_stage IS 'Etapa sucesional según Ernst Götsch';
COMMENT ON COLUMN plant_species.external_ref IS 'ID de referencia en API externa (ej: Permapeople)';

-- ============================================================================
-- TABLA: plots (Parcelas/Lechos de cultivo)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plots (
    id BIGSERIAL PRIMARY KEY,
    plantation_id BIGINT NOT NULL REFERENCES plantations(id) ON DELETE CASCADE,
    plot_type VARCHAR(50) NOT NULL CHECK (plot_type IN ('line', 'island', 'guild')),
    length_m DECIMAL(10,2), -- Para líneas
    width_m DECIMAL(10,2),  -- Para líneas
    diameter_m DECIMAL(10,2), -- Para islas
    --geometry GEOMETRY(POLYGON, 4326), -- GeoJSON opcional (PostGIS)
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plots
CREATE INDEX IF NOT EXISTS idx_plots_plantation_id ON plots(plantation_id);
CREATE INDEX IF NOT EXISTS idx_plots_plot_type ON plots(plot_type);
CREATE INDEX IF NOT EXISTS idx_plots_deleted_at ON plots(deleted_at);

-- Índice espacial para geometrías
--CREATE INDEX IF NOT EXISTS idx_plots_geometry ON plots USING GIST(geometry);

-- Comentarios
COMMENT ON TABLE plots IS 'Parcelas o lechos de cultivo con diferentes disposiciones';
COMMENT ON COLUMN plots.plot_type IS 'Tipo de parcela: line (línea), island (isla), guild (gremio)';
COMMENT ON COLUMN plots.length_m IS 'Longitud en metros (para líneas)';
COMMENT ON COLUMN plots.width_m IS 'Ancho en metros (para líneas)';
COMMENT ON COLUMN plots.diameter_m IS 'Diámetro en metros (para islas)';
--COMMENT ON COLUMN plots.geometry IS 'Geometría GeoJSON de la parcela (PostGIS)';

-- ============================================================================
-- TABLA: plant_instances (Instancias de plantas)
-- ============================================================================
CREATE TABLE IF NOT EXISTS plant_instances (
    id BIGSERIAL PRIMARY KEY,
    plot_id BIGINT NOT NULL REFERENCES plots(id) ON DELETE CASCADE,
    species_id BIGINT NOT NULL REFERENCES plant_species(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    role VARCHAR(50) CHECK (role IN ('objetivo', 'servicio', 'acompañante')),
    status VARCHAR(50) NOT NULL CHECK (status IN (
        'planned', 'germinated', 'planted', 'established', 
        'productive', 'dormant', 'dead'
    )),
    position TEXT, -- GeoJSON o descripción textual
    planted_at DATE, -- Fecha de plantación
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para plant_instances
CREATE INDEX IF NOT EXISTS idx_plant_instances_plot_id ON plant_instances(plot_id);
CREATE INDEX IF NOT EXISTS idx_plant_instances_species_id ON plant_instances(species_id);
CREATE INDEX IF NOT EXISTS idx_plant_instances_role ON plant_instances(role);
CREATE INDEX IF NOT EXISTS idx_plant_instances_status ON plant_instances(status);
CREATE INDEX IF NOT EXISTS idx_plant_instances_planted_at ON plant_instances(planted_at);
CREATE INDEX IF NOT EXISTS idx_plant_instances_deleted_at ON plant_instances(deleted_at);

-- Índice compuesto para consultas frecuentes
CREATE INDEX IF NOT EXISTS idx_plant_instances_plot_species ON plant_instances(plot_id, species_id);

-- Comentarios
COMMENT ON TABLE plant_instances IS 'Instancias específicas de plantas en parcelas';
COMMENT ON COLUMN plant_instances.quantity IS 'Cantidad de plantas de esta especie';
COMMENT ON COLUMN plant_instances.role IS 'Rol de la planta: objetivo, servicio, acompañante';
COMMENT ON COLUMN plant_instances.status IS 'Estado actual de la instancia de planta';
COMMENT ON COLUMN plant_instances.position IS 'Posición GeoJSON o descripción textual';
COMMENT ON COLUMN plant_instances.planted_at IS 'Fecha de plantación';

-- ============================================================================
-- TABLA: suggestion_templates (Plantillas de sugerencias)
-- ============================================================================
CREATE TABLE IF NOT EXISTS suggestion_templates (
    id BIGSERIAL PRIMARY KEY,
    plantation_id BIGINT NOT NULL REFERENCES plantations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    rules JSONB, -- Reglas de densidad, estrato, sucesión
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete
);

-- Índices para suggestion_templates
CREATE INDEX IF NOT EXISTS idx_suggestion_templates_plantation_id ON suggestion_templates(plantation_id);
CREATE INDEX IF NOT EXISTS idx_suggestion_templates_name ON suggestion_templates(name);
CREATE INDEX IF NOT EXISTS idx_suggestion_templates_deleted_at ON suggestion_templates(deleted_at);

-- Índice GIN para búsquedas en JSON
CREATE INDEX IF NOT EXISTS idx_suggestion_templates_rules ON suggestion_templates USING GIN(rules);

-- Comentarios
COMMENT ON TABLE suggestion_templates IS 'Plantillas de sugerencias para plantaciones';
COMMENT ON COLUMN suggestion_templates.name IS 'Nombre de la plantilla';
COMMENT ON COLUMN suggestion_templates.description IS 'Descripción de la plantilla';
COMMENT ON COLUMN suggestion_templates.rules IS 'Reglas JSON para densidad, estrato, sucesión';

-- ============================================================================
-- VISTAS ÚTILES
-- ============================================================================

-- Vista: Instancias de plantas con información completa
CREATE OR REPLACE VIEW v_plant_instances_full AS
SELECT 
    pi.id,
    pi.quantity,
    pi.role,
    pi.status,
    pi.position,
    pi.planted_at,
    pi.created_at,
    
    -- Información de la especie
    ps.common_name as species_common_name,
    ps.scientific_name as species_scientific_name,
    ps.stratum as species_stratum,
    ps.function_ecol as species_function,
    ps.succession_stage as species_succession,
    
    -- Información de la parcela
    p.plot_type,
    p.length_m,
    p.width_m,
    p.diameter_m,
    
    -- Información de la plantación
    pl.name as plantation_name,
    pl.area_m2 as plantation_area,
    
    -- Información del sitio
    s.name as site_name,
    s.area_m2 as site_area,
    
    -- Cálculos
    CASE 
        WHEN p.plot_type = 'line' AND p.length_m IS NOT NULL AND p.width_m IS NOT NULL 
        THEN p.length_m * p.width_m
        WHEN p.plot_type = 'island' AND p.diameter_m IS NOT NULL 
        THEN PI() * POWER(p.diameter_m / 2, 2)
        ELSE NULL
    END as plot_area_m2,
    
    CASE 
        WHEN p.plot_type = 'line' AND p.length_m > 0 AND p.width_m > 0 
        THEN pi.quantity / (p.length_m * p.width_m)
        WHEN p.plot_type = 'island' AND p.diameter_m > 0 
        THEN pi.quantity / (PI() * POWER(p.diameter_m / 2, 2))
        ELSE NULL
    END as density_per_m2

FROM plant_instances pi
JOIN plant_species ps ON pi.species_id = ps.id
JOIN plots p ON pi.plot_id = p.id
JOIN plantations pl ON p.plantation_id = pl.id
JOIN sites s ON pl.site_id = s.id
WHERE pi.deleted_at IS NULL 
  AND ps.deleted_at IS NULL 
  AND p.deleted_at IS NULL 
  AND pl.deleted_at IS NULL 
  AND s.deleted_at IS NULL;

COMMENT ON VIEW v_plant_instances_full IS 'Vista completa de instancias de plantas con información jerárquica';

-- Vista: Resumen por plantación
CREATE OR REPLACE VIEW v_plantation_summary AS
SELECT 
    pl.id,
    pl.name,
    pl.area_m2,
    s.name as site_name,
    COUNT(DISTINCT p.id) as total_plots,
    COUNT(DISTINCT pi.id) as total_plant_instances,
    COUNT(DISTINCT pi.species_id) as unique_species,
    COALESCE(SUM(pi.quantity), 0) as total_plant_count,
    COUNT(DISTINCT st.id) as suggestion_templates_count
FROM plantations pl
JOIN sites s ON pl.site_id = s.id
LEFT JOIN plots p ON pl.id = p.plantation_id AND p.deleted_at IS NULL
LEFT JOIN plant_instances pi ON p.id = pi.plot_id AND pi.deleted_at IS NULL
LEFT JOIN suggestion_templates st ON pl.id = st.plantation_id AND st.deleted_at IS NULL
WHERE pl.deleted_at IS NULL AND s.deleted_at IS NULL
GROUP BY pl.id, pl.name, pl.area_m2, s.name;

COMMENT ON VIEW v_plantation_summary IS 'Resumen estadístico por plantación';

-- Vista: Especies más utilizadas
CREATE OR REPLACE VIEW v_popular_species AS
SELECT 
    ps.id,
    ps.common_name,
    ps.scientific_name,
    ps.stratum,
    ps.function_ecol,
    ps.succession_stage,
    COUNT(pi.id) as usage_count,
    SUM(pi.quantity) as total_quantity,
    COUNT(DISTINCT pi.plot_id) as plots_used,
    COUNT(DISTINCT p.plantation_id) as plantations_used
FROM plant_species ps
LEFT JOIN plant_instances pi ON ps.id = pi.species_id AND pi.deleted_at IS NULL
LEFT JOIN plots p ON pi.plot_id = p.id AND p.deleted_at IS NULL
WHERE ps.deleted_at IS NULL
GROUP BY ps.id, ps.common_name, ps.scientific_name, ps.stratum, ps.function_ecol, ps.succession_stage
ORDER BY usage_count DESC, total_quantity DESC;

COMMENT ON VIEW v_popular_species IS 'Especies ordenadas por frecuencia de uso';

-- ============================================================================
-- FUNCIONES ÚTILES
-- ============================================================================

-- Función para calcular área de una parcela
CREATE OR REPLACE FUNCTION calculate_plot_area(
    p_plot_type VARCHAR(50),
    p_length_m DECIMAL(10,2),
    p_width_m DECIMAL(10,2),
    p_diameter_m DECIMAL(10,2)
) RETURNS DECIMAL(12,2) AS $$
BEGIN
    CASE p_plot_type
        WHEN 'line' THEN
            IF p_length_m IS NOT NULL AND p_width_m IS NOT NULL AND p_length_m > 0 AND p_width_m > 0 THEN
                RETURN p_length_m * p_width_m;
            END IF;
        WHEN 'island' THEN
            IF p_diameter_m IS NOT NULL AND p_diameter_m > 0 THEN
                RETURN PI() * POWER(p_diameter_m / 2, 2);
            END IF;
        ELSE
            RETURN NULL;
    END CASE;
    
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION calculate_plot_area IS 'Calcula el área de una parcela según su tipo y dimensiones';

-- Función para calcular densidad de plantación
CREATE OR REPLACE FUNCTION calculate_plant_density(
    p_plot_id BIGINT,
    p_quantity INTEGER
) RETURNS DECIMAL(10,4) AS $$
DECLARE
    plot_area DECIMAL(12,2);
BEGIN
    SELECT calculate_plot_area(plot_type, length_m, width_m, diameter_m)
    INTO plot_area
    FROM plots 
    WHERE id = p_plot_id AND deleted_at IS NULL;
    
    IF plot_area IS NULL OR plot_area <= 0 THEN
        RETURN NULL;
    END IF;
    
    RETURN p_quantity / plot_area;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION calculate_plant_density IS 'Calcula la densidad de plantación (plantas por m²)';

-- Función para obtener jerarquía completa de una instancia de planta
CREATE OR REPLACE FUNCTION get_plant_instance_hierarchy(p_instance_id BIGINT)
RETURNS TABLE(
    site_name VARCHAR(255),
    plantation_name VARCHAR(255),
    plot_type VARCHAR(50),
    species_common_name VARCHAR(255),
    quantity INTEGER
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        s.name as site_name,
        pl.name as plantation_name,
        p.plot_type,
        ps.common_name as species_common_name,
        pi.quantity
    FROM plant_instances pi
    JOIN plant_species ps ON pi.species_id = ps.id
    JOIN plots p ON pi.plot_id = p.id
    JOIN plantations pl ON p.plantation_id = pl.id
    JOIN sites s ON pl.site_id = s.id
    WHERE pi.id = p_instance_id
      AND pi.deleted_at IS NULL 
      AND ps.deleted_at IS NULL 
      AND p.deleted_at IS NULL 
      AND pl.deleted_at IS NULL 
      AND s.deleted_at IS NULL;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION get_plant_instance_hierarchy IS 'Obtiene la jerarquía completa de una instancia de planta';

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
DROP TRIGGER IF EXISTS update_sites_updated_at ON sites;
CREATE TRIGGER update_sites_updated_at 
    BEFORE UPDATE ON sites 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plantations_updated_at ON plantations;
CREATE TRIGGER update_plantations_updated_at 
    BEFORE UPDATE ON plantations 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plant_species_updated_at ON plant_species;
CREATE TRIGGER update_plant_species_updated_at 
    BEFORE UPDATE ON plant_species 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plots_updated_at ON plots;
CREATE TRIGGER update_plots_updated_at 
    BEFORE UPDATE ON plots 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_plant_instances_updated_at ON plant_instances;
CREATE TRIGGER update_plant_instances_updated_at 
    BEFORE UPDATE ON plant_instances 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS update_suggestion_templates_updated_at ON suggestion_templates;
CREATE TRIGGER update_suggestion_templates_updated_at 
    BEFORE UPDATE ON suggestion_templates 
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================================
-- DATOS DE EJEMPLO
-- ============================================================================

-- Insertar sitios de ejemplo
INSERT INTO sites (name, area_m2, length_m, width_m, notes) VALUES 
('Finca La Esperanza', 10000.00, 100.0, 100.0, 'Sitio principal de agricultura sintrópica'),
('Parcela Experimental', 2500.00, 50.0, 50.0, 'Área de pruebas y experimentación')
ON CONFLICT DO NOTHING;

-- Insertar plantaciones de ejemplo
INSERT INTO plantations (site_id, name, area_m2, notes) 
SELECT 
    s.id,
    'Zona Norte',
    3000.00,
    'Área principal con mejor exposición solar'
FROM sites s 
WHERE s.name = 'Finca La Esperanza'
ON CONFLICT DO NOTHING;

INSERT INTO plantations (site_id, name, area_m2, notes) 
SELECT 
    s.id,
    'Huerta Central',
    1500.00,
    'Zona de hortalizas y plantas medicinales'
FROM sites s 
WHERE s.name = 'Finca La Esperanza'
ON CONFLICT DO NOTHING;

-- Insertar especies de plantas de ejemplo
INSERT INTO plant_species (common_name, scientific_name, stratum, function_ecol, succession_stage) VALUES 
('Aguacate Hass', 'Persea americana', 'alto', 'alimentario', 'secundaria'),
('Frijol Caupí', 'Vigna unguiculata', 'bajo', 'fijador_nitrogeno', 'pionera'),
('Bambú Guadua', 'Guadua angustifolia', 'alto', 'maderable', 'pionera'),
('Plátano Dominico', 'Musa acuminata', 'medio', 'alimentario', 'pionera'),
('Moringa', 'Moringa oleifera', 'medio', 'medicinal', 'pionera'),
('Leucaena', 'Leucaena leucocephala', 'medio', 'fijador_nitrogeno', 'pionera')
ON CONFLICT (external_ref) DO NOTHING;

-- Insertar parcelas de ejemplo
INSERT INTO plots (plantation_id, plot_type, length_m, width_m, notes) 
SELECT 
    pl.id,
    'line',
    50.0,
    3.0,
    'Línea de frutales con orientación norte-sur'
FROM plantations pl 
WHERE pl.name = 'Zona Norte'
ON CONFLICT DO NOTHING;

INSERT INTO plots (plantation_id, plot_type, diameter_m, notes) 
SELECT 
    pl.id,
    'island',
    8.0,
    'Isla circular de leguminosas'
FROM plantations pl 
WHERE pl.name = 'Huerta Central'
ON CONFLICT DO NOTHING;

-- Insertar instancias de plantas de ejemplo
INSERT INTO plant_instances (plot_id, species_id, quantity, role, status, position) 
SELECT 
    p.id,
    ps.id,
    5,
    'objetivo',
    'planted',
    'Espaciados cada 10 metros'
FROM plots p
JOIN plantations pl ON p.plantation_id = pl.id
JOIN plant_species ps ON ps.common_name = 'Aguacate Hass'
WHERE pl.name = 'Zona Norte' AND p.plot_type = 'line'
ON CONFLICT DO NOTHING;

INSERT INTO plant_instances (plot_id, species_id, quantity, role, status, position) 
SELECT 
    p.id,
    ps.id,
    20,
    'servicio',
    'germinated',
    'Distribuidos uniformemente en la isla'
FROM plots p
JOIN plantations pl ON p.plantation_id = pl.id
JOIN plant_species ps ON ps.common_name = 'Frijol Caupí'
WHERE pl.name = 'Huerta Central' AND p.plot_type = 'island'
ON CONFLICT DO NOTHING;

-- Insertar plantilla de sugerencias de ejemplo
INSERT INTO suggestion_templates (plantation_id, name, description, rules)
SELECT 
    pl.id,
    'Sistema Agroforestal Básico',
    'Plantilla para sistema agroforestal con frutales y leguminosas',
    '{"densidad_maxima": 2.5, "estratos_requeridos": ["alto", "medio", "bajo"], "sucesion_minima": ["pionera", "secundaria"]}'::jsonb
FROM plantations pl 
WHERE pl.name = 'Zona Norte'
ON CONFLICT DO NOTHING;

-- ============================================================================
-- MENSAJE DE CONFIRMACIÓN
-- ============================================================================

DO $$
BEGIN
    RAISE NOTICE '🌱 Migración 002 - Nuevo Modelo completada!';
    RAISE NOTICE '📊 Nuevas tablas: sites, plantations, plots, plant_species, plant_instances, suggestion_templates';
    RAISE NOTICE '🔍 Vistas: v_plant_instances_full, v_plantation_summary, v_popular_species';
    RAISE NOTICE '⚙️ Funciones: calculate_plot_area, calculate_plant_density, get_plant_instance_hierarchy';
    RAISE NOTICE '🗺️ PostGIS habilitado para geometrías';
    RAISE NOTICE '📋 JSONB para reglas de sugerencias';
    RAISE NOTICE '🌱 Datos de ejemplo insertados con nuevo modelo';
END $$;