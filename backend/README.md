# Sintronia Backend API

API REST para el sistema de agricultura sintrópica desarrollada en Go con Gin.

## 🚀 Inicio Rápido

```bash
# Instalar dependencias
go mod tidy

# Ejecutar en desarrollo
go run cmd/api/main.go

# Compilar
go build -o bin/sintronia-api cmd/api/main.go
```

## 📡 Endpoints

### Especies de Plantas
- `GET /api/v1/plantas` - Listar plantas (público)
- `POST /api/v1/plantas` - Crear planta (requiere auth)
- `GET /api/v1/plantas/:id` - Obtener planta (público)
- `PUT /api/v1/plantas/:id` - Actualizar planta (requiere auth)
- `DELETE /api/v1/plantas/:id` - Eliminar planta (requiere auth)

### Sitios
- `GET /api/v1/sites` - Listar sitios (público)
- `POST /api/v1/sites` - Crear sitio (requiere auth)
- `GET /api/v1/sites/:id` - Obtener sitio (público)
- `DELETE /api/v1/sites/:id` - Eliminar sitio (requiere auth)

### Plantaciones
- `GET /api/v1/plantations` - Listar plantaciones (público)
- `POST /api/v1/plantations` - Crear plantacion (requiere auth)
- `GET /api/v1/plantations/:id` - Obtener plantacion (público)
- `DELETE /api/v1/plantations/:id` - Eliminar plantacion (requiere auth)

### Parcelas sintrópicas
- `GET /api/v1/plots` - Listar parcelas sintrópicas (público)
- `POST /api/v1/plots` - Crear parcela sintrópica (requiere auth)
- `PATCH /api/v1/plots/:id/status` - Actualizar estado (requiere auth)
- `DELETE /api/v1/plots/:id` - Eliminar parcela sintrópica (requiere auth)

### Instancias de plantas
- `GET /api/v1/plant_instances` - Listar instancias de plantas (público)
- `POST /api/v1/plant_instances` - Crear instancia de planta (requiere auth)
- `GET /api/v1/plant_instances/:id` - Obtener instancia de planta (público)
- `PUT /api/v1/plant_instances/:id` - Actualizar instancia de planta (requiere auth)
- `DELETE /api/v1/plant_instances/:id` - Eliminar instancia de planta (requiere auth)

### Plantillas
- `GET /api/v1/suggestion_templates` - Listar plantillas (público)
- `POST /api/v1/suggestion_templates` - Crear plantilla (requiere auth)
- `PATCH /api/v1/suggestion_templates/:id/status` - Actualizar plantilla (requiere auth)
- `DELETE /api/v1/suggestion_templates/:id` - Eliminar plantilla (requiere auth)

### Utilidades
- `GET /api/v1/constants` - Obtener constantes del sistema
- `GET /api/v1/health` - Estado del servicio

## 🔐 Autenticación

Para endpoints protegidos, incluir header:
```
Authorization: Bearer <token>
```

Tokens válidos para testing:
- `test-token` - Usuario normal
- `admin-token` - Administrador
- `user-token` - Usuario normal

## 🏗️ Arquitectura

```
backend/
├── cmd/api/          # Punto de entrada
├── internal/         # Código interno
│   ├── db/           # Conexión a la BD
│   ├── handlers/     # Controladores HTTP
│   ├── middleware/   # Middleware personalizado
│   ├── repositories/ # Repos
│   └── routes/       # Configuración de rutas
├── migrations/       # Código reutilizable
├── pkg/              # Código reutilizable
│    └── models/      # Modelos de datos
docs/                 # Documentos

```

## 🌱 Variables de Entorno

```bash
# Puerto del servidor
PORT=8080

# Base de datos (futuro)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=sintropia
DB_USER=user
DB_PASSWORD=password
```