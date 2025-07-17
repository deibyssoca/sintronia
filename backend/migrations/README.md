# 🗄️ Migraciones de Base de Datos - Backend Sintropia

Este directorio contiene las migraciones SQL para el sistema de agricultura sintrópica.

## 📋 Migraciones disponibles

### `001_initial_schema.sql` - Esquema inicial
- ✅ Tablas principales: `locations`, `plants`, `arrangements`, `plantings`
- ✅ Índices optimizados para performance
- ✅ Constraints y validaciones
- ✅ Vistas útiles para consultas complejas
- ✅ Funciones SQL para cálculos
- ✅ Triggers para `updated_at` automático
- ✅ Datos de ejemplo para testing

## 🚀 Cómo ejecutar las migraciones

### Opción 1: PostgreSQL directo
```bash
# Ejecutar migración inicial
psql -d sintropia -f backend/migrations/001_initial_schema.sql
```

### Opción 2: Desde Go (futuro)
```go
// En tu aplicación Go, podrías usar:
// - golang-migrate/migrate
// - pressly/goose
// - Otras herramientas de migración
```

### Opción 3: Docker
```bash
# Si usas Docker Compose
docker-compose exec postgres psql -U sintropia_user -d sintropia -f /migrations/001_initial_schema.sql
```

## 📊 Estructura del esquema

```
locations (ubicaciones)
    ↓ 1:N
arrangements (lechos)
    ↓ 1:N  
plantings (plantaciones)
    ↑ N:1
plants (catálogo)
```

## 🔍 Vistas disponibles

- **`v_plantings_full`** - Plantaciones con información completa
- **`v_location_summary`** - Resumen estadístico por ubicación

## ⚙️ Funciones disponibles

- **`calculate_arrangement_area()`** - Calcula área de lechos
- **`calculate_planting_density()`** - Calcula densidad de plantación
- **`update_updated_at_column()`** - Actualiza timestamp automáticamente

## 🌱 Datos de ejemplo

La migración incluye datos de ejemplo:
- 3 ubicaciones (Zona Norte, Huerta Central, Ladera Este)
- 6 plantas (Aguacate, Frijol, Bambú, Plátano, Moringa, Leucaena)
- 2 lechos (Línea de Frutales, Isla de Leguminosas)

## 📝 Notas importantes

- ✅ **Idempotente** - Se puede ejecutar múltiples veces sin errores
- ✅ **Soft delete** - Todas las tablas usan `deleted_at`
- ✅ **Timestamps** - `created_at` y `updated_at` automáticos
- ✅ **Constraints** - Validaciones a nivel de base de datos
- ✅ **Índices** - Optimizados para consultas frecuentes

## 🔄 Diferencia con Supabase

Este directorio (`backend/migrations/`) es para el backend Go con PostgreSQL directo.
Las migraciones de Supabase están en `supabase/migrations/` y son gestionadas por Supabase.

## 🔄 Próximas migraciones

Para agregar nuevas migraciones:
1. Crear archivo `002_nombre_descriptivo.sql`
2. Usar `IF NOT EXISTS` para evitar conflictos
3. Documentar cambios en este README
4. Probar en entorno de desarrollo primero