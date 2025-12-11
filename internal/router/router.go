package router

import (
	"log/slog"
	"net/http"

	"flexsupport/internal/config"
	mw "flexsupport/internal/middleware"
	"flexsupport/static"

	// "net/http"
	"flexsupport/internal/routes/api"
	"flexsupport/internal/routes/dashboard"
	"flexsupport/internal/routes/tickets"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(log *slog.Logger, cfg *config.Config) *chi.Mux {
	r := chi.NewMux()
	// Dashboard

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Compress(5),
		)
		r.Handle("/assets/*",
			disableCacheInDevMode(
				http.StripPrefix("/assets/",
					static.AssetRouter(cfg)),
				cfg,
			),
		)
		r.Handle("/public/*",
			disableCacheInDevMode(
				http.StripPrefix("/public/",
					static.PublicRouter(cfg)),
				cfg,
			),
		)
	})
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

func disableCacheInDevMode(next http.Handler, cfg *config.Config) http.Handler {
	if cfg.Environment == config.PROD {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
