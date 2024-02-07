package bootstrap

import "github.com/project/go-microservices/db"

type Application struct {
	Env   *Env
	Mongo db.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
