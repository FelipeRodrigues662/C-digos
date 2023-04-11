package main

import (
	"log"
	"net/http"

	"github.com/FelipeRodrigues662/ApiRestGo/api"
	"github.com/FelipeRodrigues662/ApiRestGo/database"
)

func main() {
	// Inicialização da conexão com o banco de dados
	database.InitDatabase()

	// Inicialização do roteador da API
	router := api.NewRouter()

	// Inicialização do servidor HTTP
	log.Fatal(http.ListenAndServe(":8000", router))
}
