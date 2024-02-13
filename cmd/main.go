package main

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	route "github.com/project/go-microservices/api/route"
	"github.com/project/go-microservices/bootstrap"
)

func LoggingInit(logEnv string) {
	// load environment variable
	// setup logrus
	logLevel, err := log.ParseLevel(logEnv)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	app := bootstrap.App()

	env := app.Env
	LoggingInit(env.LoggingLevel)
	// defer db.PGDataBase.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, app.Postgres, gin)

	gin.Run(env.ServerAddress)
}
