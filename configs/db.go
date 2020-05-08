package configs

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

// Connecting to db
func Connect(user string, password string, addr string, name string) *pg.DB {
	opts := &pg.Options{
		User:     user,
		Password: password,
		Addr:     addr,
		Database: name,
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	return db
}
