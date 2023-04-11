package api

import (
	"net/http"

	"github.com/FelipeRodrigues662/ApiRestGo/api/swagger"
	"github.com/FelipeRodrigues662/ApiRestGo/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) {
	router := mux.NewRouter()
	handlers.RegisterHandlers(router, db)
	swagger.RegisterSwagger(router)
	http.ListenAndServe(":8000", router)
}
