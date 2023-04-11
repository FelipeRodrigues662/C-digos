package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection representa a conexão com o banco de dados
type Connection struct {
	DB *gorm.DB
}

// Connect realiza a conexão com o banco de dados
func Connect(dsn string) (*Connection, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Migrate as tabelas
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
	}

	return &Connection{DB: db}, nil
}
