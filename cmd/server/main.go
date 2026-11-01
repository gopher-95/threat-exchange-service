package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gopher-95/threat-exchange-service/internal/config"
	"github.com/gopher-95/threat-exchange-service/internal/handler"
	"github.com/gopher-95/threat-exchange-service/internal/repo"
	"github.com/gopher-95/threat-exchange-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("failed connection to db: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("ping: %v", err)
	}

	db := repo.NewDB(pool)
	svc := service.NewService(db)
	h := handler.NewHandler(svc)

	router := newRouter(h)
	if err = http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func newRouter(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/info/{md5}", h.GetInfoHandle)
	r.Put("/info/{md5}", h.UpdateInfoHandle)
	r.Post("/info", h.CreateInfoHandle)

	return r
}
