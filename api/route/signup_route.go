package route

import (
	"time"

	"github.com/project/go-microservices/api/controller"
	"github.com/project/go-microservices/bootstrap"
	"github.com/project/go-microservices/domain"
	"github.com/project/go-microservices/db"
	"github.com/project/go-microservices/repository"
	"github.com/project/go-microservices/usecase"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
