package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

// Function to connect to the database
func Connect() error {
	var err error

	for {
		DB, err = pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))

		if err != nil {
			log.Println(err)
			time.Sleep(15 * time.Second)
		} else {
			log.Println("Successfully connected to the database!")
			break
		}
	}

	return nil
}
