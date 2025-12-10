package router

import (
	"log/slog"
	"net/http"

	mw "flexsupport/internal/middleware"
	"flexsupport/static"

	// "net/http"
	"flexsupport/internal/routes/api"
	"flexsupport/internal/routes/dashboard"
	"flexsupport/internal/routes/tickets"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var AppDev string = "development"

func NewRouter(log *slog.Logger) *chi.Mux {
	r := chi.NewMux()
	// Dashboard

	r.Handle("/assets/*",
		disableCacheInDevMode(
			http.StripPrefix("/assets/",
				static.AssetRouter()),
		),
	)
	// Tickets
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			middleware.Recoverer,
			// mw.CSPMiddleware,
			mw.Logging(log),
			mw.TextHTMLMiddleware,
		)
		dashboard.Mount(r, dashboard.NewHandler(log, dashboard.NewService(log)))
		tickets.Mount(r, tickets.NewHandler(log, tickets.NewService(log)))
	})
	api.Mount(r, api.NewHandler(log, api.NewService(log)))

	return r
}

func disableCacheInDevMode(next http.Handler) http.Handler {
	if AppDev == "development" {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
