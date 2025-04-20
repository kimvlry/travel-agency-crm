package impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"log"
	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V4Seeder struct {
	*seeder.BaseSeeder
}

func NewV4Seeder(tx *sqlx.Tx, seedCount int) *V4Seeder {
	return &V4Seeder{
		BaseSeeder: seeder.NewBaseSeeder(tx, seedCount),
	}
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
	s.db.Find(&tours)

	for i := 0; i < s.SeedCount; i++ {
		if len(tours) == 0 {
			log.Fatal("no tours found, cannot seed tour iterations")
		}

		iteration := models.TourIteration{
			TourID:    seeder.getRandomFromSlice(tours).ID,
			StartDate: gofakeit.Date(),
			EndDate:   gofakeit.Date(),
		}
		if err := s.db.Create(&iteration).Error; err != nil {
			log.Println("error seeding tour iteration:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedTourRoutes() error {
	var tours []models.Tour
	s.db.Find(&tours)

	for i := 0; i < s.SeedCount; i++ {
		if len(tours) == 0 {
			log.Fatal("no tours found, cannot seed tour routes")
		}

		route := models.TourRoute{
			TourID: seeder.getRandomFromSlice(tours).ID,
		}
		if err := s.db.Create(&route).Error; err != nil {
			log.Println("error seeding tour route:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedRoutePoints() error {
	var routes []models.TourRoute
	var cities []models.City
	s.db.Find(&routes)
	s.db.Find(&cities)

	for i := 0; i < s.SeedCount; i++ {
		if len(routes) == 0 || len(cities) == 0 {
			log.Fatal("no routes or cities found, cannot seed route points")
		}

		point := models.RoutePoint{
			RouteID: seeder.getRandomFromSlice(routes).ID,
			CityID:  seeder.getRandomFromSlice(cities).ID,
			Name:    gofakeit.City(),
			Address: gofakeit.Address().Address,
			DurationTime: fmt.Sprintf("%dh %dm",
				gofakeit.Number(1, 5),
				gofakeit.Number(0, 59)),
			InRouteOrderPosition: gofakeit.Number(1, 10),
		}
		if err := s.db.Create(&point).Error; err != nil {
			log.Println("error seeding route point:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedTransportServices() error {
	for i := 0; i < s.SeedCount; i++ {
		service := models.TransportService{
			Company: gofakeit.Company(),
			Model:   gofakeit.CarModel(),
		}
		if err := s.db.Create(&service).Error; err != nil {
			log.Println("error seeding transport service:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedTransfers() error {
	var tours []models.Tour
	var transportServices []models.TransportService
	var routePoints []models.RoutePoint
	s.db.Find(&tours)
	s.db.Find(&transportServices)
	s.db.Find(&routePoints)

	for i := 0; i < s.SeedCount; i++ {
		if len(tours) == 0 || len(transportServices) == 0 || len(routePoints) == 0 {
			log.Fatal("missing data for transfers")
		}

		transfer := models.Transfer{
			TourID:         seeder.getRandomFromSlice(tours).ID,
			TransportID:    seeder.getRandomFromSlice(transportServices).ID,
			DeparturePoint: seeder.getRandomFromSlice(routePoints).ID,
			ArrivalPoint:   seeder.getRandomFromSlice(routePoints).ID,
			DepartureTime:  gofakeit.Date(),
		}
		if err := s.db.Create(&transfer).Error; err != nil {
			log.Println("error seeding transfer:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedOrganizers() error {
	for i := 0; i < s.SeedCount; i++ {
		organizer := models.Organizer{
			Name:  gofakeit.Name(),
			Phone: gofakeit.Phone(),
			Email: gofakeit.Email(),
		}
		if err := s.db.Create(&organizer).Error; err != nil {
			log.Println("error seeding organizer:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedExcursions() error {
	var tours []models.Tour
	var organizers []models.Organizer
	var routePoints []models.RoutePoint
	s.db.Find(&tours)
	s.db.Find(&organizers)
	s.db.Find(&routePoints)

	for i := 0; i < s.SeedCount; i++ {
		if len(tours) == 0 || len(organizers) == 0 || len(routePoints) == 0 {
			log.Fatal("missing data for excursions")
		}

		excursion := models.Excursion{
			TourID:          seeder.getRandomFromSlice(tours).ID,
			OrganizerID:     seeder.getRandomFromSlice(organizers).ID,
			Name:            gofakeit.Word(),
			MeetingLocation: seeder.getRandomFromSlice(routePoints).ID,
			MeetingTime:     gofakeit.Date(),
		}
		if err := s.db.Create(&excursion).Error; err != nil {
			log.Println("error seeding excursion:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedInsuranceCompanies() error {
	for i := 0; i < s.SeedCount; i++ {
		company := models.InsuranceCompany{
			Name:  gofakeit.Company(),
			Phone: gofakeit.Phone(),
			Email: gofakeit.Email(),
		}
		if err := s.db.Create(&company).Error; err != nil {
			log.Println("error seeding insurance company:", err)
			return err
		}
	}
	return nil
}

func (s *V4Seeder) seedInsurances() error {
	var tours []models.Tour
	var companies []models.InsuranceCompany
	s.db.Find(&tours)
	s.db.Find(&companies)

	for i := 0; i < s.SeedCount; i++ {
		if len(tours) == 0 || len(companies) == 0 {
			log.Fatal("missing data for insurances")
		}

		insurance := models.Insurance{
			TourID:             seeder.getRandomFromSlice(tours).ID,
			InsuranceCompanyID: seeder.getRandomFromSlice(companies).ID,
			CoverageType:       seeder.getRandomFromSlice(models.InsuranceTypes),
		}
		if err := s.db.Create(&insurance).Error; err != nil {
			log.Println("error seeding insurance:", err)
			return err
		}
	}
	return nil
}
