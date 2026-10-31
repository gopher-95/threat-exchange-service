package main

import (
	"context"
	"log"

	"github.com/gopher-95/threat-exchange-service/internal/config"
	"github.com/gopher-95/threat-exchange-service/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("ошибка загрузки конфига: %v", err)
	}

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("ошибка подключения к бд: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("бд недоступна: %v", err)
	}

	_ = repo.NewDB(pool)
}
