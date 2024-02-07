package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/project/go-microservices/domain"
)

type NewController struct {
	NewsUsecase domain.NewsUsecase
}

func (tc *NewController) Create(c *gin.Context) {
	var news domain.News

	err := c.ShouldBind(&news)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	news.ID = primitive.NewObjectID()

	news.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.NewsUsecase.Create(c, &news)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *NewController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	news, err := u.NewsUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (u *NewController) Delete(c *gin.Context) {
	userID := c.GetString("x-user-id")

	news, err := u.NewsUsecase.DeleteByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}
