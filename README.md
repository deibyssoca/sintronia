# 🌱 Sintropia - Sistema de Agricultura Sintrópica

Sistema completo para la gestión y planificación de proyectos de agricultura sintrópica, inspirado en los principios de Ernst Götsch.

## 🏗️ Arquitectura

Este proyecto está dividido en dos aplicaciones independientes:

- **Backend**: API REST en Go con Gin framework
- **Frontend**: Aplicación web en React con TypeScript

```
sintropia/
├── backend/          # API REST en Go
│   ├── cmd/api/      # Punto de entrada
│   ├── internal/     # Lógica interna
│   └── pkg/          # Modelos y utilidades
└── frontend/         # App React
    ├── src/          # Código fuente
    └── public/       # Archivos estáticos
```

## 🚀 Inicio Rápido

### Prerrequisitos
- **Go 1.21+** para el backend
- **Node.js 18+** para el frontend

### Desarrollo Local

1. **Clonar el repositorio**
```bash
git clone <repo-url>
cd sintropia
```

2. **Iniciar Backend**
```bash
cd backend
go mod tidy
go run cmd/api/main.go
# Servidor corriendo en http://localhost:8080
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

# Backend: http://localhost:8080
# Frontend: http://localhost:5173
# PostgreSQL: localhost:5432
```

## 🌿 Funcionalidades

### ✅ Gestión de Plantas
- Inventario completo de especies
- Clasificación por estratos y funciones ecológicas
- Etapas sucesionales según Ernst Götsch
- Lista de plantas deseadas

### ✅ Ubicaciones y Lechos
- Organización por zonas/áreas
- Lechos lineales, islas y gremios
- Cálculo automático de áreas
- Tipos de suelo y modalidades de plantación

### ✅ Planificación de Plantaciones
- Vinculación plantas-lechos
- Estados del ciclo de vida
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
- `GET /api/v1/plantas` - Listar plantas
- `GET /api/v1/locations` - Listar ubicaciones
- `GET /api/v1/beds` - Listar lechos
- `GET /api/v1/plantings` - Listar plantaciones
- `GET /api/v1/constants` - Constantes del sistema

### Protegidos (requieren autenticación)
- `POST /api/v1/plantas` - Crear planta
- `PUT /api/v1/plantas/:id` - Actualizar planta
- `DELETE /api/v1/plantas/:id` - Eliminar planta
- Y más...

## 🛠️ Stack Tecnológico

### Backend
- **Go 1.21** - Lenguaje principal
- **Gin** - Framework web
- **PostgreSQL** - Base de datos (próximamente)
- **JWT** - Autenticación (próximamente)

### Frontend
- **React 18** - Biblioteca de UI
- **TypeScript** - Tipado estático
- **Vite** - Build tool
- **Tailwind CSS** - Framework CSS
- **Lucide React** - Iconos

## 🌍 Principios de Agricultura Sintrópica

Este sistema está basado en los principios de **Ernst Götsch**:

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