package main

import (
	"inventory_backend/config"
	"inventory_backend/handler"
	"inventory_backend/repository"
	"inventory_backend/service"
	"log"
	"net/http"

	jwtmiddleware "inventory_backend/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Load env
	config.LoadEnv()

	// 2. Init DB
	db := config.DBInit()

	e := echo.New()

	// Cors setup
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // frontend
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Validator Setup
	config.RegisterValidator(e)

	// Products setup
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	// User setup
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "🚀 Server running and DB connected!")
	})

	productGroup := e.Group("/products")
	productGroup.Use(jwtmiddleware.JWTMiddleware()) // apply jwt middleware

	productGroup.POST("", productHandler.Create)
	productGroup.GET("", productHandler.FindAll)
	productGroup.GET("/:id", productHandler.FindByID)
	productGroup.PUT("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)
	productGroup.GET("/export", productHandler.ExportCSV)

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	defer db.Close()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(e.Start(":8080"))
}
