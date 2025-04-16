package seeder

type V1Seeder struct{}

func NewV1Seeder() *V1Seeder {
	return &V1Seeder{}
}

func (s *V1Seeder) Seed() {}
