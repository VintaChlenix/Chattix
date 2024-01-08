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

	h.router.HandleFunc("/", h.Index)

	return h
}

func (c *Chat) Handler() http.Handler {
	return c.router
}

func (c *Chat) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/chat/templates/index.html")
	if err != nil {
		slog.Error("failed to parse templates file: %w", err)
		return
	}

	switch r.Method {
	case "GET":
		if err := tmpl.Execute(w, nil); err != nil {
			slog.Error("failed to execute templates: %w", err)
			return
		}
	case "POST":

	}

}
