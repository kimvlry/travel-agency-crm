package impl

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"

	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V5Seeder struct {
	*seeder.BaseSeeder
}

func NewV5Seeder(tx *sqlx.Tx, seedCount int) *V5Seeder {
	return &V5Seeder{BaseSeeder: seeder.NewBaseSeeder(tx, seedCount)}
}

func (s *V5Seeder) Seed() error {
	if err := s.seedHotels(); err != nil {
		return err
	}
	if err := s.seedHotelRoomCategories(); err != nil {
		return err
	}
	if err := s.seedAmenities(); err != nil {
		return err
	}
	if err := s.seedHotelNextContactReminders(); err != nil {
		return err
	}
	if err := s.seedHotelInteractions(); err != nil {
		return err
	}
	return nil
}

func (s *V5Seeder) seedHotels() error {
	op := []seeder.SeedOperation{{
		Name:  "insert hotels",
		Query: `INSERT INTO hotels (name, address, cancellation_terms) VALUES ($1, $2, $3)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				f.Company(),
				f.Address().Address,
				f.Paragraph(1, 2, 5, " "),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V5Seeder) seedHotelRoomCategories() error {
	var hotels []models.Hotel
	if err := s.TryGetAllEntities("hotels", &hotels, "SELECT id FROM hotels"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert hotel_room_categories",
		Query: `INSERT INTO hotel_room_categories (hotel_id, name, price_per_night, max_guests) 
				VALUES ($1, $2, $3, $4)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			hotel := models.GetRandomFromSlice(hotels)
			return []interface{}{
				hotel.ID,
				f.Word() + " Room",
				f.Price(50, 1000),
				f.Number(1, 5),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V5Seeder) seedAmenities() error {
	op := []seeder.SeedOperation{{
		Name:  "insert amenities",
		Query: `INSERT INTO amenities (name, description) VALUES ($1, $2)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			return []interface{}{
				f.Word(),
				f.Paragraph(1, 2, 5, " "),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V5Seeder) seedHotelNextContactReminders() error {
	var hotels []models.Hotel
	if err := s.TryGetAllEntities("hotels", &hotels, "SELECT id FROM hotels"); err != nil {
		return err
	}
	op := []seeder.SeedOperation{{
		Name: "insert hotel_next_contact_reminders",
		Query: `INSERT INTO hotel_next_contact_reminders (hotel_id, preferred_communication_channel, message, send_date) 
				VALUES ($1, $2, $3, $4)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			hotel := models.GetRandomFromSlice(hotels)
			channel := models.GetRandomFromSlice(models.CommunicationChannels)
			return []interface{}{
				hotel.ID,
				channel,
				f.Paragraph(1, 2, 5, " "),
				f.Date(),
			}
		},
	}}
	return s.BatchSeed(op)
}

func (s *V5Seeder) seedHotelInteractions() error {
	var hotels []models.Hotel
	var reminders []models.HotelNextContactReminder
	if err := s.TryGetAllEntities("hotels", &hotels, "SELECT id FROM hotels"); err != nil {
		return err
	}
	query := "SELECT id FROM hotel_next_contact_reminders"
	if err := s.TryGetAllEntities("hotel_next_contact_reminders", &reminders, query); err != nil {
		return err
	}

	op := []seeder.SeedOperation{{
		Name: "insert hotel_interactions",
		Query: `INSERT INTO hotel_interactions (hotel_id, date_utc, communication_channel, type, summary, agreements, next_contact_reminder_id) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		Builder: func(f *gofakeit.Faker) []interface{} {
			var reminderID interface{} = nil
			if len(reminders) > 0 && f.Bool() {
				reminder := models.GetRandomFromSlice(reminders)
				reminderID = reminder.ID
			}
			return []interface{}{
				models.GetRandomFromSlice(hotels).ID,
				f.Date(),
				models.GetRandomFromSlice(models.CommunicationChannels),
				models.GetRandomFromSlice(models.InteractionTypes),
				f.Paragraph(1, 2, 5, " "),
				f.Paragraph(1, 2, 5, " "),
				reminderID,
			}
		},
	}}
	return s.BatchSeed(op)
}
