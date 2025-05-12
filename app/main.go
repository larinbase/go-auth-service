package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
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
	jwtService := service.NewJWTService(cfg.JWTSecret, 7200000)
	refreshSessionRepository := repository.NewRefreshSessionRepository(db)
	emailService := service.NewEmailService(cfg.EmailApiUsername, cfg.EmailApiPassword)
	emailManager := service.NewEmailManager(10, emailService)
	userService := service.NewUserService(userRepository, jwtService, refreshSessionRepository, emailManager)
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandler(userService)

	// Инициализация роутера Gin
	r := gin.Default()

	// Настройка
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	// Группа маршрутов API
	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-tokens", authHandler.RefreshTokens)
		auth.POST("/v2/login", authHandler.LoginV2)
		auth.POST("/v2/sendCode", authHandler.SendCode)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.PATCH("/change-password", userHandler.ChangePassword)
	}

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
