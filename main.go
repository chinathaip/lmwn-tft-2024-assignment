package main

import (
	"flag"

	"github.com/chinathaip/lmwn-tft-2024-assignment/presentation/router"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "port for the http server")

	r := router.NewGinRouter()

	r.Start(port)
}
