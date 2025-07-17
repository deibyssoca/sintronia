# Multi-stage Dockerfile optimizado para Go 1.24
FROM golang:1.24-alpine AS builder

# Instalar dependencias del sistema
RUN apk add --no-cache git ca-certificates tzdata

# Configurar directorio de trabajo
WORKDIR /app

# Copiar archivos de módulos Go primero (para aprovechar cache de Docker)
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download && go mod verify

# Copiar código fuente
COPY . .

# Compilar aplicación con optimizaciones para Go 1.24
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"\' \
    -a -installsuffix cgo \
    -o main ./cmd/api

# Imagen final mínima
FROM scratch

# Copiar certificados SSL y zona horaria
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copiar binario compilado
COPY --from=builder /app/main /main

# Configurar usuario no-root (por seguridad)
USER 65534:65534

# Exponer puerto
EXPOSE 3000

# Variables de entorno por defecto
ENV PORT=3000
ENV GIN_MODE=release

# Comando por defecto
ENTRYPOINT ["/main"]