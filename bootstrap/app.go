package bootstrap

import (
	"gorm.io/gorm"

	"github.com/project/go-microservices/db"
)

type Application struct {
	Env *Env
	Postgres *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	newPGDB := db.NewPGDatabase()
	NewPGDatabase(&newPGDB, app.Env)
	app.Postgres = newPGDB.DBConnect
	return *app
}