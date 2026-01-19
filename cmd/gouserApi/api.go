package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/spudmashmedia/gouser/internal/config"
	"github.com/spudmashmedia/gouser/internal/health"
	"github.com/spudmashmedia/gouser/internal/users"
	"github.com/spudmashmedia/gouser/pkg/randomuser"
)

type application struct {
	config *config.ApiConfig
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Health Endpoint
	healthHandler := health.NewHandler(nil)
	r.Get("/health", healthHandler.GetHealth)

	// Users Endpoint
	usersService := users.NewService(
		randomuser.NewService(
			app.config.ExtRandomuser.Host,
			app.config.ExtRandomuser.Route,
		),
	)

	usersHandler := users.NewHandler(usersService)

	r.Route("/user", func(r chi.Router) {
		r.Use(users.UserCtx)
		r.Get("/", usersHandler.GetUser)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.GouserApi.Addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	slog.Info(
		"Server started at ",
		"app.config.addr", app.config.GouserApi.Addr)

	return srv.ListenAndServe()
}
