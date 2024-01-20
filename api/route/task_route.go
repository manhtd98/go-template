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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
	group.GET("/update", tc.Update)
}
