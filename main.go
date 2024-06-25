package main

import (
	"flag"
	"log"

	repository "github.com/chinathaip/lmwn-tft-2024-assignment/data/repository_impl"
	"github.com/chinathaip/lmwn-tft-2024-assignment/domain/controller"
	"github.com/chinathaip/lmwn-tft-2024-assignment/presentation/router"
)

const dataSource = "https://static.wongnai.com/devinterview/covid-cases.json"

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "port for the http server")

	covidDatasource, err := repository.FetchCovidCase(dataSource)
	if err != nil {
		log.Fatal(err)
	}

	covidRepository := repository.NewCovidRepositoryImpl(covidDatasource)
	covidController := controller.NewCovidController(covidRepository)
	r := router.NewGinRouter(covidController)

	r.Start(port)
}
