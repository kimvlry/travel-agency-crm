package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

type Tour struct {
	ID                 uint    `gorm:"primaryKey"`
	Title              string  `gorm:"not null;unique"`
	PriceEUR           float64 `gorm:"type:decimal(10,2);not null"`
	Quota              int     `gorm:"not null"`
	MealsType          string
	IsLastMinute       bool
	LastMinuteApproved *bool
	BaseDurationDays   int `gorm:"not null"`
}

func (s *V3Seeder) seedTours() {
	for i := 0; i < s.count; i++ {
		approved := gofakeit.Bool()
		tour := Tour{
			Title:              gofakeit.Company() + " Tour",
			PriceEUR:           gofakeit.Price(100, 10000),
			Quota:              gofakeit.Number(10, 100),
			MealsType:          getRandomFromSlice(MealsTypes),
			IsLastMinute:       approved,
			LastMinuteApproved: &approved,
			BaseDurationDays:   gofakeit.Number(3, 14),
		}
		if err := s.db.Create(&tour).Error; err != nil {
			log.Println("error seeding tour:", err)
		}
	}
}

type Booking struct {
	ID             uint `gorm:"primaryKey"`
	TourID         uint
	Tour           Tour
	Status         string    `gorm:"type:booking_status;default:'draft'"`
	ContractNumber string    `gorm:"unique"`
	CreatedAt      time.Time `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"default:current_timestamp"`
}

func (s *V3Seeder) seedBookings() {
	var tours []Tour
	s.db.Find(&tours)

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 {
			log.Fatal("no tours found, cannot seed bookings")
		}

		booking := Booking{
			TourID:         getRandomFromSlice(tours).ID,
			Status:         getRandomFromSlice(BookingStatuses),
			ContractNumber: gofakeit.LetterN(10),
		}
		if err := s.db.Create(&booking).Error; err != nil {
			log.Println("error seeding booking:", err)
		}
	}
}

type BookingAgreement struct {
	ID         uint `gorm:"primaryKey"`
	ClientID   uint
	BookingID  uint
	SignedDate time.Time `gorm:"not null"`
}

func (s *V3Seeder) seedBookingAgreements() {
	var clients []Client
	var bookings []Booking
	s.db.Find(&clients)
	s.db.Find(&bookings)

	for i := 0; i < s.count; i++ {
		if len(clients) == 0 || len(bookings) == 0 {
			log.Fatal("missing clients or bookings")
		}

		ba := BookingAgreement{
			ClientID:   getRandomFromSlice(clients).ID,
			BookingID:  getRandomFromSlice(bookings).ID,
			SignedDate: gofakeit.Date(),
		}
		if err := s.db.Create(&ba).Error; err != nil {
			log.Println("error seeding booking agreement:", err)
		}
	}
}

type Assignee struct {
	ID       uint `gorm:"primaryKey"`
	ClientID uint
}

func (s *V3Seeder) seedAssignees() {
	var clients []Client
	s.db.Find(&clients)

	for i := 0; i < s.count; i++ {
		if len(clients) == 0 {
			log.Fatal("no clients found, cannot seed assignees")
		}
		assignee := Assignee{
			ClientID: getRandomFromSlice(clients).ID,
		}
		if err := s.db.Create(&assignee).Error; err != nil {
			log.Println("error seeding assignee:", err)
		}
	}
}

type ContractTemplate struct {
	ID                 uint   `gorm:"primaryKey"`
	Name               string `gorm:"not null"`
	Content            string `gorm:"not null"`
	ValidityPeriodDays int    `gorm:"not null"`
}

func (s *V3Seeder) seedContractTemplates() {
	for i := 0; i < s.count; i++ {
		ct := ContractTemplate{
			Name:               gofakeit.Word() + " contract",
			Content:            gofakeit.Paragraph(1, 3, 10, " "),
			ValidityPeriodDays: gofakeit.Number(30, 365),
		}
		if err := s.db.Create(&ct).Error; err != nil {
			log.Println("error seeding contract template:", err)
		}
	}
}

type Contract struct {
	ID         uint `gorm:"primaryKey"`
	AssigneeID uint
	TemplateID uint
	IssueDate  time.Time `gorm:"not null"`
	SignDate   *time.Time
	Status     string  `gorm:"not null"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null"`
}

func (s *V3Seeder) seedContracts() {
	var assignees []Assignee
	var templates []ContractTemplate
	s.db.Find(&assignees)
	s.db.Find(&templates)

	for i := 0; i < s.count; i++ {
		if len(assignees) == 0 || len(templates) == 0 {
			log.Fatal("missing assignees or templates")
		}

		issueDate := gofakeit.Date()
		signDate := issueDate.AddDate(0, 0, gofakeit.Number(1, 7))
		contract := Contract{
			AssigneeID: getRandomFromSlice(assignees).ID,
			TemplateID: getRandomFromSlice(templates).ID,
			IssueDate:  issueDate,
			SignDate:   &signDate,
			Status:     getRandomFromSlice(ConsentStatuses),
			TotalPrice: gofakeit.Price(500, 10000),
		}
		if err := s.db.Create(&contract).Error; err != nil {
			log.Println("error seeding contract:", err)
		}
	}
}

type ConsentTemplate struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Type    string `gorm:"not null"`
	Content string `gorm:"not null"`
}

func (s *V3Seeder) seedConsentTemplates() {
	for i := 0; i < s.count; i++ {
		ct := ConsentTemplate{
			Name:    gofakeit.AppName(),
			Type:    getRandomFromSlice(AgreementConsentTypes),
			Content: gofakeit.Paragraph(1, 2, 10, " "),
		}
		if err := s.db.Create(&ct).Error; err != nil {
			log.Println("error seeding consent template:", err)
		}
	}
}

type AgreementConsent struct {
	ID         uint `gorm:"primaryKey"`
	AssigneeID uint
	ContractID uint
	TemplateID *uint
	Date       time.Time
	Status     string
}

func (s *V3Seeder) seedAgreementConsents() {
	var assignees []Assignee
	var contracts []Contract
	var templates []ConsentTemplate
	s.db.Find(&assignees)
	s.db.Find(&contracts)
	s.db.Find(&templates)

	for i := 0; i < s.count; i++ {
		if len(assignees) == 0 || len(contracts) == 0 {
			log.Fatal("missing assignees or contracts")
		}

		var tmplID *uint
		if len(templates) > 0 && gofakeit.Bool() {
			id := getRandomFromSlice(templates).ID
			tmplID = &id
		}

		ac := AgreementConsent{
			AssigneeID: getRandomFromSlice(assignees).ID,
			ContractID: getRandomFromSlice(contracts).ID,
			TemplateID: tmplID,
			Date:       gofakeit.Date(),
			Status:     getRandomFromSlice(ConsentStatuses),
		}
		if err := s.db.Create(&ac).Error; err != nil {
			log.Println("error seeding agreement consent:", err)
		}
	}
}

type PaymentLink struct {
	ID        uint `gorm:"primaryKey"`
	BookingID uint
	URL       string `gorm:"not null"`
	QRCode    string `gorm:"not null"`
}

func (s *V3Seeder) seedPaymentLinks() {
	var bookings []Booking
	s.db.Find(&bookings)

	for i := 0; i < s.count; i++ {
		if len(bookings) == 0 {
			log.Fatal("no bookings found for payment links")
		}
		link := PaymentLink{
			BookingID: getRandomFromSlice(bookings).ID,
			URL:       gofakeit.URL(),
			QRCode:    gofakeit.UUID(),
		}
		if err := s.db.Create(&link).Error; err != nil {
			log.Println("error seeding payment link:", err)
		}
	}
}

type V3Seeder struct {
	db    *gorm.DB
	count int
}

func NewV3Seeder(db *gorm.DB, count int) *V3Seeder {
	return &V3Seeder{
		db:    db,
		count: count,
	}
}

func (s *V3Seeder) Seed() {
	s.seedTours()
	s.seedBookings()
	s.seedBookingAgreements()
	s.seedAssignees()
	s.seedContractTemplates()
	s.seedContracts()
	s.seedConsentTemplates()
	s.seedAgreementConsents()
	s.seedPaymentLinks()
}
