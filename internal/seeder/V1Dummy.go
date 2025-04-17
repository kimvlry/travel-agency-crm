package seeder

type V1DummySeeder struct{}

func NewV1DummySeeder() *V1DummySeeder {
	return &V1DummySeeder{}
}

func (s *V1DummySeeder) Seed() error {
	return nil
}
