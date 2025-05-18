package main

import (
	_ "github.com/lib/pq"
	"log"
	"travel-agency-seeder/internal"
	"travel-agency-seeder/internal/monitoring"
)

func main() {
	db := connect.ToDb()
	log.Printf("⏳ Starting monitoring...")
	monitoring.Start(db)
	log.Printf("✅ Monitoring started")

	select {}
}
