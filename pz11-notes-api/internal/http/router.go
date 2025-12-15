package httpx

import (
	"example.com/notes-api/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/v1/notes", h.CreateNote)
	r.Get("/api/v1/notes", h.List)
	r.Get("/api/v1/notes/{id}", h.GetNote)
	r.Patch("/api/v1/notes/{id}", h.EditNote)
	r.Delete("/api/v1/notes/{id}", h.DeleteNote)
	return r
}