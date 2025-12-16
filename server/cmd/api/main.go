package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/AmiyoKm/todo-micro/configs"
	"github.com/AmiyoKm/todo-micro/internal/infra/cache"
	"github.com/AmiyoKm/todo-micro/internal/infra/db"
	"github.com/AmiyoKm/todo-micro/internal/repository"
	"github.com/AmiyoKm/todo-micro/internal/service"
	"github.com/AmiyoKm/todo-micro/utils"
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
		Name:        utils.GetEnv("POSTGRES_DB", "todo_micro"),
		User:        utils.GetEnv("POSTGRES_USER", "postgres"),
		Password:    utils.GetEnv("POSTGRES_PASSWORD", "password"),
		SslMode:     utils.GetEnv("POSTGRES_SSLMODE", "disable"),
		MaxConnOpen: utils.GetEnvInt("POSTGRES_MAX_CONN_OPEN", 60),
		MaxIdleConn: utils.GetEnvInt("POSTGRES_MAX_IDLE_CONN", 60),
		MaxIdleTime: utils.GetEnv("POSTGRES_MAX_IDLE_TIME", "15m"),
	}

	redisConfig := configs.RedisConfig{
		Host:       utils.GetEnv("REDIS_HOST", "localhost"),
		Port:       utils.GetEnv("REDIS_PORT", "6379"),
		Password:   utils.GetEnv("REDIS_PASSWORD", ""),
		DB:         utils.GetEnvInt("REDIS_DB_NUMBER", 0),
		Expiration: time.Duration(utils.GetEnvInt("REDIS_EXPIRATION", 300)) * time.Second,
	}

	config := configs.Config{
		Port:  utils.GetEnvInt("PORT", 8080),
		Env:   utils.GetEnv("ENVIRONMENT", "DEVELOPMENT"),
		Db:    dbConfig,
		Redis: redisConfig,
	}

	dbConn, err := db.New(&config.Db)
	if err != nil {
		log.Fatal("error while connecting to database, err: ", err)
		return
	}
	defer dbConn.Close()
	log.Println("DB connection pool established")

	redisClient, err := cache.New(&config.Redis)
	if err != nil {
		log.Fatal("error while connecting to redis, err: ", err)
		return
	}
	defer redisClient.Close()
	log.Println("Redis connection pool established")

	repo := repository.NewRepo(dbConn)
	cacheRepo := cache.NewCacheRepository(redisClient)
	svc := service.NewService(repo, cacheRepo)
	server(config, svc)
}
