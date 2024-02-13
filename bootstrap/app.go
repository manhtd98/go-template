package bootstrap

import (
	"gorm.io/gorm"

	"github.com/project/go-microservices/db"
	"github.com/project/go-microservices/domain"
)

type Application struct {
	Env      *Env
	Postgres *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()

	dbConfig := domain.DBConfig{
		DBHost:         app.Env.DBHost,
		DBUserName:     app.Env.DBUser,
		DBName:         app.Env.DBName,
		DBUserPassword: app.Env.DBPass,
		DBPort:         app.Env.DBPort,
	}
	newPGDB := db.NewPGDatabase(dbConfig)
	NewPGDatabase(newPGDB, app.Env)
	app.Postgres = newPGDB.DBConnect
	return *app
}
