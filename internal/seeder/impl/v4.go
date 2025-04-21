package impl

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"

	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V4Seeder struct {
	*seeder.BaseSeeder
}

func NewV4Seeder(tx *sqlx.Tx, seedCount int) *V4Seeder {
	return &V4Seeder{BaseSeeder: seeder.NewBaseSeeder(tx, seedCount)}
}

func (s *V4Seeder) Seed() error {
	if err := s.seedTourIterations(); err != nil {
		return err
	}
	if err := s.seedTourRoutes(); err != nil {
		return err
	}
	if err := s.seedRoutePoints(); err != nil {
		return err
	}
	if err := s.seedTransportServices(); err != nil {
		return err
	}
	if err := s.seedTransfers(); err != nil {
		return err
	}
	if err := s.seedOrganizers(); err != nil {
		return err
	}
	if err := s.seedExcursions(); err != nil {
		return err
	}
	if err := s.seedInsuranceCompanies(); err != nil {
		return err
	}
	if err := s.seedInsurances(); err != nil {
		return err
	}
	return nil
}

func (s *V4Seeder) seedTourIterations() error {
	var tours []models.Tour
	if err := s.TryGetAllEntities(
		"tours",
		&tours, "SELECT id FROM tours",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert tour_iterations",
		Query: `INSERT INTO tour_iterations 
    			(tour_id, start_date, end_date) 
				VALUES ($1, $2, $3)
				ON CONFLICT (tour_id, start_date) do nothing 
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			t := models.GetRandomFromSlice(tours)
			return []interface{}{
				t.ID,
				faker.Date(),
				faker.Date(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedTourRoutes() error {
	var tours []models.Tour
	if err := s.TryGetAllEntities(
		"tours",
		&tours,
		"SELECT id FROM tours",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert tour_routes",
		Query: `INSERT INTO tour_routes 
                (tour_id) 
				VALUES ($1)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			t := models.GetRandomFromSlice(tours)
			return []interface{}{
				t.ID,
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedRoutePoints() error {
	var routes []models.TourRoute
	var cities []models.City

	if err := s.TryGetAllEntities(
		"tour_routes",
		&routes,
		"SELECT id FROM tour_routes",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"cities",
		&cities,
		"SELECT id FROM cities",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert route_points",
		Query: `INSERT INTO route_points 
    			(route_id, city_id, name, address, duration_time, in_route_order_position) 
				VALUES ($1, $2, $3, $4, $5, $6)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			r := models.GetRandomFromSlice(routes)
			c := models.GetRandomFromSlice(cities)
			return []interface{}{
				r.ID,
				c.ID,
				faker.City(),
				faker.Address().Address,
				fmt.Sprintf(
					"%dh %dm",
					faker.Number(1, 5),
					faker.Number(0, 59),
				),
				faker.Number(1, 10),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedTransportServices() error {
	op := []seeder.SeedOperation{{
		Name: "insert transport_services",
		Query: `INSERT INTO transport_services (company, model) 
				VALUES ($1, $2)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.Company(),
				faker.CarModel(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedTransfers() error {
	var tours []models.Tour
	var transports []models.TransportService
	var points []models.RoutePoint

	if err := s.TryGetAllEntities(
		"tours",
		&tours,
		"SELECT id FROM tours",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"transport_services",
		&transports,
		"SELECT id FROM transport_services",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"route_points",
		&points,
		"SELECT id FROM route_points",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert transfers",
		Query: `INSERT INTO transfers 
                (tour_id, transport_id, departure_point, arrival_point, departure_time) 
				VALUES ($1, $2, $3, $4, $5)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				models.GetRandomFromSlice(tours).ID,
				models.GetRandomFromSlice(transports).ID,
				models.GetRandomFromSlice(points).ID,
				models.GetRandomFromSlice(points).ID,
				faker.Date(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedOrganizers() error {
	op := []seeder.SeedOperation{{
		Name: "insert organizers",
		Query: `INSERT INTO organizers 
                (name, phone, email) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.Name(),
				faker.Phone(),
				faker.Email()}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedExcursions() error {
	var tours []models.Tour
	var organizers []models.Organizer
	var points []models.RoutePoint

	if err := s.TryGetAllEntities(
		"tours",
		&tours,
		"SELECT id FROM tours",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"organizers",
		&organizers,
		"SELECT id FROM organizers",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"route_points",
		&points, "SELECT id FROM route_points",
	); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert excursions",
		Query: `INSERT INTO excursions 
    			(tour_id, organizer_id, name, meeting_location, meeting_time) 
				VALUES ($1, $2, $3, $4, $5)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				models.GetRandomFromSlice(tours).ID,
				models.GetRandomFromSlice(organizers).ID,
				faker.Word(),
				models.GetRandomFromSlice(points).ID,
				faker.Date(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedInsuranceCompanies() error {
	op := []seeder.SeedOperation{{
		Name: "insert insurance_companies",
		Query: `INSERT INTO insurance_companies 
    			(name, phone, email) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				f.Company(),
				f.Phone(),
				f.Email(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V4Seeder) seedInsurances() error {
	var companies []models.InsuranceCompany
	var tours []models.Tour

	if err := s.TryGetAllEntities(
		"tours",
		&tours,
		"SELECT id FROM tours",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"insurance_companies",
		&companies,
		"SELECT id FROM insurance_companies",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert insurances",
		Query: `INSERT INTO insurances 
    			(tour_id, insurance_company_id, coverage_type) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				models.GetRandomFromSlice(tours).ID,
				models.GetRandomFromSlice(companies).ID,
				models.GetRandomFromSlice(models.InsuranceTypes),
			}
		},
	}}
	return s.BatchSeed(op)
}
