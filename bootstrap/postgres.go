package bootstrap

import (
	"context"
	"time"

	"github.com/project/go-microservices/db"
)

func NewPGDatabase(env *Env) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConfig := db.Config{
		DBHost: env.DBHost,
		DBUserName: env.DBUser,
		DBName: env.DBName,
		DBUserPassword: env.DBPass,
		DBPort: env.DBPort,
	}
	db.ConnectDB(&dbConfig)
}
