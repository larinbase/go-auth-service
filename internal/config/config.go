package config

import (
	"auth-service/internal/domain"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	JWTSecret        string
	EmailApiKey      string
	EmailApiUsername string
	EmailApiPassword string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file"))
	}
}

func LoadConfig() *Config {
	return &Config{
		DBHost:           getEnv("DB_HOST"),
		DBPort:           getEnv("DB_PORT"),
		DBUser:           getEnv("DB_USER"),
		DBPassword:       getEnv("DB_PASSWORD"),
		DBName:           getEnv("DB_NAME"),
		JWTSecret:        getEnv("JWT_SECRET"),
		EmailApiUsername: getEnv("EMAIL_API_USERNAME"),
		EmailApiPassword: getEnv("EMAIL_API_PASSWORD"),
	}
}

func (c *Config) GetDBConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Errorf("failed to get enviroment %s", key))
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}

func InitDB(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.AutoMigrate(&domain.Role{}, &domain.User{}, &domain.RefreshSession{})

	// Создание дефолтных ролей
	var count int64
	db.Model(&domain.Role{}).Count(&count)
	if count == 0 {
		db.Create(&domain.Role{Name: domain.Author})
		db.Create(&domain.Role{Name: domain.Reviewer})
		db.Create(&domain.Role{Name: domain.Reader})
		db.Create(&domain.Role{Name: domain.Editor})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
