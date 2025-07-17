#!/bin/bash

echo "🚀 Configurando entorno de desarrollo Backend Sintropia con Go 1.24..."

# Asegurar que estamos en el directorio correcto
cd /workspace

# Verificar versión de Go
echo "📋 Verificando versión de Go..."
go version

# Crear directorios de cache si no existen
mkdir -p .devcontainer/go-cache
mkdir -p .devcontainer/vscode-extensions

# Configurar backend
echo "📦 Configurando backend Go 1.24..."
if [ -d "backend" ]; then
    cd backend
    
    # Verificar que go.mod tiene la versión correcta
    if ! grep -q "go 1.24" go.mod; then
        echo "⚠️ Actualizando go.mod para Go 1.24..."
        sed -i 's/go [0-9]\+\.[0-9]\+/go 1.24/' go.mod
    fi
    
    # Limpiar módulos y descargar dependencias
    echo "  - Limpiando y descargando módulos Go..."
    go clean -modcache 2>/dev/null || true
    go mod tidy
    go mod download
    
    # Verificar que las dependencias están correctas
    echo "  - Verificando dependencias..."
    go mod verify
    
    # Pre-compilar dependencias comunes para acelerar builds
    echo "  - Pre-compilando dependencias..."
    go build -i ./... 2>/dev/null || true
    
    # Verificar que el proyecto compila
    echo "  - Verificando compilación..."
    go build -o /tmp/test-build ./cmd/api && rm -f /tmp/test-build
    
    cd /workspace
else
    echo "⚠️ Directorio backend no encontrado"
fi

# Crear archivos de configuración si no existen
echo "⚙️ Creando archivos de configuración..."

# Backend .env
if [ -f "backend/.env.example" ] && [ ! -f "backend/.env" ]; then
    cp backend/.env.example backend/.env
    echo "✅ Creado backend/.env desde .env.example"
elif [ ! -f "backend/.env" ]; then
    # Crear .env básico si no existe el ejemplo
    cat > backend/.env << 'EOF'
# Variables de entorno para el backend Go 1.24
PORT=3000
GIN_MODE=debug

# Base de datos PostgreSQL
DB_HOST=postgres
DB_PORT=5432
DB_NAME=sintropia
DB_USER=sintropia_user
DB_PASSWORD=sintropia_pass
DB_SSLMODE=disable

# JWT (futuro)
JWT_SECRET=tu-secreto-super-seguro-aqui-go-1-24
JWT_EXPIRATION=24h

# Configuración de CORS - Permitir todos los orígenes para desarrollo
CORS_ALLOWED_ORIGINS=*
CORS_ALLOW_ALL=true

# Límites de paginación
DEFAULT_MAX_PAGINATION_LIMIT=100
ADMIN_MAX_PAGINATION_LIMIT=1000

# Go 1.24 específico
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org
GO111MODULE=on
EOF
    echo "✅ Creado backend/.env con configuración para Go 1.24"
fi

# Configuración de Air para hot reload del backend (optimizada para Go 1.24)
if [ -d "backend" ] && [ ! -f "backend/.air.toml" ]; then
    cat > backend/.air.toml << 'EOF'
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/api"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "node_modules", ".git"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false

# Configuración específica para Go 1.24
[env]
  GO111MODULE = "on"
  GOPROXY = "https://proxy.golang.org,direct"
  GOSUMDB = "sum.golang.org"
  PORT = "3000"
  GIN_MODE = "debug"
  CORS_ALLOW_ALL = "true"
EOF
    echo "✅ Creado backend/.air.toml para hot reload con Go 1.24"
fi

# Crear Makefile para comandos backend (actualizado para Go 1.24)
if [ ! -f "Makefile.backend" ]; then
    cat > Makefile.backend << 'EOF'
.PHONY: help dev build clean install test go-version lint

# Mostrar ayuda
help:
	@echo "🌱 Sintropia Backend - Comandos disponibles (Go 1.24):"
	@echo ""
	@echo "  make dev          - Iniciar backend con hot reload"
	@echo "  make build        - Compilar el backend"
	@echo "  make clean        - Limpiar archivos temporales"
	@echo "  make install      - Instalar dependencias"
	@echo "  make test         - Ejecutar tests"
	@echo "  make lint         - Verificar código"
	@echo "  make go-version   - Mostrar versión de Go"
	@echo ""

# Verificar versión de Go
go-version:
	@echo "📋 Versión de Go:"
	@go version
	@echo ""
	@echo "📋 Variables de entorno Go:"
	@go env GOVERSION GOOS GOARCH GOROOT GOPATH

# Desarrollo con hot reload
dev:
	@echo "🔧 Iniciando backend con hot reload (Go 1.24)..."
	cd backend && air

# Compilar backend
build:
	@echo "📦 Compilando backend con Go 1.24..."
	mkdir -p bin
	cd backend && go build -ldflags="-w -s" -o ../bin/sintropia-api ./cmd/api

# Limpiar archivos temporales
clean:
	@echo "🧹 Limpiando archivos temporales..."
	rm -rf backend/tmp
	rm -rf bin
	cd backend && go clean -cache -modcache -testcache

# Instalar dependencias
install:
	@echo "📦 Instalando dependencias (Go 1.24)..."
	cd backend && go mod tidy && go mod download && go mod verify

# Ejecutar tests
test:
	@echo "🧪 Ejecutando tests con Go 1.24..."
	cd backend && go test -v ./...

# Verificar código Go
lint:
	@echo "🔍 Verificando código Go..."
	cd backend && golangci-lint run ./...

# Actualizar dependencias de Go
update-deps:
	@echo "🔄 Actualizando dependencias de Go..."
	cd backend && go get -u ./... && go mod tidy
EOF
    echo "✅ Creado Makefile.backend con comandos para Go 1.24"
fi

echo ""
echo "🎉 ¡Configuración de backend completada para Go 1.24!"
echo ""
echo "📋 Información del entorno:"
go version
echo ""
echo "📋 Comandos disponibles:"
echo "  make -f Makefile.backend help         # Ver todos los comandos"
echo "  make -f Makefile.backend go-version   # Verificar versión de Go"
echo "  make -f Makefile.backend dev          # Iniciar backend con hot reload"
echo "  cd backend && go run cmd/api/main.go  # Iniciar backend directamente"
echo ""
echo "📁 Estructura del backend:"
echo "  backend/cmd/api/                      # Punto de entrada"
echo "  backend/internal/                     # Lógica interna"
echo "  backend/pkg/                          # Modelos y utilidades"
echo ""
echo "🌐 URLs cuando esté ejecutándose:"
echo "  Backend:  http://localhost:3000"
echo "  Health:   http://localhost:3000/api/v1/health"
echo "  pgAdmin:  http://localhost:5050"
echo ""
echo "🚀 Para iniciar: make -f Makefile.backend dev"