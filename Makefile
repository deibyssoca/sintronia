.PHONY: help dev dev-backend dev-frontend build clean install test docker-up docker-down logs

# Mostrar ayuda
help:
	@echo "🌱 Sintropia - Comandos disponibles:"
	@echo ""
	@echo "  make dev          - Iniciar backend y frontend simultáneamente"
	@echo "  make dev-backend  - Iniciar solo el backend con hot reload"
	@echo "  make dev-frontend - Iniciar solo el frontend"
	@echo "  make build        - Compilar el proyecto"
	@echo "  make clean        - Limpiar archivos temporales"
	@echo "  make install      - Instalar dependencias"
	@echo "  make test         - Ejecutar tests"
	@echo "  make docker-up    - Iniciar con Docker Compose"
	@echo "  make docker-down  - Detener Docker Compose"
	@echo "  make logs         - Ver logs de Docker"
	@echo ""

# Desarrollo completo
dev:
	@echo "🚀 Iniciando desarrollo completo..."
	npm run dev

# Solo backend
dev-backend:
	@echo "🔧 Iniciando backend ..."
	cd backend && go run ./cmd/api

# Solo frontend  
dev-frontend:
	@echo "🎨 Iniciando frontend..."
	cd frontend && npm run dev

# Compilar proyecto
build:
	@echo "📦 Compilando proyecto..."
	mkdir -p bin
	cd backend && go build -o ../bin/sintropia-api ./cmd/api
	cd frontend && npm run build

# Limpiar archivos temporales
clean:
	@echo "🧹 Limpiando archivos temporales..."
	rm -rf backend/tmp
	rm -rf frontend/dist
	rm -rf bin
	cd backend && go clean -cache -modcache -testcache

# Instalar dependencias
install:
	@echo "📦 Instalando dependencias..."
	cd backend && go mod tidy && go mod download
# 	cd frontend && npm install

# Ejecutar tests
test:
	@echo "🧪 Ejecutando tests..."
	cd backend && go test ./...
	cd frontend && npm test 2>/dev/null || echo "No hay tests configurados en frontend"

# Docker Compose
docker-up:
	@echo "🐳 Iniciando con Docker Compose..."
	docker-compose -f docker-compose.dev.yml up -d

docker-down:
	@echo "🐳 Deteniendo Docker Compose..."
	docker-compose -f docker-compose.dev.yml down

logs:
	@echo "📋 Mostrando logs de Docker..."
	docker-compose -f docker-compose.dev.yml logs -f