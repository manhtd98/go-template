package route

import (
	"time"

	"backend-task/api/controller"
	"backend-task/bootstrap"
	"backend-task/domain"
	"backend-task/mongo"
	"backend-task/repository"
	"backend-task/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
