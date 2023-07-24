package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:password@localhost:5432/users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connecting to database!")
	defer conn.Close(context.Background())

	file, err := os.ReadFile("users.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open SQL file: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(context.Background(), string(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute SQL command: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("SQL file has been successfully loaded into the database.")
}
