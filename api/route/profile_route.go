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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
