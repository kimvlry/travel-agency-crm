package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"reflect"
)

type SeedOperation struct {
	Name    string
	Query   string
	Builder func(f *gofakeit.Faker) []interface{}
}

func (s *BaseSeeder) BatchSeed(ops []SeedOperation) error {
	faker := gofakeit.New(0)
	for _, op := range ops {
		for i := 0; i < s.SeedCount; i++ {
			args := op.Builder(faker)
			if err := s.TryExecute(op.Name, op.Query, args...); err != nil {
				s.Logger.Printf("[BATCH-SEED:ERROR]} in %q (iteration %d): %v", op.Name, i, err)
				return err
			} else {
				s.Logger.Printf("[BATCH-SEED:SUCCESS] %q (iteration %d)", op.Name, i)
			}
		}
	}
	return nil
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
