package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	}
}

func (handler *HTTPUser) GetAll(c *gin.Context) {
	list, err := handler.useCaseInterface.GetAll()
	if err != nil {
		panic(err)
	}

	if len(list) == 0 {
		c.JSON(
			http.StatusOK,
			gin.H{
				"responseCode":    "00",
				"responseMessage": "success",
				"listData":        &[]string{},
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"responseCode":    "00",
			"responseMessage": "success",
			"listData":        &list,
		},
	)
}
