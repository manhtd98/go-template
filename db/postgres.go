package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/project/go-microservices/domain"
)

type PGDatabase struct {
	DBConnect *gorm.DB
	config    domain.DBConfig
}

func NewPGDatabase(cfg domain.DBConfig) *PGDatabase {
	return &PGDatabase{config: cfg}
}
func (ur *PGDatabase) ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", ur.config.DBHost,
		ur.config.DBUserName, ur.config.DBUserPassword,
		ur.config.DBName, ur.config.DBPort)

	ur.DBConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
