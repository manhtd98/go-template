package main

import (
	"time"

	"github.com/gin-gonic/gin"

	route "github.com/project/go-microservices/api/route"
	"github.com/project/go-microservices/bootstrap"
	"github.com/project/go-microservices/db"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	// defer db.PGDataBase.Close()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db.PGDataBase, gin)

	gin.Run(env.ServerAddress)
}
