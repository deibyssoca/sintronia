# 🌱 Diagrama del Modelo de Datos - Sintropia

## 📊 Diagrama de Entidad-Relación (ASCII)

```
┌─────────────────────────────────────┐
│              PLANT                  │
├─────────────────────────────────────┤
│ 🔑 id (PK)                         │
│    name                            │
│    scientific                      │
│    stratum                         │
│    function                        │
│    succession_stage                │
│    external_id (UNIQUE)            │
│    desired                         │
│    notes                           │
│    created_at                      │
│    updated_at                      │
│    deleted_at                      │
└─────────────────────────────────────┘
                    │
                    │ 1:N
                    ▼
┌─────────────────────────────────────┐
│             PLANTING                │
├─────────────────────────────────────┤
│ 🔑 id (PK)                         │
│ 🔗 plant_id (FK)                   │
│ 🔗 arrangement_id (FK)             │
│    quantity                        │
│    status                          │
│    position                        │
│    notes                           │
│    planted_at                      │
│    created_at                      │
│    updated_at                      │
│    deleted_at                      │
└─────────────────────────────────────┘
                    ▲
                    │ N:1
                    │
┌─────────────────────────────────────┐
│            ARRANGEMENT              │
├─────────────────────────────────────┤
│ 🔑 id (PK)                         │
│ 🔗 location_id (FK)                │
│    name                            │
│    type                            │
│    length                          │
│    width                           │
│    diameter                        │
│    soil_type                       │
│    planting_mode                   │
│    notes                           │
│    created_at                      │
│    updated_at                      │
│    deleted_at                      │
└─────────────────────────────────────┘
                    ▲
                    │ N:1
                    │
┌─────────────────────────────────────┐
│             LOCATION                │
├─────────────────────────────────────┤
│ 🔑 id (PK)                         │
│    name (UNIQUE)                   │
│    notes                           │
│    created_at                      │
│    updated_at                      │
│    deleted_at                      │
└─────────────────────────────────────┘
```

## 🔗 Relaciones

### **1. Location → Arrangement (1:N)**
- Una ubicación puede tener múltiples lechos
- Un lecho pertenece a una sola ubicación

### **2. Arrangement → Planting (1:N)**
- Un lecho puede tener múltiples plantaciones
- Una plantación pertenece a un solo lecho

### **3. Plant → Planting (1:N)**
- Una planta puede estar en múltiples plantaciones
- Una plantación tiene una sola especie de planta

## 📋 Descripción de Entidades

### **🌿 PLANT (Catálogo de Plantas)**
```
Propósito: Inventario de especies vegetales
Ejemplos: Aguacate Hass, Frijol Caupí, Bambú
```

### **📍 LOCATION (Ubicaciones/Zonas)**
```
Propósito: Organización territorial del terreno
Ejemplos: "Zona Norte", "Huerta Principal", "Ladera Este"
```

### **🏗️ ARRANGEMENT (Lechos de Cultivo)**
```
Propósito: Disposiciones específicas de plantación
Tipos: Línea, Isla, Gremio
Ejemplos: "Línea de Frutales", "Isla de Leguminosas"
```

### **🌱 PLANTING (Plantaciones)**
```
Propósito: Instancias específicas de plantas en lechos
Ejemplos: "5 aguacates en Línea Norte", "20 frijoles en Isla Central"
```

## 🎯 Flujo de Datos

```
1. Crear LOCATION
   ↓
2. Crear ARRANGEMENT en esa location
   ↓
3. Agregar PLANT al catálogo
   ↓
4. Crear PLANTING (vincular plant + arrangement)
```

## 📊 Cardinalidades

```
Location (1) ←→ (N) Arrangement
Arrangement (1) ←→ (N) Planting  
Plant (1) ←→ (N) Planting
```

## 🔍 Índices Importantes

```sql
-- Búsquedas frecuentes
CREATE INDEX idx_plants_name ON plants(name);
CREATE INDEX idx_plants_stratum ON plants(stratum);
CREATE INDEX idx_plants_desired ON plants(desired);
CREATE INDEX idx_plantings_status ON plantings(status);
CREATE INDEX idx_arrangements_type ON arrangements(type);
```

## 🚀 Queries Típicas

### **Plantas por estrato:**
```sql
SELECT * FROM plants WHERE stratum = 'alto';
```

### **Plantaciones de una ubicación:**
```sql
SELECT p.*, pl.name as plant_name, a.name as arrangement_name
FROM plantings p
JOIN plants pl ON p.plant_id = pl.id
JOIN arrangements a ON p.arrangement_id = a.id
JOIN locations l ON a.location_id = l.id
WHERE l.name = 'Zona Norte';
```

### **Densidad por lecho:**
```sql
SELECT 
    a.name,
    a.type,
    SUM(p.quantity) as total_plants,
    (a.length * a.width) as area_m2,
    SUM(p.quantity) / (a.length * a.width) as density
FROM arrangements a
JOIN plantings p ON a.id = p.arrangement_id
GROUP BY a.id;
```