package impl

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strings"
	"travel-agency-seeder/internal/seeder"
)

type V6Seeder struct {
	*seeder.BaseSeeder
}

func NewV6Seeder(tx *sqlx.Tx, seedCount int) *V6Seeder {
	return &V6Seeder{
		BaseSeeder: seeder.NewBaseSeeder(tx, seedCount),
	}
}

func (s *V6Seeder) Seed() error {
	if err := s.SeedAnalysts(); err != nil {
		return err
	}
	return nil
}

func (s *V6Seeder) SeedAnalysts() error {
	analystNames := os.Getenv("ANALYST_NAMES")
	if analystNames == "" {
		log.Println("ANALYST_NAMES environment variable is not set")
		return fmt.Errorf("ANALYST_NAMES environment variable is not set")
	}
	names := strings.Split(analystNames, ",")

	for _, name := range names {
		password := fmt.Sprintf("%s_123", name)

		query := fmt.Sprintf(`
            DO $$
            BEGIN
                IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = '%s') THEN
                    CREATE ROLE "%s" WITH LOGIN PASSWORD '%s' INHERIT;
                END IF;
            END
            $$;
        `, name, name, password)

		_, err := s.Tx.Exec(query)
		if err != nil {
			log.Printf("failed to create user %s: %v\n", name, err)
			return fmt.Errorf("failed to create user %s: %w", name, err)
		}

		_, err = s.Tx.Exec(fmt.Sprintf(`
            GRANT analytic TO "%s";
        `, name))
		if err != nil {
			log.Printf("failed to grant role 'analytic' to user %s: %v\n", name, err)
			return fmt.Errorf("failed to grant role 'analytic' to user %s: %w", name, err)
		}
	}

	log.Println("successfully seeded analysts")
	return nil
}
