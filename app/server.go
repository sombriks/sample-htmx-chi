package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"sample-htmx-chi/app/configs"
	"sample-htmx-chi/app/controllers"
	"sample-htmx-chi/app/services"
)

type Server struct {
	port       string
	router     *chi.Mux
	config     *configs.Config
	service    *services.TodoService
	controller *controllers.TodoController
}

func NewServer() *Server {
	var server Server

	var set bool
	server.port, set = os.LookupEnv("PORT")
	if !set {
		server.port = "4000"
	}

	server.config = configs.NewConfig()
	server.service = services.NewTodoService(server.config)
	server.controller = controllers.NewTodoController(server.config, server.service)

	server.router = chi.NewRouter()
	server.router.Use(middleware.Logger)
	server.router.Get("/htmx-1.9.11.min.js", server.controller.AssetsHandler)
	server.router.Route("/", func(r chi.Router) {
		r.Get("/", server.controller.IndexHandler)
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", server.controller.ListHandler)
			r.Post("/", server.controller.InsertHandler)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", server.controller.FindHandler)
				r.Put("/", server.controller.UpdateHandler)
				r.Delete("/", server.controller.DeleteHandler)
			})
		})
	})

	// we're done configuring the server
	return &server
}

func (server *Server) Start() error {
	log.Printf("starting server on http://localhost:%s\n", server.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", server.port), server.router)
}
