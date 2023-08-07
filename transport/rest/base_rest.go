package rest

import (
	"net/http"
	"template/contract"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type templateRestServer struct{}

func getTemplateRestServer() contract.TemplateServiceRestInterface {
	return &templateRestServer{}
}

func NewTemplateRestServer() *http.Server {
	mux := chi.NewMux()
	mux.Use(middleware.Logger)

	server := getTemplateRestServer()
	mapRouteToHandler(mux, server)

	return &http.Server{
		Handler: mux,
	}
}

func mapRouteToHandler(mux *chi.Mux, server contract.TemplateServiceRestInterface) {
	mux.Post("/add", server.Add)
}
