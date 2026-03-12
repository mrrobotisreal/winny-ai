package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"io.winapps/internal/catalog"
	"io.winapps/internal/domain"
	"io.winapps/internal/service"
)

type ChatService interface {
	StreamChat(context.Context, domain.ChatRequest, func(domain.StreamEvent) error) error
	Reset(context.Context) error
}

type Handler struct {
	chatService ChatService
}

func NewHandler(chatService ChatService) *Handler {
	return &Handler{chatService: chatService}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeMethodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"assistant": catalog.AssistantName,
	})
}

func (h *Handler) StreamChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}

	var request domain.ChatRequest
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "streaming not supported"})
		return
	}

	w.Header().Set("Content-Type", "application/x-ndjson; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	encoder := json.NewEncoder(w)
	err := h.chatService.StreamChat(r.Context(), request, func(event domain.StreamEvent) error {
		if err := encoder.Encode(event); err != nil {
			return err
		}
		flusher.Flush()
		return nil
	})
	if err == nil {
		return
	}

	if errors.Is(err, service.ErrInvalidRequest) {
		_ = encoder.Encode(domain.StreamEvent{
			Type:  "error",
			Error: err.Error(),
		})
		flusher.Flush()
		return
	}

	message := "something went wrong while talking to Winny"
	if strings.TrimSpace(err.Error()) != "" {
		message = err.Error()
	}
	_ = encoder.Encode(domain.StreamEvent{
		Type:  "error",
		Error: message,
	})
	flusher.Flush()
}

func (h *Handler) ResetChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeMethodNotAllowed(w)
		return
	}
	if err := h.chatService.Reset(r.Context()); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeMethodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}
