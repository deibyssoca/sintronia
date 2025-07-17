# Sintropia Backend API

API REST para el sistema de agricultura sintrópica desarrollada en Go con Gin.

## 🚀 Inicio Rápido

```bash
# Instalar dependencias
go mod tidy

# Ejecutar en desarrollo
go run cmd/api/main.go

# Compilar
go build -o bin/sintropia-api cmd/api/main.go
```

## 📡 Endpoints

### Plantas
- `GET /api/v1/plantas` - Listar plantas (público)
- `POST /api/v1/plantas` - Crear planta (requiere auth)
- `GET /api/v1/plantas/:id` - Obtener planta (público)
- `PUT /api/v1/plantas/:id` - Actualizar planta (requiere auth)
- `DELETE /api/v1/plantas/:id` - Eliminar planta (requiere auth)

### Ubicaciones
- `GET /api/v1/locations` - Listar ubicaciones (público)
- `POST /api/v1/locations` - Crear ubicación (requiere auth)
- `GET /api/v1/locations/:id` - Obtener ubicación (público)
- `DELETE /api/v1/locations/:id` - Eliminar ubicación (requiere auth)

### Lechos
- `GET /api/v1/beds` - Listar lechos (público)
- `POST /api/v1/beds` - Crear lecho (requiere auth)
- `GET /api/v1/beds/:id` - Obtener lecho (público)
- `DELETE /api/v1/beds/:id` - Eliminar lecho (requiere auth)

### Plantaciones
- `GET /api/v1/plantings` - Listar plantaciones (público)
- `POST /api/v1/plantings` - Crear plantación (requiere auth)
- `PATCH /api/v1/plantings/:id/status` - Actualizar estado (requiere auth)
- `DELETE /api/v1/plantings/:id` - Eliminar plantación (requiere auth)

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
│   ├── handlers/     # Controladores HTTP
│   ├── middleware/   # Middleware personalizado
│   └── routes/       # Configuración de rutas
└── pkg/             # Código reutilizable
    └── models/      # Modelos de datos
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