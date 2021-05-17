package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rizkyrmsyah/gin-gonic/cmd/database"
	"github.com/rizkyrmsyah/gin-gonic/config"
	"github.com/rizkyrmsyah/gin-gonic/repository/mysql"
	transportHTTP "github.com/rizkyrmsyah/gin-gonic/transport/http"
	"github.com/rizkyrmsyah/gin-gonic/usecase"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewConnection(config.DB)

	r := gin.New()
	r.Use(logger.SetLogger())

	defer db.Close()

	user := mysql.NewUserRepository(db)
	userUseCase := usecase.NewUser(user)

	transportHTTP.NewUserHTTPHandler(r, userUseCase)
	r.Run(":" + strconv.Itoa(config.AppPort))
	// log.Info().Msg("Service Running at port : " + port)
}
