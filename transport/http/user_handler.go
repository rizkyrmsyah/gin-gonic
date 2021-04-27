package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	useCaseI "github.com/rizkyrmsyah/gin-gonic/interfc"
)

type HTTPUser struct {
	useCaseInterface useCaseI.UserInterface
}

func NewUserHTTPHandler(r *gin.Engine, useCaseInterface useCaseI.UserUseCaseI) {
	handler := &HTTPUser{useCaseInterface}

	api := r.Group("/api/user")
	{
		api.GET("/list", handler.GetAll)
	}
}

func (handler *HTTPUser) GetAll(c *gin.Context) {
	list, err := handler.useCaseInterface.FindAll()
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"responseCode":    "50",
				"responseMessage": "Terjadi kesalahan pada sistem",
			},
		)
		return
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
