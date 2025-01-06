package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "youthcongressnepal-6418.j77.cockroachlabs.cloud"
	port     = 26257              
	user     = "chetanbudathoki"   
	password = "0XDNmjBAseohph82xkl4GQ"
	dbname   = "youthcongressnepal"
	sslmode  = "require"
)

func Connection() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	
	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}

	log.Println("Connected to the database 'youthcongressnepal' ")	
}