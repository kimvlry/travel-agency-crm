package main

import (
	_ "github.com/lib/pq"
	"log"
	"travel-agency-seeder/db-load-simulation/pkg"
	connect "travel-agency-seeder/shared"
)

func main() {
	db := connect.ToDb()
	log.Printf("⏳ Starting db-load simulation...")
	simulation.Start(db)
	log.Printf("✅ simulation started")

	select {}
}
