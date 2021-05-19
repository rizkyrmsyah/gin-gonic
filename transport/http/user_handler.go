package http

import (
	"net/http"
	"strconv"

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
		api.GET("/show/:id", handler.Show)
		api.DELETE("/delete/:id", handler.Delete)
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

func (handler *HTTPUser) Show(c *gin.Context) {
	id := c.Param("id")

	intId, _ := strconv.ParseInt(id, 10, 64)
	data, err := handler.useCaseInterface.Show(intId)
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    "00",
			"message": "success",
			"result":  &data,
		},
	)
}

func (handler *HTTPUser) Delete(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.ParseInt(id, 10, 64)
	err := handler.useCaseInterface.Delete(intId)
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    "00",
			"message": "success",
		},
	)
}
