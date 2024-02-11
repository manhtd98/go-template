package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/project/go-microservices/api/controller"
	"github.com/project/go-microservices/bootstrap"
	"github.com/project/go-microservices/repository"
	"github.com/project/go-microservices/usecase"
)

func NewsRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewNewsRepository(db)
	tc := &controller.NewController{
		NewsUsecase: usecase.NewNewsUsecase(tr, timeout),
	}
	group.POST("/news", tc.Create)
	group.GET("/news", tc.GetByID)
	group.DELETE("/news", tc.DeleteByUserID)
}
