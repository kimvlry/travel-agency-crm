package impl

import (
	"database/sql"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"

	"travel-agency-seeder/models"
	"travel-agency-seeder/seeding/pkg"
)

type V2Seeder struct {
	*seeder.BaseSeeder
}

func NewV2Seeder(tx *sqlx.Tx, seedCount int) *V2Seeder {
	return &V2Seeder{
		BaseSeeder: seeder.NewBaseSeeder(tx, seedCount),
	}
}

func (s *V2Seeder) Seed() error {
	if err := s.seedCountries(); err != nil {
		return err
	}
	if err := s.seedCities(); err != nil {
		return err
	}
	if err := s.seedClients(); err != nil {
		return err
	}
	if err := s.seedBans(); err != nil {
		return err
	}
	if err := s.seedPassports(); err != nil {
		return err
	}
	if err := s.seedPromotions(); err != nil {
		return err
	}
	if err := s.seedNotificationTemplates(); err != nil {
		return err
	}
	if err := s.seedReminders(); err != nil {
		return err
	}
	if err := s.seedInteractions(); err != nil {
		return err
	}
	return s.seedPersonalNotifications()
}

func (s *V2Seeder) seedCountries() error {
	op := []seeder.SeedOperation{{
		Name:  "insert countries",
		Query: `INSERT INTO countries (name) VALUES ($1)`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.Country(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedCities() error {
	var countries []models.Country
	if err := s.TryGetAllEntities(
		"countries",
		&countries,
		"SELECT id FROM countries",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert cities",
		Query: `INSERT INTO cities 
    			(name, country_id, timezone_offset) 
				VALUES ($1, $2, $3)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(countries)
			return []interface{}{
				faker.City(),
				c.ID,
				faker.Number(-12, 12)}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedClients() error {
	var cities []models.City
	if err := s.TryGetAllEntities(
		"cities",
		&cities,
		"SELECT id FROM cities",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert clients",
		Query: `INSERT INTO clients 
    			(full_name, phone, email, birth_date, city_id) 
				VALUES ($1, $2, $3, $4, $5)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(cities)
			return []interface{}{
				faker.Name(),
				faker.Phone(),
				faker.Email(),
				faker.DateRange(
					time.Now().AddDate(-60, 0, 0),
					time.Now().AddDate(-18, 0, 0),
				),
				c.ID,
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedBans() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{
		{
			Name: "insert bans",
			Query: `INSERT INTO bans 
    				(client_id, ban_reason) 
					VALUES ($1, $2)
					`,
			Builder: func(faker *gofakeit.Faker) []interface{} {
				c := models.GetRandomFromSlice(clients)
				return []interface{}{
					c.ID,
					faker.Sentence(5),
				}
			},
		},
		{
			Name:  "update blacklist",
			Query: `UPDATE clients SET is_blacklisted = true WHERE id = $1`,
			Builder: func(f *gofakeit.Faker) []interface{} {
				c := models.GetRandomFromSlice(clients)
				return []interface{}{
					c.ID,
				}
			},
		},
	}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedPassports() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert passports",
		Query: `INSERT INTO passports 
    			(client_id, type, number, expiration_date, issue_date) 
				VALUES ($1, $2, $3, $4, $5)
				ON CONFLICT (client_id, type) DO NOTHING
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			return []interface{}{
				c.ID,
				models.GetRandomFromSlice(models.PassportTypes),
				faker.UUID(),
				faker.DateRange(
					time.Now().AddDate(5, 0, 0),
					time.Now().AddDate(10, 0, 0),
				),
				faker.DateRange(
					time.Now().AddDate(-10, 0, 0),
					time.Now().AddDate(-1, 0, 0),
				),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedPromotions() error {
	op := []seeder.SeedOperation{{
		Name: "insert promotions",
		Query: `INSERT INTO promotions 
    			(title, content, promo_type, created_at) 
				VALUES ($1, $2, $3, $4)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				faker.Sentence(2),
				faker.Paragraph(1, 2, 5, " "),
				models.GetRandomFromSlice(models.PromotionTypes),
				time.
					Now().
					Add(-time.Duration(faker.Number(1, 365)) * 24 * time.Hour),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedNotificationTemplates() error {
	op := []seeder.SeedOperation{{
		Name: "insert notification_templates",
		Query: `INSERT INTO notification_templates 
    			(type, message_template) 
				VALUES ($1, $2)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			return []interface{}{
				models.GetRandomFromSlice(models.NotificationTypes),
				faker.Sentence(5),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedReminders() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert reminders",
		Query: `INSERT INTO client_next_contact_reminders 
    			(client_id, preferred_communication_channel, message, send_time) 
				VALUES ($1, $2, $3, $4)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			return []interface{}{
				c.ID,
				models.GetRandomFromSlice(models.CommunicationChannels),
				faker.Sentence(5),
				faker.DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0),
				),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedInteractions() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}
	var reminders []models.ClientNextContactReminder
	if err := s.TryGetAllEntities(
		"client_next_contact_reminders",
		&reminders,
		"SELECT id FROM client_next_contact_reminders",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert interactions",
		Query: `INSERT INTO client_interactions 
    			(client_id, time, communication_channel, meeting_location, type, summary, agreements, reminder_id) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			var rid interface{} = nil
			if len(reminders) > 0 && faker.Bool() {
				r := models.GetRandomFromSlice(reminders)
				rid = sql.NullInt32{Int32: r.ID, Valid: true}
			}
			return []interface{}{
				c.ID,
				faker.DateRange(
					time.Now().AddDate(-1, 0, 0),
					time.Now(),
				),
				models.GetRandomFromSlice(models.CommunicationChannels),
				sql.NullString{
					String: faker.Address().City,
					Valid:  true,
				},
				models.GetRandomFromSlice(models.InteractionTypes),
				sql.NullString{
					String: faker.Sentence(5),
					Valid:  true,
				},
				sql.NullString{
					String: faker.Sentence(5),
					Valid:  true,
				},
				rid}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V2Seeder) seedPersonalNotifications() error {
	var clients []models.Client
	if err := s.TryGetAllEntities(
		"clients",
		&clients,
		"SELECT id FROM clients",
	); err != nil {
		return err
	}
	var templates []models.NotificationTemplate
	if err := s.TryGetAllEntities(
		"notification_templates",
		&templates,
		"SELECT id FROM notification_templates",
	); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert personal_notifications",
		Query: `INSERT INTO client_personal_notifications 
    			(client_id, preferred_communication_channel, send_time, template_id) 
				VALUES ($1, $2, $3, $4)
				`,
		Builder: func(faker *gofakeit.Faker) []interface{} {
			c := models.GetRandomFromSlice(clients)
			t := models.GetRandomFromSlice(templates)

			return []interface{}{
				c.ID,
				models.GetRandomFromSlice(models.CommunicationChannels),
				faker.DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0)),
				t.ID,
			}
		},
	}}
	return s.BatchSeed(op)
}
