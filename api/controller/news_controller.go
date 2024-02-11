package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

	news.UserID = userID
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
