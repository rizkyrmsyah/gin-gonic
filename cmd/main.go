package main

import (
	"log"

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

	// httpAddr := flag.String("http.addr", config.AppPort, "HTTP listen address")

	db := database.NewConnection(config.DB)

	r := gin.New()

	defer db.Close()

	user := mysql.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(user)

	transportHTTP.NewUserHTTPHandler(r, userUseCase)
	r.Run(":3001")
	// log.Info().Msg("Service Running at port : " + port)
}
