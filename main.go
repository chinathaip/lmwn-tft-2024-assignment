package main

import (
	"flag"

	repository "github.com/chinathaip/lmwn-tft-2024-assignment/data/repository_impl"
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/controller"
	"github.com/chinathaip/lmwn-tft-2024-assignment/presentation/router"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "port for the http server")

	covidRepository := repository.NewCovidRepositoryImpl()
	covidController := controller.NewCovidController(covidRepository)
	r := router.NewGinRouter(covidController)

	r.Start(port)
}
