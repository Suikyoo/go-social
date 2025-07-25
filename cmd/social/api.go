package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/go-chi/cors"
)

type config struct{
  addr string
}
type application struct {
  config config
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
    r.Post("/", createPostHandler)

    r.Get("/", getPostFeedHandler)
    r.Get("/{id}", getPostHandler)

  })

  r.Route("/users", func(r chi.Router) {

    r.Get("/", getUserFeedHandler)
    r.Get("/{id}", getUserHandler)

  })

  r.Route("/auth", func(r chi.Router) {
    r.Post("/user", registerUserHandler)
    r.Post("/token", createTokenHandler)

  })

  return r
}

func (app *application) Run(handler http.Handler) {
  srv := http.Server{
    Addr: app.config.addr,
    Handler: handler,
  }

  log.Fatal(srv.ListenAndServe())
}
