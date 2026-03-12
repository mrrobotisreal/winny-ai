package api

import (
	"context"
	"net/http"
	"strings"

	"io.winapps/internal/config"
	"io.winapps/internal/runner"
	"io.winapps/internal/service"
)

type App struct {
	Router http.Handler
}

func NewApp(_ context.Context, cfg config.Config) (*App, error) {
	chatService := service.NewChatService(runner.NewOllamaCLIRunner(cfg), cfg)
	handler := NewHandler(chatService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", handler.Health)
	mux.HandleFunc("/api/v1/chat/stream", handler.StreamChat)
	mux.HandleFunc("/api/v1/chat/reset", handler.ResetChat)

	return &App{
		Router: corsMiddleware(cfg.AllowedOrigins, mux),
	}, nil
}

func corsMiddleware(allowedOrigins []string, next http.Handler) http.Handler {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed[origin] = struct{}{}
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := strings.TrimSpace(r.Header.Get("Origin"))
		if origin != "" {
			if _, ok := allowed[origin]; !ok {
				writeJSON(w, http.StatusForbidden, map[string]string{"error": "origin not allowed"})
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With, Cache-Control")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
