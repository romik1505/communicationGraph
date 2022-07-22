package main

import (
	"log"

	"github.com/romik1505/communicationGraph/internal/app/config"
	"github.com/romik1505/communicationGraph/internal/app/service"
	"github.com/romik1505/communicationGraph/internal/server"
)

// @title           Communication Graph
// @version         1.0
// @description     This is user communication service.

// @host      localhost:8080
// @BasePath  /
func main() {
	s := service.NewCommunicationService(config.NewPostgresConnection())
	h := server.NewHandler(s)

	app := server.NewApp(h.InitRoutes())
	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
