package api

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

func NewApplication(config *config.ApiConfig) *application {
	return &application{
		config: config,
	}
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Prepare Users Service for DI
	usersService := users.NewService(
		randomuser.NewService(
			app.config.ExtRandomuser.Host,
			app.config.ExtRandomuser.Route,
		),
	)

	// Register Chi Rest Routes
	RegisterHealthRouter(r)
	RegisterUserRouter(r, usersService)

	return r
}

func RegisterHealthRouter(r *chi.Mux) {
	healthHandler := health.NewHandler(nil)
	r.Get("/health", healthHandler.GetHealth)
}

func RegisterUserRouter(r *chi.Mux, svc users.Service) {
	usersHandler := users.NewHandler(svc)
	r.Route("/user", func(r chi.Router) {
		r.Use(users.UserCtx)
		r.Get("/", usersHandler.GetUser)
	})
}

func (app *application) Run(h http.Handler) error {
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
