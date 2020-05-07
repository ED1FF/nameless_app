package configs

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "",
		Addr:     "localhost:5432",
		Database: "nameless_app",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	return db
}
