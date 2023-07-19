package main

import (
	data "auth/db"
	"database/sql"
	"log"
	"net/http"
	"fmt"
)

const PORT = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service...")

	//conection to db


	//config
	app := Config{}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", PORT),
		Handler: app.routes()
	}


}
