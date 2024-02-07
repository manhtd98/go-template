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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
	group.GET("/update", tc.Update)
}
