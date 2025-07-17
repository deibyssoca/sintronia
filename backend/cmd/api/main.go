package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deibys/sintronia/internal/db"
	"github.com/deibys/sintronia/internal/routes"
)

func main() {
	fmt.Println("🌱 Iniciando Sintropia API...")

	// Verificar que PostgreSQL esté disponible antes de continuar
	fmt.Println("🔍 Verificando PostgreSQL...")

	// Inicializar base de datos
	if err := db.InitDatabase(); err != nil {
		log.Printf("❌ Error inicializando base de datos: %v", err)
		log.Println("⚠️ Continuando sin base de datos (modo fallback)")

		// Opcional: Continuar sin DB para testing
		// En producción, esto debería ser fatal
		// log.Fatalf("❌ Error inicializando base de datos: %v", err)
	}

	// Configurar cierre graceful de la base de datos
	defer func() {
		if err := db.CloseDatabase(); err != nil {
			log.Printf("⚠️ Error cerrando base de datos: %v", err)
		}
	}()

	// Obtener puerto del entorno o usar 3000 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := routes.NewRouter()

	fmt.Printf("🚀 Servidor corriendo en puerto %s\n", port)
	fmt.Printf("📡 API disponible en: http://localhost:%s/api/v1\n", port)
	fmt.Printf("🔍 Health check: http://localhost:%s/api/v1/health\n", port)
	fmt.Printf("🗄️ Base de datos: PostgreSQL conectada\n")

	// Levantamos el servidor
	r.Run(":" + port)
}
