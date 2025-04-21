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
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				ON CONFLICT (title) do nothing
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			lm := faker.Bool()
			approved := lm && faker.Bool()
			return []interface{}{
				faker.Country() + faker.UUID(),
				faker.Price(100, 10000),
				faker.Number(10, 100),
				models.GetRandomFromSlice(models.MealsTypes),
				sql.NullBool{
					Bool:  lm,
					Valid: true},
				sql.NullBool{
					Bool:  approved,
					Valid: true},
				int32(faker.Number(3, 14)),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedBookings() error {
	var tours []models.Tour
	if err := s.TryGetAllEntities(
		"tours",
		&tours,
		"SELECT id FROM tours",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert bookings",
		Query: `INSERT INTO bookings 
    			(tour_id, status, contract_number) 
				VALUES ($1, $2, $3)
				ON CONFLICT (contract_number) do nothing
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			t := models.GetRandomFromSlice(tours)
			return []interface{}{
				t.ID,
				models.GetRandomFromSlice(models.BookingStatuses),
				faker.LetterN(10),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedBookingAgreements() error {
	var clients []models.Client
	var bookings []models.Booking
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"bookings",
		&bookings,
		"SELECT id FROM bookings",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert booking_agreements",
		Query: `INSERT INTO booking_agreements 
    			(client_id, booking_id, signed_date) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			b := models.GetRandomFromSlice(bookings)
			return []interface{}{
				c.ID,
				b.ID,
				faker.Date(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedAssignees() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert assignees",
		Query: `INSERT INTO assignees 
    			(client_id) 
				VALUES ($1)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			return []interface{}{
				c.ID,
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedContractTemplates() error {
	op := []seeder.SeedOperation{{
		Name: "insert contract_templates",
		Query: `INSERT INTO contract_templates 
    			(name, content, validity_period_days) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.Word() + " contract",
				faker.Paragraph(1, 3, 10, " "),
				faker.Number(30, 365),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedContracts() error {
	var assignees []models.Assignee
	var templates []models.ContractTemplate
	if err := s.TryGetAllEntities(
		"assignees",
		&assignees,
		"SELECT id FROM assignees",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"contract_templates",
		&templates,
		"SELECT id FROM contract_templates",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert contracts",
		Query: `INSERT INTO contracts 
    			(assignee_id, template_id, issue_date, sign_date, status, total_price) 
				VALUES ($1, $2, $3, $4, $5, $6)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			a := models.GetRandomFromSlice(assignees)
			t := models.GetRandomFromSlice(templates)
			issue := faker.Date()
			signed := issue.AddDate(0, 0, faker.Number(1, 7))
			return []interface{}{
				a.ID,
				t.ID,
				issue,
				sql.NullTime{
					Time:  signed,
					Valid: true,
				},
				models.GetRandomFromSlice(models.ConsentStatuses),
				faker.Price(500, 10000),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedConsentTemplates() error {
	op := []seeder.SeedOperation{{
		Name: "insert consent_templates",
		Query: `INSERT INTO consent_templates 
    			(name, type, content) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.AppName(),
				models.GetRandomFromSlice(models.AgreementConsentTypes),
				faker.Paragraph(1, 2, 10, " "),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedAgreementConsents() error {
	var assignees []models.Assignee
	var contracts []models.Contract
	var templates []models.ConsentTemplate
	if err := s.TryGetAllEntities(
		"assignees",
		&assignees,
		"SELECT id FROM assignees",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"contracts",
		&contracts,
		"SELECT id FROM contracts",
	); err != nil {
		return err
	}
	if err := s.TryGetAllEntities(
		"consent_templates",
		&templates,
		"SELECT id FROM consent_templates",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert agreement_consents",
		Query: `INSERT INTO agreement_consents 
    			(assignee_id, contract_id, template_id, date, status) 
				VALUES ($1, $2, $3, $4, $5)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			a := models.GetRandomFromSlice(assignees)
			c := models.GetRandomFromSlice(contracts)
			t := models.GetRandomFromSlice(templates)
			return []interface{}{
				a.ID,
				c.ID,
				t.ID,
				faker.Date(),
				models.GetRandomFromSlice(models.ConsentStatuses),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V3Seeder) seedPaymentLinks() error {
	var bookings []models.Booking
	if err := s.TryGetAllEntities(
		"bookings",
		&bookings,
		"SELECT id FROM bookings",
	); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert payment_links",
		Query: `INSERT INTO payment_links 
    			(booking_id, url, qr_code) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			b := models.GetRandomFromSlice(bookings)
			return []interface{}{
				b.ID,
				faker.URL(),
				faker.UUID(),
			}
		},
	}}
	return s.BatchSeed(op)
}
