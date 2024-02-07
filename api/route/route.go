package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/project/go-microservices/api/middleware"
	"github.com/project/go-microservices/bootstrap"
	"github.com/project/go-microservices/db"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db db.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	publicRouter.Use(middleware.CORSMiddleware())
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	protectedRouter.Use(middleware.CORSMiddleware())
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewsRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
