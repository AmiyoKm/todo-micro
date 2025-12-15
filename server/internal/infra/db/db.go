package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AmiyoKm/todo-micro/configs"
	_ "github.com/lib/pq"
)

func New(cfg *configs.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SslMode))

	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxConnOpen)
	db.SetMaxIdleConns(cfg.MaxIdleConn)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
