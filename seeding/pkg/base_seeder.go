package seeder

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type BaseSeeder struct {
	Tx        *sqlx.Tx
	SeedCount int
	Logger    *log.Logger
}

func NewBaseSeeder(tx *sqlx.Tx, seedCount int) *BaseSeeder {
	return &BaseSeeder{
		Tx:        tx,
		SeedCount: seedCount,
		Logger:    log.New(os.Stdout, "[seeder] ", log.LstdFlags),
	}
}
