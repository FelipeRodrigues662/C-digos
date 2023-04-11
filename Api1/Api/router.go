package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) {
	router := mux.NewRouter()
	userHandler := NewUserHandler(db)
	RegisterHandlers(router, userHandler)
	RegisterSwagger(router)

	// Adicionar tratamento para rota não encontrada (404)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Endpoint não encontrado.")
	})

	// Iniciar o servidor HTTP
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
