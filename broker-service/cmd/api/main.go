package main

import (
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {}

func main () {
	app := Config{}

	log.Println("Starting service on port %s", webPort)

	// define http server

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s")
	}


}