package main

import (
	"log"

	api "new.com/events/Api"
	database "new.com/events/Database"
)

func main() {
	// Inicialização da conexão com o banco de dados
	dsn := "sqlserver://sa:543219876Dk@localhost:1433?database=Api"
	db, err := database.InitDatabase(dsn)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Inicialização do roteador da API
	api.StartRouter(db.DB)
}
