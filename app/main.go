package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация бд
	config.Init()
	cfg := config.LoadConfig()
	db, err := config.InitDB(cfg)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Инициализация сервисов
	userRepository := repository.NewUserRepository(db)
	jwtService := service.NewJWTService(cfg.JWTSecret, 144000000)
	userService := service.NewUserService(userRepository, jwtService)
	authHandler := handler.NewAuthHandler(userService)

	// Инициализация роутера Gin
	r := gin.Default()

	// Настройка middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Группа маршрутов API
	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
