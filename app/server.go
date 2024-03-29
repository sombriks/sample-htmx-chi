package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"sample-htmx-chi/app/controllers"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	var server Server
	server.router = chi.NewRouter()
	server.router.Use(middleware.Logger)
	server.router.Get("/htmx.js", controllers.HtmxHandler)
	server.router.Route("/", func(r chi.Router) {
		r.Get("/", controllers.IndexHandler)
	})
	return &server
}

func (server *Server) Start() error {
	port := 4000
	log.Printf("starting server on http://localhost:%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), server.router)
}
