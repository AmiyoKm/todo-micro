package main

import (
	"log"
	"os"
	"sync"

	"github.com/AmiyoKm/user-micro/configs"
	"github.com/AmiyoKm/user-micro/internal/infra/db"
	"github.com/AmiyoKm/user-micro/internal/repository"
	"github.com/AmiyoKm/user-micro/internal/service"
	"github.com/AmiyoKm/user-micro/utils"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func main() {
	if _, err := os.Stat(".env"); err == nil {
		log.Print(".env file found")
		if err := godotenv.Load(); err != nil {
			log.Fatal(".env file not loaded")
		}
	}

	dbConfig := configs.DbConfig{
		Host:        utils.GetEnv("POSTGRES_HOST", "localhost"),
		Port:        utils.GetEnv("POSTGRES_PORT", "5432"),
		Name:        utils.GetEnv("POSTGRES_DB", "user_micro"),
		User:        utils.GetEnv("POSTGRES_USER", "postgres"),
		Password:    utils.GetEnv("POSTGRES_PASSWORD", "password"),
		SslMode:     utils.GetEnv("POSTGRES_SSLMODE", "disable"),
		MaxConnOpen: utils.GetEnvInt("POSTGRES_MAX_CONN_OPEN", 60),
		MaxIdleConn: utils.GetEnvInt("POSTGRES_MAX_IDLE_CONN", 60),
		MaxIdleTime: utils.GetEnv("POSTGRES_MAX_IDLE_TIME", "15m"),
	}

	jwtConfig := configs.JWTConfig{
		Secret:          utils.GetEnv("JWT_SECRET", "your-secret-key"),
		ExpirationHours: utils.GetEnvInt("JWT_EXPIRATION_HOURS", 24),
	}

	config := configs.Config{
		Port: utils.GetEnvInt("PORT", 8081),
		Env:  utils.GetEnv("ENVIRONMENT", "DEVELOPMENT"),
		Db:   dbConfig,
		JWT:  jwtConfig,
	}

	dbConn, err := db.New(&config.Db)
	if err != nil {
		log.Fatal("error while connecting to database, err: ", err)
		return
	}
	defer dbConn.Close()
	log.Println("DB connection pool established")

	repo := repository.NewRepo(dbConn)
	svc := service.NewService(repo, config.JWT.Secret, config.JWT.ExpirationHours)
	server(config, svc)
}
