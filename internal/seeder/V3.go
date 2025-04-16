package seeder

import (
	"fmt"
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

func (s *V3Seeder) seedTours() error {
	for i := 0; i < s.count; i++ {

		approved := gofakeit.Bool()
		tour := Tour{
			Title:              gofakeit.Country() + gofakeit.UUID(),
			PriceEUR:           gofakeit.Price(100, 10000),
			Quota:              gofakeit.Number(10, 100),
			MealsType:          getRandomFromSlice(MealsTypes),
			IsLastMinute:       approved,
			LastMinuteApproved: &approved,
			BaseDurationDays:   gofakeit.Number(3, 14),
		}
		if err := s.db.Create(&tour).Error; err != nil {
			log.Println("error seeding tour:", err)
			return err
		}
	}
	return nil
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

func (s *V3Seeder) seedBookings() error {
	var tours []Tour
	if err := s.db.Find(&tours).Error; err != nil {
		return fmt.Errorf("failed to fetch tours: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 {
			return fmt.Errorf("no tours found, cannot seed bookings")
		}

		booking := Booking{
			TourID:         getRandomFromSlice(tours).ID,
			Status:         getRandomFromSlice(BookingStatuses),
			ContractNumber: gofakeit.LetterN(10),
		}
		if err := s.db.Create(&booking).Error; err != nil {
			log.Println("error seeding booking:", err)
			return err
		}
	}
	return nil
}

type BookingAgreement struct {
	ID         uint `gorm:"primaryKey"`
	ClientID   uint
	BookingID  uint
	SignedDate time.Time `gorm:"not null"`
}

func (s *V3Seeder) seedBookingAgreements() error {
	var clients []Client
	var bookings []Booking
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}
	if err := s.db.Find(&bookings).Error; err != nil {
		return fmt.Errorf("failed to fetch bookings: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(clients) == 0 || len(bookings) == 0 {
			return fmt.Errorf("missing clients or bookings")
		}

		ba := BookingAgreement{
			ClientID:   getRandomFromSlice(clients).ID,
			BookingID:  getRandomFromSlice(bookings).ID,
			SignedDate: gofakeit.Date(),
		}
		if err := s.db.Create(&ba).Error; err != nil {
			log.Println("error seeding booking agreement:", err)
			return err
		}
	}
	return nil
}

type Assignee struct {
	ID       uint `gorm:"primaryKey"`
	ClientID uint
}

func (s *V3Seeder) seedAssignees() error {
	var clients []Client
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(clients) == 0 {
			return fmt.Errorf("no clients found, cannot seed assignees")
		}
		assignee := Assignee{
			ClientID: getRandomFromSlice(clients).ID,
		}
		if err := s.db.Create(&assignee).Error; err != nil {
			log.Println("error seeding assignee:", err)
			return err
		}
	}
	return nil
}

type ContractTemplate struct {
	ID                 uint   `gorm:"primaryKey"`
	Name               string `gorm:"not null"`
	Content            string `gorm:"not null"`
	ValidityPeriodDays int    `gorm:"not null"`
}

func (s *V3Seeder) seedContractTemplates() error {
	for i := 0; i < s.count; i++ {
		ct := ContractTemplate{
			Name:               gofakeit.Word() + " contract",
			Content:            gofakeit.Paragraph(1, 3, 10, " "),
			ValidityPeriodDays: gofakeit.Number(30, 365),
		}
		if err := s.db.Create(&ct).Error; err != nil {
			log.Println("error seeding contract template:", err)
			return err
		}
	}
	return nil
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

func (s *V3Seeder) seedContracts() error {
	var assignees []Assignee
	var templates []ContractTemplate
	if err := s.db.Find(&assignees).Error; err != nil {
		return fmt.Errorf("failed to fetch assignees: %w", err)
	}
	if err := s.db.Find(&templates).Error; err != nil {
		return fmt.Errorf("failed to fetch contract templates: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(assignees) == 0 || len(templates) == 0 {
			return fmt.Errorf("missing assignees or templates")
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
			return err
		}
	}
	return nil
}

type ConsentTemplate struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Type    string `gorm:"not null"`
	Content string `gorm:"not null"`
}

func (s *V3Seeder) seedConsentTemplates() error {
	for i := 0; i < s.count; i++ {
		ct := ConsentTemplate{
			Name:    gofakeit.AppName(),
			Type:    getRandomFromSlice(AgreementConsentTypes),
			Content: gofakeit.Paragraph(1, 2, 10, " "),
		}
		if err := s.db.Create(&ct).Error; err != nil {
			log.Println("error seeding consent template:", err)
			return err
		}
	}
	return nil
}

type AgreementConsent struct {
	ID         uint `gorm:"primaryKey"`
	AssigneeID uint
	ContractID uint
	TemplateID *uint
	Date       time.Time
	Status     string
}

func (s *V3Seeder) seedAgreementConsents() error {
	var assignees []Assignee
	var contracts []Contract
	var templates []ConsentTemplate
	if err := s.db.Find(&assignees).Error; err != nil {
		return fmt.Errorf("failed to fetch assignees: %w", err)
	}
	if err := s.db.Find(&contracts).Error; err != nil {
		return fmt.Errorf("failed to fetch contracts: %w", err)
	}
	if err := s.db.Find(&templates).Error; err != nil {
		return fmt.Errorf("failed to fetch consent templates: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(assignees) == 0 || len(contracts) == 0 {
			return fmt.Errorf("missing assignees or contracts")
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
			return err
		}
	}
	return nil
}

type PaymentLink struct {
	ID        uint `gorm:"primaryKey"`
	BookingID uint
	URL       string `gorm:"not null"`
	QRCode    string `gorm:"not null"`
}

func (s *V3Seeder) seedPaymentLinks() error {
	var bookings []Booking
	if err := s.db.Find(&bookings).Error; err != nil {
		return fmt.Errorf("failed to fetch bookings: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(bookings) == 0 {
			return fmt.Errorf("no bookings found for payment links")
		}
		link := PaymentLink{
			BookingID: getRandomFromSlice(bookings).ID,
			URL:       gofakeit.URL(),
			QRCode:    gofakeit.UUID(),
		}
		if err := s.db.Create(&link).Error; err != nil {
			log.Println("error seeding payment link:", err)
			return err
		}
	}
	return nil
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

func (s *V3Seeder) Seed() error {
	if err := s.seedTours(); err != nil {
		return err
	}
	if err := s.seedBookings(); err != nil {
		return err
	}
	if err := s.seedBookingAgreements(); err != nil {
		return err
	}
	if err := s.seedAssignees(); err != nil {
		return err
	}
	if err := s.seedContractTemplates(); err != nil {
		return err
	}
	if err := s.seedContracts(); err != nil {
		return err
	}
	if err := s.seedConsentTemplates(); err != nil {
		return err
	}
	if err := s.seedAgreementConsents(); err != nil {
		return err
	}
	if err := s.seedPaymentLinks(); err != nil {
		return err
	}
	return nil
}
