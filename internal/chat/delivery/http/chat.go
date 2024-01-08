package http

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Chat struct {
	router chi.Router
}

func NewChat() *Chat {
	h := &Chat{
		router: chi.NewRouter(),
	}

	h.router.Get("/", h.Index)

	return h
}

func (c *Chat) Handler() http.Handler {
	return c.router
}

func (c *Chat) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/template/index.html")
	if err != nil {
		slog.Error("failed to parse template file: %w", err)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		slog.Error("failed to execute template: %w", err)
		return
	}
}
