package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/deibys/sintronia/internal/handlers"
	"github.com/deibys/sintronia/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Registra la validación "onlyletters"
		v.RegisterValidation("onlyletters", OnlyLetters)
	}
}

// Planta representa la estructura de una planta.
// En el struct, usamos "onlyletters" en el campo
type Plant struct {
	Nombre           string `json:"nombre" binding:"required,onlyletters"`
	NombreCientifico string `json:"nombre_cientifico" binding:"required"`
	Estrato          string `json:"estrato" binding:"required,oneof=emergente alto medio bajo"`
	Sucesion         string `json:"sucesion" binding:"required,oneof=pionera secundaria climax"`
	IsObjetivo       bool   `json:"is_objetivo"`
}

// Lista que simula la persistencia en memoria.
var PlantasRegistradas []Plant

// NewRouter configura el router con Gin y define los endpoints.
// NewRouter configura el router con todas las rutas y middlewares globales.
func NewRouter() *gin.Engine {
	// Obtiene el valor de la variable de entorno GIN_MODE.
	// Si no está definida, por defecto será "debug".
	mode := os.Getenv("GIN_MODE")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	// Con gin.New() evitamos el logger y recovery por defecto para usar los nuestros
	router := gin.New()

	// Middlewares globales
	router.Use(middleware.CustomLogger())
	router.Use(middleware.ErrorHandler())
	// Use the proper CORS middleware instead of hardcoded configuration
	//router.Use(middleware.CORSMiddleware())
	// Configuración básica de CORS:
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
			"x-permapeople-key-id", "Cache-Control", "ngrok-skip-browser-warning", // <- agregamos este
		},
		AllowCredentials: false, // ⚠️ debe estar en false si AllowAllOrigins es true
		MaxAge:           12 * time.Hour,

		// // Permite todos los orígenes; en producción es mejor especificar tus orígenes permitidos.
		// AllowOrigins: []string{"*"},
		// //AllowOrigins: []string{"https://stackblitz.com", "http://localhost:5173"},
		// // Permite métodos comunes:
		// AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// // Permite encabezados comunes:
		// AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "x-permapeople-key-id"},
		// // Opcional: permite el envío de credenciales.
		// AllowCredentials: true,
		// // Establece un tiempo de vida para la configuración de CORS
		// MaxAge: 12 * time.Hour,
	}))
	// Puedes aplicar el AuthMiddleware globalmente si lo considerás necesario,
	// o solo en rutas específicas
	// router.Use(middleware.AuthMiddleware())

	RegisterRoutes(router)

	router.GET("/error", func(c *gin.Context) {
		// Forzamos un error agregándolo al context
		c.Error(fmt.Errorf("error forzado para prueba "))
		// Luego, no se llama a c.Next() o se deja caer sin responder
	})

	router.GET("/api/plants", middleware.AuthMiddleware(), func(c *gin.Context) {
		log.Println("**************************************   ")
		// Llamada a la API de Permapeople
		req, err := http.NewRequest("GET", "https://permapeople.org/api/plants", nil)
		if err != nil {
			c.Error(fmt.Errorf("error creando request"))

			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al crear la solicitud"})
			return
		}

		authHeader := c.GetHeader("Authorization") // Bearer token
		keyID := c.GetHeader("x-permapeople-key-id")
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")
		// Si necesitás pasar una API key (ejemplo)
		req.Header.Set("x-permapeople-key-id", keyID)
		req.Header.Set("x-permapeople-key-secret", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.Error(fmt.Errorf("error haciendo request"))
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al llamar a la API externa"})
			return
		}
		defer resp.Body.Close()

		// Leer respuesta y reenviarla
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.Error(fmt.Errorf("error leyendo respuesta"))
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al leer la respuesta"})
			return
		}

		c.Data(resp.StatusCode, "application/json", body)

	})

	// Endpoint POST /plant
	router.POST("/plants2", middleware.AuthMiddleware(), func(c *gin.Context) {
		var nuevaPlanta Plant

		// Lógica para crear una planta
		if err := c.ShouldBindJSON(&nuevaPlanta); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		PlantasRegistradas = append(PlantasRegistradas, nuevaPlanta)
		c.JSON(http.StatusCreated, gin.H{"mensaje": "Planta creada"})
	})

	return router
}

func RegisterRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")
	{
		// Endpoint de prueba, solo accesible públicamente
		api.GET("/public/saludo", handlers.SaludoHandler)

	}

	// Grupo para la API protegida (por ejemplo, para plantas)
	plantas := api.Group("/plantas")
	{
		// Rutas públicas (sin autenticación)

		plantas.GET("", handlers.GetPlantsHandler)
		plantas.GET("/:id", handlers.GetPlantHandler)

		// Rutas protegidas (con autenticación)
		plantasAuth := plantas.Group("")
		plantasAuth.Use(middleware.AuthMiddleware())
		{
			plantasAuth.POST("", handlers.CreatePlantHandler)
			plantasAuth.PUT("/:id", handlers.UpdatePlantHandler)
			plantasAuth.DELETE("/:id", handlers.DeletePlantHandler)
		}
	}

	// ubicaciones := api.Group("/locations")
	// {
	// 	// Rutas de ubicaciones
	// 	ubicaciones.POST("", handlers.CreateLocationHandler)
	// 	ubicaciones.GET("", handlers.GetLocationsHandler)
	// 	ubicaciones.GET("/:id", handlers.GetLocationHandler)
	// 	ubicaciones.DELETE("/:id", handlers.DeleteLocationHandler)
	// }

	// disposicion := api.Group("/arrangement")
	// {
	// 	// Rutas de lechos
	// 	disposicion.POST("", handlers.CreateArrangementHandler)
	// 	disposicion.GET("", handlers.GetArrangementsHandler)
	// 	disposicion.GET("/:id", handlers.GetArrangementHandler)
	// 	disposicion.DELETE("/:id", handlers.DeleteArrangementHandler)
	// }

	// plantaciones := api.Group("/plantings")
	// {
	// 	// Rutas de plantaciones
	// 	plantaciones.POST("", handlers.CreatePlantingHandler)
	// 	plantaciones.GET("", handlers.GetPlantingsHandler)
	// 	plantaciones.PATCH("/:id/status", handlers.UpdatePlantingStatusHandler)
	// 	plantaciones.DELETE("/:id", handlers.DeletePlantingHandler)
	// }

	// Rutas de utilidad
	api.GET("/constants", handlers.GetConstantsHandler)

	// Ruta de salud
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "sintropia-api",
		})
	})

	// Grupo de rutas para usuarios (si llegas a implementarlo)
	// userGroup := router.Group("/usuarios")
	// {
	// 	// Por ejemplo, podrías aplicar otro middleware específico para usuarios
	// 	userGroup.POST("/registro", handlers.RegisterUserHandler)
	// 	userGroup.POST("/login", handlers.LoginUserHandler)
	// 	// etc.
	// }
}

// Función de validación personalizada que verifica que el nombre solo tenga letras
func OnlyLetters(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	return re.MatchString(value)
}
