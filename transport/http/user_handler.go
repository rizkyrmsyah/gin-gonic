package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizkyrmsyah/gin-gonic/model"
	"github.com/rizkyrmsyah/gin-gonic/usecase"
)

type HTTPUser struct {
	useCaseInterface usecase.UserUseCaseI
}

func NewUserHTTPHandler(r *gin.Engine, useCaseInterface usecase.UserUseCaseI) {
	handler := &HTTPUser{useCaseInterface}

	api := r.Group("/api/user")
	{
		api.GET("/list", handler.GetAll)
		api.POST("/store", handler.Store)
	}
}

func (handler *HTTPUser) GetAll(c *gin.Context) {
	list, err := handler.useCaseInterface.GetAll()
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    "00",
			"message": "success",
			"result":  &list,
		},
	)
}

func (handler *HTTPUser) Store(c *gin.Context) {
	var user model.StoreUserRequest
	c.BindJSON(&user)

	err := handler.useCaseInterface.Store(&user)
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"code":    "00",
			"message": "user created",
		},
	)
}
