package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"travel-agency-seeder/migrations_seeders"
)

func connectToDb() *gorm.DB {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	if dbName == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("environment variables not set")
	}

	connectionString := fmt.
		Sprintf("user=%s password=%s dbname=%s host=postgres port=5432 sslmode=disable",
			dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to db", err)
	}

	return db
}

func main() {
	db := connectToDb()

	seedCountStr := os.Getenv("SEED_COUNT")
	seedCount, err := strconv.Atoi(seedCountStr)
	if err != nil {
		log.Fatal("invalid seed count env", err)
	}

	v2 := migrations_seeders.NewV2Seeder(db, seedCount)
	v2.Seed()
}
