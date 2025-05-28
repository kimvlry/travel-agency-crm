package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"travel-agency-seeder/db-load-simulation/pkg"
	connect "travel-agency-seeder/shared"
)

func main() {
	db := connect.ToDb()
	log.Printf("⏳ Starting db-load simulation...")
	simulation.Start(db)
	log.Printf("✅ simulation started")

	simulation.SetupMetrics()
	log.Printf("✅ Metrics server starting on :8080/metrics")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("❌ Metrics server error: %v\n", err)
	}

	select {}
}
