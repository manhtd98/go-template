package bootstrap

import (
	"gorm.io/gorm"

	"github.com/project/go-microservices/db"
)

type Application struct {
	Env *Env
	// Mongo db.Client
	Postgres *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	NewPGDatabase(app.Env)
	app.Postgres = db.PGDataBase
	return *app
}