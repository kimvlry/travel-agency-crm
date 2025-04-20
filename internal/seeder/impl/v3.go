package impl

import (
	"database/sql"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"

	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V3Seeder struct {
	*seeder.BaseSeeder
}

func NewV3Seeder(tx *sqlx.Tx, seedCount int) *V3Seeder {
	return &V3Seeder{BaseSeeder: seeder.NewBaseSeeder(tx, seedCount)}
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
	op := []seeder.SeedOperation{{
		Name: "insert tours",
		Query: `INSERT INTO tours
			(title, price_eur, quota, meals_type, is_last_minute, last_minute_approved, base_duration_days)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			lm := f.Bool()
			approved := lm && f.Bool()
			return []interface{}{
				f.Country() + f.UUID(),
				f.Price(100, 10000),
				f.Number(10, 100),
				models.MealsType(models.GetRandomFromSlice(models.MealsTypes)),
				sql.NullBool{Bool: lm, Valid: true},
				sql.NullBool{Bool: approved, Valid: true},
				int32(f.Number(3, 14)),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedBookings() error {
	var tours []models.Tour
	if err := s.TryGetAllEntities("tours", &tours, "SELECT id FROM tours"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name:  "insert bookings",
		Query: `INSERT INTO bookings (tour_id, status, contract_number) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			t := models.GetRandomFromSlice(tours)
			return []interface{}{t.ID,
				models.BookingStatus(models.GetRandomFromSlice(models.BookingStatuses)),
				f.LetterN(10),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedBookingAgreements() error {
	var clients []models.Client
	var bookings []models.Booking
	if err := s.TryGetAllEntities("clients", &clients, "SELECT id FROM clients"); err != nil {
		return err
	}
	if err := s.TryGetAllEntities("bookings", &bookings, "SELECT id FROM bookings"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name:  "insert booking_agreements",
		Query: `INSERT INTO booking_agreements (client_id, booking_id, signed_date) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			b := models.GetRandomFromSlice(bookings)
			return []interface{}{c.ID, b.ID, f.Date()}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedAssignees() error {
	var clients []models.Client
	if err := s.TryGetAllEntities("clients", &clients, "SELECT id FROM clients"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name:  "insert assignees",
		Query: `INSERT INTO assignees (client_id) VALUES ($1)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			return []interface{}{c.ID}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedContractTemplates() error {
	op := []seeder.SeedOperation{{
		Name:  "insert contract_templates",
		Query: `INSERT INTO contract_templates (name, content, validity_period_days) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				f.Word() + " contract",
				f.Paragraph(1, 3, 10, " "),
				f.Number(30, 365),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedContracts() error {
	var assignees []models.Assignee
	var templates []models.ContractTemplate
	if err := s.TryGetAllEntities("assignees", &assignees, "SELECT id FROM assignees"); err != nil {
		return err
	}
	if err := s.TryGetAllEntities("contract_templates", &templates, "SELECT id FROM contract_templates"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert contracts",
		Query: `INSERT INTO contracts (assignee_id, template_id, issue_date, sign_date, status, total_price) 
				VALUES ($1, $2, $3, $4, $5, $6)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			a := models.GetRandomFromSlice(assignees)
			t := models.GetRandomFromSlice(templates)
			issue := f.Date()
			signed := issue.AddDate(0, 0, f.Number(1, 7))
			return []interface{}{
				a.ID,
				t.ID,
				issue,
				sql.NullTime{Time: signed, Valid: true},
				models.ConsentStatus(models.GetRandomFromSlice(models.ConsentStatuses)),
				f.Price(500, 10000),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedConsentTemplates() error {
	op := []seeder.SeedOperation{{
		Name:  "insert consent_templates",
		Query: `INSERT INTO consent_templates (name, type, content) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				f.AppName(),
				models.AgreementConsentType(models.GetRandomFromSlice(models.AgreementConsentTypes)),
				f.Paragraph(1, 2, 10, " "),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedAgreementConsents() error {
	var assignees []models.Assignee
	var contracts []models.Contract
	var templates []models.ConsentTemplate
	if err := s.TryGetAllEntities("assignees", &assignees, "SELECT id FROM assignees"); err != nil {
		return err
	}
	if err := s.TryGetAllEntities("contracts", &contracts, "SELECT id FROM contracts"); err != nil {
		return err
	}
	if err := s.TryGetAllEntities("consent_templates", &templates, "SELECT id FROM consent_templates"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert agreement_consents",
		Query: `INSERT INTO agreement_consents (assignee_id, contract_id, template_id, date, status) 
				VALUES ($1, $2, $3, $4, $5)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			a := models.GetRandomFromSlice(assignees)
			c := models.GetRandomFromSlice(contracts)
			var tid interface{} = nil
			if len(templates) > 0 && f.Bool() {
				t := models.GetRandomFromSlice(templates)
				tid = t.ID
			}
			return []interface{}{
				a.ID,
				c.ID,
				tid,
				f.Date(),
				models.ConsentStatus(models.GetRandomFromSlice(models.ConsentStatuses)),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedPaymentLinks() error {
	var bookings []models.Booking
	if err := s.TryGetAllEntities("bookings", &bookings, "SELECT id FROM bookings"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name:  "insert payment_links",
		Query: `INSERT INTO payment_links (booking_id, url, qr_code) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			b := models.GetRandomFromSlice(bookings)
			return []interface{}{b.ID, f.URL(), f.UUID()}
		},
	}}
	return s.BatchSeed(op)
}
