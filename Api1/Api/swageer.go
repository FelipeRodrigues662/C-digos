package api

import (
	_ "github.com/FelipeRodrigues662/ApiRestGo/api/docs"
	"github.com/gorilla/mux"
)

// @title My API
// @description This is a sample CRUD API with Swagger integration
// @version 1
// @host localhost:8000
// @BasePath /api/v1
func RegisterSwagger(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition"
	))
}
