package connect

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

func ToDb() *sqlx.DB {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbName == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("environment variables not set")
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort,
	)

	var db *sqlx.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			log.Printf("✅ connected to DB")
			return db
		}

		log.Printf("⏳ waiting for db... (%v)", err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("❌ couldn't connect to db after multiple attempts: %v", err)
	return nil
}
