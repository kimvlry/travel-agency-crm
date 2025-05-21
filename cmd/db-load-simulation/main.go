package main

import (
	_ "github.com/lib/pq"
	"log"
	"travel-agency-seeder/internal"
	"travel-agency-seeder/internal/db-load-simulation"
)

func main() {
	db := connect.ToDb()
	log.Printf("⏳ Starting db-load simulation...")
	db_load_simulation.Start(db)
	log.Printf("✅ simulation started")

	select {}
}
