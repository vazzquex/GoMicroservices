package main

import (
	"fmt"
	"log"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	router := app.routes()

	//start server
	err := router.Run(fmt.Sprintf(":%s", webPort))
	if err != nil {
		log.Panic(err)
	}

}
