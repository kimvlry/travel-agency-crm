package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"travel-agency-seeder/internal/seeder"
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

func ensureSeedHistoryTable(db *gorm.DB) {
	err := db.Exec(`
		CREATE TABLE IF NOT EXISTS seed_history (
			id SERIAL PRIMARY KEY,
			version VARCHAR(10) NOT NULL UNIQUE,
			seeded_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);
	`).Error
	if err != nil {
		log.Fatalf("failed to create seed_history table: %v", err)
	}
}

func checkAppliedSeedVersions(db *gorm.DB) map[string]bool {
	var versions []string
	result := db.Table("seed_history").Pluck("version", &versions)
	if result.Error != nil {
		log.Fatalf("error reading seed_history: %v", result.Error)
	}

	applied := make(map[string]bool)
	for _, v := range versions {
		applied[v] = true
	}
	return applied
}

func checkAppliedFlywayVersions(db *gorm.DB) map[string]bool {
	var versions []string
	result := db.
		Table("flyway_schema_history").
		Where("success = true").
		Pluck("version", &versions)
	if result.Error != nil {
		log.Fatalf("error reading flyway_schema_history: %v", result.Error)
	}

	applied := make(map[string]bool)
	for _, v := range versions {
		applied[v] = true
	}
	return applied
}

func markSeedingApplied(db *gorm.DB, version string) {
	err := db.Exec(
		"INSERT INTO seed_history (version) VALUES (?) ON CONFLICT DO NOTHING", version,
	).Error
	if err != nil {
		log.Fatalf("failed to mark seed as applied: %v", err)
	}
}

type Seeder struct {
	Version string
	Run     func()
}

func main() {
	db := connectToDb()
	ensureSeedHistoryTable(db)

	seedCount, errSeed := strconv.Atoi(os.Getenv("SEED_COUNT"))
	if errSeed != nil {
		log.Fatal("invalid SEED_COUNT env ", errSeed)
	}

	seeders := []Seeder{
		{"2", func() { seeder.NewV2Seeder(db, seedCount).Seed() }},
		{"3", func() { seeder.NewV3Seeder(db, seedCount).Seed() }},
		{"4", func() { seeder.NewV4Seeder(db, seedCount).Seed() }},
		{"5", func() { seeder.NewV5Seeder(db, seedCount).Seed() }},
	}

	appliedMigrations := checkAppliedFlywayVersions(db)
	appliedSeedings := checkAppliedSeedVersions(db)

	for _, s := range seeders {
		if !appliedMigrations[s.Version] {
			log.Printf("Skipping version %s (migration not applied)\n", s.Version)
			continue
		}
		if appliedSeedings[s.Version] {
			log.Printf("Skipping version %s (already seeded)\n", s.Version)
			continue
		}

		log.Printf("Seeding for version %s\n", s.Version)
		err := func() error {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Seeding panicked for version %s: %v\n", s.Version, r)
				}
			}()
			s.Run()
			return nil
		}()

		if err != nil {
			log.Printf("Error seeding for version %s: %v\n", s.Version, err)
			continue
		}
		markSeedingApplied(db, s.Version)
	}
}
