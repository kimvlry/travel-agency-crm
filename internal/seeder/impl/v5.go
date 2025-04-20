package impl

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
)

type V5Seeder struct {
	*seeder.BaseSeeder
}

func NewV5Seeder(tx *sqlx.Tx, seedCount int) *V5Seeder {
	return &V5Seeder{
		BaseSeeder: seeder.NewBaseSeeder(tx, seedCount),
	}
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
	for i := 0; i < s.SeedCount; i++ {
		hotel := models.Hotel{
			Name:              gofakeit.Company(),
			Address:           gofakeit.Address().Address,
			CancellationTerms: gofakeit.Paragraph(1, 2, 5, " "),
		}
		if err := s.db.Create(&hotel).Error; err != nil {
			return fmt.Errorf("error seeding hotel: %w", err)
		}
	}
	return nil
}

func (s *V5Seeder) seedHotelRoomCategories() error {
	var hotels []models.Hotel
	s.db.Find(&hotels)

	for i := 0; i < s.SeedCount; i++ {
		if len(hotels) == 0 {
			return fmt.Errorf("no hotels found, cannot seed hotel room categories")
		}
		category := models.HotelRoomCategory{
			HotelID:       seeder.getRandomFromSlice(hotels).ID,
			Name:          gofakeit.Word() + " Room",
			PricePerNight: gofakeit.Price(50, 1000),
			MaxGuests:     gofakeit.Number(1, 5),
		}
		if err := s.db.Create(&category).Error; err != nil {
			return fmt.Errorf("error seeding hotel room category: %w", err)
		}
	}
	return nil
}

func (s *V5Seeder) seedAmenities() error {
	for i := 0; i < s.SeedCount; i++ {
		amenity := models.Amenity{
			Name:        gofakeit.Word(),
			Description: gofakeit.Paragraph(1, 2, 5, " "),
		}
		if err := s.db.Create(&amenity).Error; err != nil {
			return fmt.Errorf("error seeding amenity: %w", err)
		}
	}
	return nil
}

func (s *V5Seeder) seedHotelNextContactReminders() error {
	var hotels []models.Hotel
	s.db.Find(&hotels)

	for i := 0; i < s.SeedCount; i++ {
		if len(hotels) == 0 {
			return fmt.Errorf("no hotels found, cannot seed hotel next contact reminders")
		}
		reminder := models.HotelNextContactReminder{
			HotelID:                       seeder.getRandomFromSlice(hotels).ID,
			PreferredCommunicationChannel: seeder.getRandomFromSlice(models.CommunicationChannels),
			Message:                       gofakeit.Paragraph(1, 2, 5, " "),
			SendDate:                      gofakeit.Date(),
		}
		if err := s.db.Create(&reminder).Error; err != nil {
			return fmt.Errorf("error seeding hotel next contact reminder: %w", err)
		}
	}
	return nil
}

func (s *V5Seeder) seedHotelInteractions() error {
	var hotels []models.Hotel
	var reminders []models.HotelNextContactReminder
	s.db.Find(&hotels)
	s.db.Find(&reminders)

	for i := 0; i < s.SeedCount; i++ {
		if len(hotels) == 0 {
			return fmt.Errorf("no hotels found, cannot seed hotel interactions")
		}

		var nextContactReminderID *uint
		if len(reminders) > 0 && gofakeit.Bool() {
			id := seeder.getRandomFromSlice(reminders).ID
			nextContactReminderID = &id
		}

		interaction := models.HotelInteraction{
			HotelID:               seeder.getRandomFromSlice(hotels).ID,
			DateUtc:               gofakeit.Date(),
			CommunicationChannel:  seeder.getRandomFromSlice(models.CommunicationChannels),
			Type:                  seeder.getRandomFromSlice(models.InteractionTypes),
			Summary:               gofakeit.Paragraph(1, 2, 5, " "),
			Agreements:            gofakeit.Paragraph(1, 2, 5, " "),
			NextContactReminderID: nextContactReminderID,
		}
		if err := s.db.Create(&interaction).Error; err != nil {
			return fmt.Errorf("error seeding hotel interaction: %w", err)
		}
	}
	return nil
}
