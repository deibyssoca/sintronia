# 🌱 Esquema de Base de Datos v2.0 - Sistema Sintronia

## 📊 Modelo Jerárquico

### 🏗️ Estructura Jerárquica

```
SITE (Sitio/Terreno)
  ↓ 1:N
PLANTATION (Plantación/Zona)
  ↓ 1:N  
PLOT (Parcela/Lecho)
  ↓ 1:N
PLANT_INSTANCE (Instancia de Planta)
  ↑ N:1
PLANT_SPECIES (Especie de Planta)
```

## 📋 Descripción de Entidades

### **🏞️ SITE (Sitios/Terrenos)**
```sql
sites (
    id, name, area_m2, length_m, width_m, notes,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Representa el terreno principal o finca
- **Ejemplos**: "Finca La Esperanza", "Parcela Experimental"
- **Características**: Área total, dimensiones generales

### **🌿 PLANTATION (Plantaciones/Zonas)**
```sql
plantations (
    id, site_id, name, area_m2, notes,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Zonas de cultivo dentro de un sitio
- **Ejemplos**: "Zona Norte", "Huerta Central", "Área Agroforestal"
- **Características**: Área específica, propósito definido

### **🌱 PLANT_SPECIES (Especies de Plantas)**
```sql
plant_species (
    id, common_name, scientific_name, stratum, 
    function_ecol, succession_stage, external_ref, notes,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Catálogo de especies vegetales
- **Ejemplos**: "Aguacate Hass", "Frijol Caupí", "Bambú Guadua"
- **Características**: Clasificación sintrópica completa

### **📐 PLOT (Parcelas/Lechos)**
```sql
plots (
    id, plantation_id, plot_type, length_m, width_m, 
    diameter_m, geometry, notes,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Disposiciones específicas de plantación
- **Tipos**: `line` (línea), `island` (isla), `guild` (gremio)
- **Características**: Dimensiones, geometría opcional (PostGIS)

### **🌿 PLANT_INSTANCE (Instancias de Plantas)**
```sql
plant_instances (
    id, plot_id, species_id, quantity, role, status, 
    position, planted_at, notes,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Plantas específicas en parcelas específicas
- **Roles**: `objetivo`, `servicio`, `acompañante`
- **Estados**: `planned`, `germinated`, `planted`, `established`, `productive`, `dormant`, `dead`

### **💡 SUGGESTION_TEMPLATE (Plantillas de Sugerencias)**
```sql
suggestion_templates (
    id, plantation_id, name, description, rules,
    created_at, updated_at, deleted_at
)
```
- **Propósito**: Plantillas para sugerir plantaciones
- **Características**: Reglas JSON para densidad, estratos, sucesión

## 🔗 Relaciones Principales

### **1. Site → Plantation (1:N)**
- Un sitio puede tener múltiples plantaciones
- Una plantación pertenece a un solo sitio

### **2. Plantation → Plot (1:N)**
- Una plantación puede tener múltiples parcelas
- Una parcela pertenece a una sola plantación

### **3. Plot → PlantInstance (1:N)**
- Una parcela puede tener múltiples instancias de plantas
- Una instancia pertenece a una sola parcela

### **4. PlantSpecies → PlantInstance (1:N)**
- Una especie puede estar en múltiples instancias
- Una instancia tiene una sola especie

### **5. Plantation → SuggestionTemplate (1:N)**
- Una plantación puede tener múltiples plantillas
- Una plantilla pertenece a una sola plantación

## 🎯 Modelo

1. **🏗️ Jerarquía más clara**
   - Separación entre sitio, plantación y parcela
   - Mejor organización territorial

2. **🎯 Roles de plantas definidos**
   - `objetivo`: Plantas de producción principal
   - `servicio`: Plantas de apoyo ecológico
   - `acompañante`: Plantas complementarias

3. **📐 Geometrías avanzadas**
   - Soporte PostGIS para formas complejas
   - GeoJSON para posicionamiento preciso

4. **💡 Sistema de sugerencias**
   - Plantillas con reglas JSON
   - Automatización de recomendaciones

5. **📊 Mejor escalabilidad**
   - Estructura más normalizada
   - Consultas más eficientes

## 🔍 Vistas Principales

### **v_plant_instances_full**
Vista completa con toda la jerarquía:
```sql
SELECT site_name, plantation_name, plot_type, 
       species_common_name, quantity, role, status,
       plot_area_m2, density_per_m2
FROM v_plant_instances_full;
```

### **v_plantation_summary**
Resumen estadístico por plantación:
```sql
SELECT name, total_plots, total_plant_instances, 
       unique_species, total_plant_count
FROM v_plantation_summary;
```

### **v_popular_species**
Especies más utilizadas:
```sql
SELECT common_name, usage_count, total_quantity,
       plots_used, plantations_used
FROM v_popular_species;
```

## ⚙️ Funciones Útiles

### **calculate_plot_area()**
```sql
SELECT calculate_plot_area('line', 50.0, 3.0, NULL); -- 150.0
SELECT calculate_plot_area('island', NULL, NULL, 8.0); -- ~50.27
```

### **calculate_plant_density()**
```sql
SELECT calculate_plant_density(plot_id, quantity);
```

### **get_plant_instance_hierarchy()**
```sql
SELECT * FROM get_plant_instance_hierarchy(instance_id);
```

### **Conceptos:**
- **Roles de plantas**: objetivo, servicio, acompañante
- **Estados simplificados**: planned, germinated, planted, etc.
- **Plantillas de sugerencias**: automatización de recomendaciones
- **Geometrías PostGIS**: formas complejas y precisas
