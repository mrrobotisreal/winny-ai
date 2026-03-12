package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"io.winapps/internal/api"
	"io.winapps/internal/config"
)

func main() {
	log.SetFlags(0)

	cfg := config.Load()
	app, err := api.NewApp(context.Background(), cfg)
	if err != nil {
		log.Fatalf("init app: %v", err)
	}

	log.Printf("winny assistant api listening on :%s", cfg.Port)
	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           app.Router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       cfg.HTTPReadTimeout,
		WriteTimeout:      cfg.HTTPWriteTimeout,
		IdleTimeout:       cfg.HTTPIdleTimeout,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("run server: %v", err)
	}
}
