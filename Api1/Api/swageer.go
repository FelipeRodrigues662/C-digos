package api

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterSwagger configura o Swagger na rota /swagger/
func RegisterSwagger(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // alteração da URL para uma rota relativa
	))
}
