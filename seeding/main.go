package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"travel-agency-seeder/seeding/pkg/impl"
	connect "travel-agency-seeder/shared"
)

func ensureSeedHistoryTable(db *sqlx.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS seed_history (
			id SERIAL PRIMARY KEY,
			version VARCHAR(10) NOT NULL UNIQUE,
			seeded_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);
	`)
	if err != nil {
		log.Fatalf("failed to create seed_history table: %v", err)
	}
}

func checkAppliedSeedVersions(db *sqlx.DB) map[string]bool {
	var versions []string
	err := db.Select(
		&versions,
		"SELECT version FROM seed_history",
	)
	if err != nil {
		log.Fatalf("error reading seed_history: %v", err)
	}

	applied := make(map[string]bool)
	for _, v := range versions {
		applied[v] = true
	}
	return applied
}

func checkAppliedFlywayVersions(db *sqlx.DB) map[string]bool {
	var versions []string
	err := db.Select(
		&versions,
		"SELECT version FROM flyway_schema_history WHERE success = true",
	)
	if err != nil {
		log.Fatalf("error reading flyway_schema_history: %v", err)
	}

	applied := make(map[string]bool)
	for _, v := range versions {
		applied[v] = true
	}
	return applied
}

func markSeedingApplied(tx *sqlx.Tx, version string) error {
	_, err := tx.Exec(
		`INSERT INTO seed_history (version) 
				VALUES ($1) 
				ON CONFLICT DO NOTHING
				`,
		version)
	return err
}

func validateSeedersMap(el *orderedmap.Element) (string, func(*sqlx.Tx, int) error, error) {
	version, ok := el.Key.(string)
	if !ok {
		return "", nil, fmt.Errorf("invalid key type for element: %v", el.Key)
	}

	seedFunc, ok := el.Value.(func(*sqlx.Tx, int) error)
	if !ok {
		return "", nil, fmt.Errorf("invalid type for element: %v", el.Value)
	}
	return version, seedFunc, nil
}

func shouldSkipVersion(version string, appliedMigrations map[string]bool, appliedSeedings map[string]bool) bool {
	if !appliedMigrations[version] {
		log.Printf("⏭️ skipping version %s (migration not applied)", version)
		return true
	}
	if appliedSeedings[version] {
		log.Printf("⏭️ skipping version %s (already seeded)", version)
		return true
	}
	return false
}

func main() {
	db := connect.ToDb()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("❌ error closing db")
		}
	}(db)

	ensureSeedHistoryTable(db)

	seedCount, err := strconv.Atoi(os.Getenv("SEED_COUNT"))
	if err != nil {
		log.Fatal("❌ invalid SEED_COUNT env: ", err)
	}

	appliedMigrations := checkAppliedFlywayVersions(db)
	appliedSeedings := checkAppliedSeedVersions(db)

	m := orderedmap.NewOrderedMap()
	m.Set("1", func(tx *sqlx.Tx, count int) error { return impl.NewV1DummySeeder().Seed() })
	m.Set("2", func(tx *sqlx.Tx, count int) error { return impl.NewV2Seeder(tx, count).Seed() })
	m.Set("3", func(tx *sqlx.Tx, count int) error { return impl.NewV3Seeder(tx, count).Seed() })
	m.Set("4", func(tx *sqlx.Tx, count int) error { return impl.NewV4Seeder(tx, count).Seed() })
	m.Set("5", func(tx *sqlx.Tx, count int) error { return impl.NewV5Seeder(tx, count).Seed() })
	m.Set("6", func(tx *sqlx.Tx, count int) error { return impl.NewV6Seeder(tx, count).Seed() })

	for el := m.Front(); el != nil; el = el.Next() {
		version, seedFunc, err := validateSeedersMap(el)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if shouldSkipVersion(version, appliedMigrations, appliedSeedings) {
			continue
		}
		log.Printf("⏳ seeding version %s...", version)

		tx, err := db.Beginx()
		if err != nil {
			log.Fatalf("❌ failed to begin transaction: %v", err)
		}

		err = seedFunc(tx, seedCount)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("❌ failed to rollback transaction for version %s: %v", version, rollbackErr)
			}
			log.Fatalf("❌ seeding failed for version %s: %v", version, err)
		}

		err = markSeedingApplied(tx, version)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("❌ failed to rollback transaction for version %s: %v", version, rollbackErr)
			}
			log.Fatalf("❌ failed to mark version %s as seeded: %v", version, err)
		}

		if err := tx.Commit(); err != nil {
			log.Fatalf("❌ transaction commit failed for version %s: %v", version, err)
		}

		log.Printf("✅ successfully seeded version %s", version)
	}
}
