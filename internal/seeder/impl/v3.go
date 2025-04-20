package impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"log"
	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V3Seeder struct {
	*seeder.BaseSeeder
}

func NewV3Seeder(tx *sqlx.Tx, seedCount int) *V3Seeder {
	return &V3Seeder{
		BaseSeeder: seeder.NewBaseSeeder(tx, seedCount),
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

func (s *V3Seeder) seedTours() error {
	for i := 0; i < s.SeedCount; i++ {

		approved := gofakeit.Bool()
		tour := models.Tour{
			Title:              gofakeit.Country() + gofakeit.UUID(),
			PriceEur:           gofakeit.Price(100, 10000),
			Quota:              gofakeit.Number(10, 100),
			MealsType:          seeder.getRandomFromSlice(models.MealsTypes),
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

func (s *V3Seeder) seedBookings() error {
	var tours []models.Tour
	if err := s.db.Find(&tours).Error; err != nil {
		return fmt.Errorf("failed to fetch tours: %w", err)
	}

	for i := 0; i < s.count; i++ {
		if len(tours) == 0 {
			return fmt.Errorf("no tours found, cannot seed bookings")
		}

		booking := models.Booking{
			TourID:         seeder.getRandomFromSlice(tours).ID,
			Status:         seeder.getRandomFromSlice(models.BookingStatuses),
			ContractNumber: gofakeit.LetterN(10),
		}
		if err := s.db.Create(&booking).Error; err != nil {
			log.Println("error seeding booking:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedBookingAgreements() error {
	var clients []models.Client
	var bookings []models.Booking
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}
	if err := s.db.Find(&bookings).Error; err != nil {
		return fmt.Errorf("failed to fetch bookings: %w", err)
	}

	for i := 0; i < s.SeedCount; i++ {
		if len(clients) == 0 || len(bookings) == 0 {
			return fmt.Errorf("missing clients or bookings")
		}

		ba := models.BookingAgreement{
			ClientID:   seeder.getRandomFromSlice(clients).ID,
			BookingID:  seeder.getRandomFromSlice(bookings).ID,
			SignedDate: gofakeit.Date(),
		}
		if err := s.db.Create(&ba).Error; err != nil {
			log.Println("error seeding booking agreement:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedAssignees() error {
	var clients []models.Client
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}

	for i := 0; i < s.SeedCount; i++ {
		if len(clients) == 0 {
			return fmt.Errorf("no clients found, cannot seed assignees")
		}
		assignee := models.Assignee{
			ClientID: seeder.getRandomFromSlice(clients).ID,
		}
		if err := s.db.Create(&assignee).Error; err != nil {
			log.Println("error seeding assignee:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedContractTemplates() error {
	for i := 0; i < s.SeedCount; i++ {
		ct := models.ContractTemplate{
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

func (s *V3Seeder) seedContracts() error {
	var assignees []models.Assignee
	var templates []models.ContractTemplate
	if err := s.db.Find(&assignees).Error; err != nil {
		return fmt.Errorf("failed to fetch assignees: %w", err)
	}
	if err := s.db.Find(&templates).Error; err != nil {
		return fmt.Errorf("failed to fetch contract templates: %w", err)
	}

	for i := 0; i < s.SeedCount; i++ {
		if len(assignees) == 0 || len(templates) == 0 {
			return fmt.Errorf("missing assignees or templates")
		}

		issueDate := gofakeit.Date()
		signDate := issueDate.AddDate(0, 0, gofakeit.Number(1, 7))
		contract := models.Contract{
			AssigneeID: seeder.getRandomFromSlice(assignees).ID,
			TemplateID: seeder.getRandomFromSlice(templates).ID,
			IssueDate:  issueDate,
			SignDate:   &signDate,
			Status:     seeder.getRandomFromSlice(models.ConsentStatuses),
			TotalPrice: gofakeit.Price(500, 10000),
		}
		if err := s.db.Create(&contract).Error; err != nil {
			log.Println("error seeding contract:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedConsentTemplates() error {
	for i := 0; i < s.SeedCount; i++ {
		ct := models.ConsentTemplate{
			Name:    gofakeit.AppName(),
			Type:    seeder.getRandomFromSlice(models.AgreementConsentTypes),
			Content: gofakeit.Paragraph(1, 2, 10, " "),
		}
		if err := s.db.Create(&ct).Error; err != nil {
			log.Println("error seeding consent template:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedAgreementConsents() error {
	var assignees []models.Assignee
	var contracts []models.Contract
	var templates []models.ConsentTemplate
	if err := s.db.Find(&assignees).Error; err != nil {
		return fmt.Errorf("failed to fetch assignees: %w", err)
	}
	if err := s.db.Find(&contracts).Error; err != nil {
		return fmt.Errorf("failed to fetch contracts: %w", err)
	}
	if err := s.db.Find(&templates).Error; err != nil {
		return fmt.Errorf("failed to fetch consent templates: %w", err)
	}

	for i := 0; i < s.SeedCount; i++ {
		if len(assignees) == 0 || len(contracts) == 0 {
			return fmt.Errorf("missing assignees or contracts")
		}

		var tmplID *uint
		if len(templates) > 0 && gofakeit.Bool() {
			id := seeder.getRandomFromSlice(templates).ID
			tmplID = &id
		}

		ac := models.AgreementConsent{
			AssigneeID: seeder.getRandomFromSlice(assignees).ID,
			ContractID: seeder.getRandomFromSlice(contracts).ID,
			TemplateID: tmplID,
			Date:       gofakeit.Date(),
			Status:     seeder.getRandomFromSlice(models.ConsentStatuses),
		}
		if err := s.db.Create(&ac).Error; err != nil {
			log.Println("error seeding agreement consent:", err)
			return err
		}
	}
	return nil
}

func (s *V3Seeder) seedPaymentLinks() error {
	var bookings []models.Booking
	if err := s.db.Find(&bookings).Error; err != nil {
		return fmt.Errorf("failed to fetch bookings: %w", err)
	}

	for i := 0; i < s.SeedCount; i++ {
		if len(bookings) == 0 {
			return fmt.Errorf("no bookings found for payment links")
		}
		link := models.PaymentLink{
			BookingID: seeder.getRandomFromSlice(bookings).ID,
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
