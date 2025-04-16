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

func markSeedingApplied(tx *gorm.DB, version string) error {
	return tx.Exec(
		"INSERT INTO seed_history (version) VALUES (?) ON CONFLICT DO NOTHING", version,
	).Error
}

func main() {
	db := connectToDb()
	ensureSeedHistoryTable(db)

	seedCount, errSeed := strconv.Atoi(os.Getenv("SEED_COUNT"))
	if errSeed != nil {
		log.Fatal("invalid SEED_COUNT env ", errSeed)
	}

	appliedMigrations := checkAppliedFlywayVersions(db)
	appliedSeedings := checkAppliedSeedVersions(db)

	var seeders = map[string]func(*gorm.DB, int) error{
		"1": func(db *gorm.DB, count int) error {
			return seeder.NewV1Seeder().Seed()
		},
		"2": func(db *gorm.DB, count int) error {
			return seeder.NewV2Seeder(db, count).Seed()
		},
		"3": func(db *gorm.DB, count int) error {
			return seeder.NewV3Seeder(db, count).Seed()
		},
		"4": func(db *gorm.DB, count int) error {
			return seeder.NewV4Seeder(db, count).Seed()
		},
		"5": func(db *gorm.DB, count int) error {
			return seeder.NewV5Seeder(db, count).Seed()
		},
		"6": func(db *gorm.DB, count int) error {
			return seeder.NewV6Seeder(db).Seed()
		},
	}

	for version, seedFunc := range seeders {
		if !appliedMigrations[version] {
			log.Printf("skipping version %s (migration not applied)\n", version)
			continue
		}
		if appliedSeedings[version] {
			log.Printf("skipping version %s (already seeded)\n", version)
			continue
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			seedErr := seedFunc(tx, seedCount)
			if seedErr != nil {
				return fmt.Errorf("seeding error: %w", seedErr)
			}
			if err := markSeedingApplied(tx, version); err != nil {
				return fmt.Errorf("failed to mark version %s as seeded: %w", version, err)
			}
			return nil
		})
		if err != nil {
			log.Printf("transaction for version %s failed: %v\n", version, err)
		} else {
			log.Printf("successfully seeded version %s\n", version)
		}
	}
}
