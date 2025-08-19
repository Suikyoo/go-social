package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
  "github.com/Suikyoo/go-social/internal/repository"
)

type config struct {
  server serverConfig
  auth authConfig
}

type serverConfig struct {
  addr string
}

type authConfig struct {
  jwtSecret []byte
  tokenExpiry int64
}

type application struct {
	config config
  store *repository.Storage
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Route("/posts", func(r chi.Router) {
    r.Use(app.AuthTokenMiddleware)
		r.Post("/", app.createPost)

		r.Get("/", app.getPostFeed)
		r.Get("/{id}", app.getPost)

	})

	r.Route("/users", func(r chi.Router) {
    r.Use(app.AuthTokenMiddleware)
		r.Get("/", app.getUserFeed)
		r.Get("/{id}", app.getUser)

	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/user", app.createUser)
		r.Post("/token", app.createToken)

	})

	return r
}

func (app *application) Run(handler http.Handler) {
	srv := http.Server{
		Addr:    app.config.server.addr,
		Handler: handler,
	}

  log.Printf("Server running at %s", srv.Addr)
  log.Printf("code just got updated")
	log.Fatal(srv.ListenAndServe())
}
