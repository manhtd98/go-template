package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/project/go-microservices/api/controller"
	"github.com/project/go-microservices/bootstrap"
	"github.com/project/go-microservices/domain"
	"github.com/project/go-microservices/db"
	"github.com/project/go-microservices/repository"
	"github.com/project/go-microservices/usecase"
)

func NewsRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	tr := repository.NewNewsRepository(db, domain.CollectionNew)
	tc := &controller.NewController{
		NewsUsecase: usecase.NewNewsUsecase(tr, timeout),
	}
	group.GET("/news", tc.Fetch)
	group.POST("/news", tc.Create)
	group.DELETE("/news", tc.Delete)
}
