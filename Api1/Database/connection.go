package database

import (
	"fmt"

	models "new.com/events/Models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Connection representa a conexão com o banco de dados
type Connection struct {
	DB *gorm.DB
}

// InitDatabase realiza a conexão com o banco de dados e executa as migrações
func InitDatabase(dsn string) (*Connection, error) {
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Executar as migrações
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("error migrating tables: %w", err)
	}

	return &Connection{DB: db}, nil
}
