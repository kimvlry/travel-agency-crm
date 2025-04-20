package seeder

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"reflect"
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

func (s *BaseSeeder) TryExecute(operationName, query string, args ...interface{}) error {
	_, err := s.Tx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s operation failed: %w", operationName, err)
	}
	return nil
}

func (s *BaseSeeder) TryGetAllEntities(
	entityName string,
	dest interface{},
	query string,
) error {
	err := s.Tx.Select(dest, query)
	if err != nil {
		return fmt.Errorf("failed to get %s: %w", entityName, err)
	}
	sliceValue := reflect.ValueOf(dest).Elem()
	if sliceValue.Len() == 0 {
		return fmt.Errorf("no %s found", entityName)
	}
	return nil
}
