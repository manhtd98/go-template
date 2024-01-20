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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
