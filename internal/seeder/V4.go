package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

type TourIteration struct {
	ID        uint      `gorm:"primaryKey"`
	TourID    uint      `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}

func (s *V4Seeder) seedTourIterations() error {
	var tours []Tour
	s.db.Find(&tours)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 {
			log.Fatal("no tours found, cannot seed tour iterations")
		}

		iteration := TourIteration{
			TourID:    getRandomFromSlice(tours).ID,
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

type TourRoute struct {
	ID     uint `gorm:"primaryKey"`
	TourID uint `gorm:"not null"`
}

func (s *V4Seeder) seedTourRoutes() error {
	var tours []Tour
	s.db.Find(&tours)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 {
			log.Fatal("no tours found, cannot seed tour routes")
		}

		route := TourRoute{
			TourID: getRandomFromSlice(tours).ID,
		}
		if err := s.db.Create(&route).Error; err != nil {
			log.Println("error seeding tour route:", err)
			return err
		}
	}
	return nil
}

type RoutePoint struct {
	ID                   uint   `gorm:"primaryKey"`
	RouteID              uint   `gorm:"not null"`
	CityID               uint   `gorm:"not null"`
	Name                 string `gorm:"not null"`
	Address              string `gorm:"not null"`
	DurationTime         string `gorm:"not null"`
	InRouteOrderPosition int    `gorm:"not null"`
}

func (s *V4Seeder) seedRoutePoints() error {
	var routes []TourRoute
	var cities []City
	s.db.Find(&routes)
	s.db.Find(&cities)

	for i := 0; i < s.count; i++ {
		if len(routes) == 0 || len(cities) == 0 {
			log.Fatal("no routes or cities found, cannot seed route points")
		}

		point := RoutePoint{
			RouteID: getRandomFromSlice(routes).ID,
			CityID:  getRandomFromSlice(cities).ID,
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

type TransportService struct {
	ID      uint   `gorm:"primaryKey"`
	Company string `gorm:"not null"`
	Model   string `gorm:"not null"`
}

func (s *V4Seeder) seedTransportServices() error {
	for i := 0; i < s.count; i++ {
		service := TransportService{
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

type Transfer struct {
	ID             uint      `gorm:"primaryKey"`
	TourID         uint      `gorm:"not null"`
	TransportID    uint      `gorm:"not null"`
	DeparturePoint uint      `gorm:"not null"`
	ArrivalPoint   uint      `gorm:"not null"`
	DepartureTime  time.Time `gorm:"not null"`
}

func (s *V4Seeder) seedTransfers() error {
	var tours []Tour
	var transportServices []TransportService
	var routePoints []RoutePoint
	s.db.Find(&tours)
	s.db.Find(&transportServices)
	s.db.Find(&routePoints)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 || len(transportServices) == 0 || len(routePoints) == 0 {
			log.Fatal("missing data for transfers")
		}

		transfer := Transfer{
			TourID:         getRandomFromSlice(tours).ID,
			TransportID:    getRandomFromSlice(transportServices).ID,
			DeparturePoint: getRandomFromSlice(routePoints).ID,
			ArrivalPoint:   getRandomFromSlice(routePoints).ID,
			DepartureTime:  gofakeit.Date(),
		}
		if err := s.db.Create(&transfer).Error; err != nil {
			log.Println("error seeding transfer:", err)
			return err
		}
	}
	return nil
}

type Organizer struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Email string `gorm:"not null"`
}

func (s *V4Seeder) seedOrganizers() error {
	for i := 0; i < s.count; i++ {
		organizer := Organizer{
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

type Excursion struct {
	ID              uint      `gorm:"primaryKey"`
	TourID          uint      `gorm:"not null"`
	OrganizerID     uint      `gorm:"not null"`
	Name            string    `gorm:"not null"`
	MeetingLocation uint      `gorm:"not null"`
	MeetingTime     time.Time `gorm:"not null"`
}

func (s *V4Seeder) seedExcursions() error {
	var tours []Tour
	var organizers []Organizer
	var routePoints []RoutePoint
	s.db.Find(&tours)
	s.db.Find(&organizers)
	s.db.Find(&routePoints)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 || len(organizers) == 0 || len(routePoints) == 0 {
			log.Fatal("missing data for excursions")
		}

		excursion := Excursion{
			TourID:          getRandomFromSlice(tours).ID,
			OrganizerID:     getRandomFromSlice(organizers).ID,
			Name:            gofakeit.Word(),
			MeetingLocation: getRandomFromSlice(routePoints).ID,
			MeetingTime:     gofakeit.Date(),
		}
		if err := s.db.Create(&excursion).Error; err != nil {
			log.Println("error seeding excursion:", err)
			return err
		}
	}
	return nil
}

type InsuranceCompany struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Phone string `gorm:"not null"`
	Email string `gorm:"not null"`
}

func (s *V4Seeder) seedInsuranceCompanies() error {
	for i := 0; i < s.count; i++ {
		company := InsuranceCompany{
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

type Insurance struct {
	ID                 uint   `gorm:"primaryKey"`
	TourID             uint   `gorm:"not null"`
	InsuranceCompanyID uint   `gorm:"not null"`
	CoverageType       string `gorm:"not null"`
}

func (s *V4Seeder) seedInsurances() error {
	var tours []Tour
	var companies []InsuranceCompany
	s.db.Find(&tours)
	s.db.Find(&companies)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 || len(companies) == 0 {
			log.Fatal("missing data for insurances")
		}

		insurance := Insurance{
			TourID:             getRandomFromSlice(tours).ID,
			InsuranceCompanyID: getRandomFromSlice(companies).ID,
			CoverageType:       getRandomFromSlice(InsuranceTypes),
		}
		if err := s.db.Create(&insurance).Error; err != nil {
			log.Println("error seeding insurance:", err)
			return err
		}
	}
	return nil
}

type V4Seeder struct {
	db    *gorm.DB
	count int
}

func NewV4Seeder(db *gorm.DB, count int) *V4Seeder {
	return &V4Seeder{
		db:    db,
		count: count,
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
