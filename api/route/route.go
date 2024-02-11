package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/project/go-microservices/api/middleware"
	"github.com/project/go-microservices/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
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
}
