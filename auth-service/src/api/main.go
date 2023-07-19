package main

import (
	data "auth/db"
	"database/sql"
	"log"
)

const PORT = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service...")
}
