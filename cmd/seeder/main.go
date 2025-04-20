package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"travel-agency-seeder/internal/seeder/impl"
)

func connectToDb() *sqlx.DB {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	if dbName == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("environment variables not set")
	}

	connectionString := fmt.
		Sprintf("user=%s password=%s dbname=%s host=postgres port=5432 sslmode=disable",
			dbUser, dbPassword, dbName)
	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		log.Fatal("couldn't connect to db", err)
	}
	return db
}

func ensureSeedHistoryTable(db *sqlx.DB) {
	query := `
		create table if not exists seed_history (
		    id serial primary key,
		    version varchar(10) not null unique,
		    seeded_at timestampz not null default now()
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("failed to create seed_history table: %v", err)
	}
}

func checkAppliedSeedVersions(db *sqlx.DB) map[string]bool {
	var versions []string
	result := db.Select(&versions, "select version from seed_history")
	if result.Error != nil {
		log.Fatalf("error reading seed_history: %v", result.Error)
	}

	applied := make(map[string]bool)
	for _, v := range versions {
		applied[v] = true
	}
	return applied
}

func checkAppliedFlywayVersions(db *sqlx.DB) map[string]bool {
	var versions []string
	query := `select version from flyway_schema_history where success = true`
	err := db.Select(&versions, query)
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
	_, err := tx.Exec(`
        insert into seed_history (version) values ($1)`,
		version,
	)
	return err
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

	var seeders = map[string]func(*sqlx.Tx, int) error{
		"1": func(tx *sqlx.Tx, count int) error {
			return impl.NewV1DummySeeder().Seed()
		},
		"2": func(tx *sqlx.Tx, count int) error {
			return impl.NewV2Seeder(tx, count).Seed()
		},
		"3": func(tx *sqlx.Tx, count int) error {
			return impl.NewV3Seeder(tx, count).Seed()
		},
		"4": func(tx *sqlx.Tx, count int) error {
			return impl.NewV4Seeder(tx, count).Seed()
		},
		"5": func(tx *sqlx.Tx, count int) error {
			return impl.NewV5Seeder(tx, count).Seed()
		},
		"6": func(tx *sqlx.Tx, count int) error {
			return impl.NewV6Seeder(tx, count).Seed()
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

		func() {
			tx, err := db.Beginx()
			if err != nil {
				log.Printf("failed to start transaction: %v", err)
				return
			}
			committed := false

			defer func() {
				if !committed {
					if err := tx.Rollback(); err != nil {
						log.Printf("rollback failed: %v", err)
					}
				}
			}()

			if err := seedFunc(tx, seedCount); err != nil {
				log.Printf("seeding failed for version %s: %v", version, err)
				return
			}
			if err := markSeedingApplied(tx, version); err != nil {
				log.Printf("failed to mark version %s as seeded: %v", version, err)
				return
			}
			if err := tx.Commit(); err != nil {
				log.Printf("commit failed for version %s: %v", version, err)
				return
			}

			committed = true
			log.Printf("successfully seeded version %s", version)
		}()
	}
}
