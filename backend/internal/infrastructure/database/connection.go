package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/internal/domain/models"
	"backend/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewPostgreConnection(cfg *config.Config) (*Database, error) {
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Madrid",
	cfg.PostgreConfig.DB_HOST, 
	cfg.PostgreConfig.DB_PORT, 
	cfg.PostgreConfig.DB_USER, 
	cfg.PostgreConfig.DB_PASSWORD, 
	cfg.PostgreConfig.DB_NAME,
)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error open DB: %v", err)
	}

	// Migrations
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.WithContext(ctx).AutoMigrate(&models.Board{}, &models.Column{}, models.Task{})
	if err != nil {
		return nil, fmt.Errorf("error making migrations %v", err)
	}

	log.Println("✅ Conection succesfully to PostgreSQL!")
	return &Database{Conn: db}, nil
}
