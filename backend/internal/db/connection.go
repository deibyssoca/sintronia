package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deibys/sintronia/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase inicializa la conexión a PostgreSQL con GORM
func InitDatabase() error {
	// Construir DSN desde variables de entorno
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "sintropia_user"),
		getEnv("DB_PASSWORD", "sintropia_pass"),
		getEnv("DB_NAME", "sintropia"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
	)

	// Configurar logger de GORM
	gormLogger := logger.Default
	if os.Getenv("GIN_MODE") == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	// Conectar a la base de datos
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return fmt.Errorf("error conectando a PostgreSQL: %w", err)
	}

	// Configurar pool de conexiones
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error obteniendo instancia SQL: %w", err)
	}

	// Configuración del pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conexión a PostgreSQL establecida")

	// Auto-migrar modelos
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("error en auto-migración: %w", err)
	}

	return nil
}

// autoMigrate ejecuta las migraciones automáticas de GORM
func autoMigrate() error {
	log.Println("🔄 Ejecutando auto-migraciones...")

	err := DB.AutoMigrate(
		&models.Plant{},
		&models.Location{},
		&models.Arrangement{},
		&models.Planting{},
	)

	if err != nil {
		return fmt.Errorf("error en auto-migración: %w", err)
	}

	log.Println("✅ Auto-migraciones completadas")
	return nil
}

// HealthCheck verifica el estado de la conexión a la base de datos
func HealthCheck() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// CloseDatabase cierra la conexión a la base de datos
func CloseDatabase() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// getEnv obtiene una variable de entorno con valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
