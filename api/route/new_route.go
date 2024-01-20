package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"backend-task/api/controller"
	"backend-task/bootstrap"
	"backend-task/domain"
	"backend-task/mongo"
	"backend-task/repository"
	"backend-task/usecase"
)

func NewsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewNewsRepository(db, domain.CollectionNew)
	tc := &controller.NewController{
		NewsUsecase: usecase.NewNewsUsecase(tr, timeout),
	}
	group.GET("/news", tc.Fetch)
	group.POST("/news", tc.Create)
	group.DELETE("/news", tc.Delete)
}
