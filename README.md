# 🌱 Sintronia - Sistema de Agricultura Sintrópica

Sistema completo para la gestión y planificación de proyectos de agricultura sintrópica.

## 🏗️ Arquitectura

Este proyecto está dividido en dos aplicaciones independientes:

- **Backend**: API REST en Go con Gin framework
- **Frontend**: Aplicación web en React con TypeScript

```
sintropia/
├── backend/          # API REST en Go
│   ├── cmd/api/          # Punto de entrada
│   ├── internal/         # Código interno
│   │   ├── db/           # Conexión a la BD
│   │   ├── handlers/     # Controladores HTTP
│   │   ├── middleware/   # Middleware personalizado
│   │   ├── repositories/ # Repos
│   │   └── routes/       # Configuración de rutas
│   ├── migrations/       # Código reutilizable
│   ├── pkg/              # Código reutilizable
│   │    └── models/      # Modelos de datos
│   docs/                 # Documentos
│
├── frontend/          # Aplicación web en React con TypeScript
```

## 🚀 Inicio Rápido

### Prerrequisitos
- **Go 1.24+** para el backend
- **Node.js 18+** para el frontend

### Desarrollo Local

1. **Clonar el repositorio**
```bash
git clone <repo-url>
cd sintronia
```

2. **Iniciar Backend**
```bash
cd backend
go mod tidy
go run cmd/api/main.go
# Servidor corriendo en http://localhost:3000
```

3. **Iniciar Frontend** (en otra terminal)
```bash
cd frontend
npm install
npm run dev
# App corriendo en http://localhost:5173
```

### Con Docker (Recomendado)

```bash
# Iniciar todo el stack
docker-compose up -d

# Backend: http://localhost:3000
# Frontend: http://localhost:5173
# PostgreSQL: localhost:5432
```

## 🌿 Funcionalidades

### ✅ Gestión de catalogo de especies de Plantas
- Inventario completo de especies
- Clasificación por estratos y funciones ecológicas
- Etapas sucesionales
- Lista de plantas

### ✅ Sitios y lugares de plantación
- Creación de sitios(espacios) y plantaciones    -- En construcción

### ✅ Diseño de parcelas sintrópicas para las zonas de plantación
- Vinculación especies a la parcela 
- Dispocición de cada especie según tipo de parcela.
- Cálculo de densidades
- Seguimiento temporal

### 🔄 En Desarrollo
- [ ] Integración con base de datos PostgreSQL
- [ ] Sistema de usuarios y autenticación JWT
- [ ] Dashboard con métricas y gráficos
- [ ] Exportación de reportes
- [ ] API de integración con Permapeople

## 🔐 Autenticación

Para testing, usar estos tokens en el header `Authorization: Bearer <token>`:

- `admin-token` - Permisos de administrador
- `user-token` - Usuario estándar
- `test-token` - Usuario de prueba

## 📡 API Endpoints

### Públicos (sin autenticación)
- `GET /api/v1/plantas` - Listar plantas (público)
- `GET /api/v1/plantas/:id` - Obtener planta (público)
- `GET /api/v1/sites` - Listar sitios (público)
- `GET /api/v1/plantations` - Listar plantaciones (público)
- `GET /api/v1/plots` - Listar parcelas sintrópicas (público)

### Protegidos (requieren autenticación)
- `POST /api/v1/plantas` - Crear planta (requiere auth)
- `POST /api/v1/sites` - Crear sitio (requiere auth)
- `POST /api/v1/plantations` - Crear plantacion (requiere auth)
- Y más...

## 🛠️ Stack Tecnológico

### Backend
- **Go 1.24** - Lenguaje principal
- **Gin** - Framework web
- **PostgreSQL** - Base de datos 
- **Gorm** - ORM
- **JWT** - Autenticación (próximamente)

### Frontend
- **React 18** - Biblioteca de UI
- **TypeScript** - Tipado estático
- **Vite** - Build tool
- **Tailwind CSS** - Framework CSS
- **Lucide React** - Iconos

## 🌍 Principios de Agricultura Sintrópica

Este sistema está basado en los principios de bosques sintrópicos:

- **Estratificación**: 7 estratos de vegetación
- **Sucesión**: 4 etapas sucesionales
- **Funciones Ecológicas**: 14 funciones principales
- **Densidad y Tiempo**: Optimización espacio-temporal

## 🤝 Contribuir

1. Fork el proyecto
2. Crear rama feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Crear Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.

## 🙏 Agradecimientos

- **Ernst Götsch** - Por los principios de agricultura sintrópica
- **Permapeople** - Por la inspiración en el manejo de datos de plantas
- Comunidad de agricultura regenerativa