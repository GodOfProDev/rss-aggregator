package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	env "github.com/godofprodev/rss-aggregator/internal/config"
	"log"
	"net/http"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Start() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", env.GetApiConfig().HandleCreateUser)

	router.Mount("/v1", v1Router)

	port := env.GetEnv().Port

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", port),
	}

	log.Printf("Server starting on port %v", port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
