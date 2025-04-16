package seeder

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

type Hotel struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Address           string `gorm:"not null"`
	CancellationTerms string `gorm:"not null"`
}

func (s *V5Seeder) seedHotels() {
	for i := 0; i < s.count; i++ {
		hotel := Hotel{
			Name:              gofakeit.Company(),
			Address:           gofakeit.Address().Address,
			CancellationTerms: gofakeit.Paragraph(1, 2, 5, " "),
		}
		if err := s.db.Create(&hotel).Error; err != nil {
			log.Println("error seeding hotel:", err)
		}
	}
}

type HotelRoomCategory struct {
	ID            uint    `gorm:"primaryKey"`
	HotelID       uint    `gorm:"not null"`
	Name          string  `gorm:"not null"`
	PricePerNight float64 `gorm:"type:decimal(10,2);not null"`
	MaxGuests     int     `gorm:"not null"`
}

func (s *V5Seeder) seedHotelRoomCategories() {
	var hotels []Hotel
	s.db.Find(&hotels)

	for i := 0; i < s.count; i++ {
		if len(hotels) == 0 {
			log.Fatal("no hotels found, cannot seed hotel room categories")
		}
		category := HotelRoomCategory{
			HotelID:       getRandomFromSlice(hotels).ID,
			Name:          gofakeit.Word() + " Room",
			PricePerNight: gofakeit.Price(50, 1000),
			MaxGuests:     gofakeit.Number(1, 5),
		}
		if err := s.db.Create(&category).Error; err != nil {
			log.Println("error seeding hotel room category:", err)
		}
	}
}

type Amenity struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Description string
}

func (s *V5Seeder) seedAmenities() {
	for i := 0; i < s.count; i++ {
		amenity := Amenity{
			Name:        gofakeit.Word(),
			Description: gofakeit.Paragraph(1, 2, 5, " "),
		}
		if err := s.db.Create(&amenity).Error; err != nil {
			log.Println("error seeding amenity:", err)
		}
	}
}

type HotelNextContactReminder struct {
	ID                            uint      `gorm:"primaryKey"`
	HotelID                       uint      `gorm:"not null"`
	PreferredCommunicationChannel string    `gorm:"not null"`
	Message                       string    `gorm:"not null"`
	SendDate                      time.Time `gorm:"not null"`
}

func (s *V5Seeder) seedHotelNextContactReminders() {
	var hotels []Hotel
	s.db.Find(&hotels)

	for i := 0; i < s.count; i++ {
		if len(hotels) == 0 {
			log.Fatal("no hotels found, cannot seed hotel next contact reminders")
		}
		reminder := HotelNextContactReminder{
			HotelID:                       getRandomFromSlice(hotels).ID,
			PreferredCommunicationChannel: getRandomFromSlice(CommunicationChannels),
			Message:                       gofakeit.Paragraph(1, 2, 5, " "),
			SendDate:                      gofakeit.Date(),
		}
		if err := s.db.Create(&reminder).Error; err != nil {
			log.Println("error seeding hotel next contact reminder:", err)
		}
	}
}

type HotelInteraction struct {
	ID                    uint      `gorm:"primaryKey"`
	HotelID               uint      `gorm:"not null"`
	DateUTC               time.Time `gorm:"not null"`
	CommunicationChannel  string    `gorm:"not null"`
	Type                  string    `gorm:"not null"`
	Summary               string
	Agreements            string
	NextContactReminderID *uint
}

func (s *V5Seeder) seedHotelInteractions() {
	var hotels []Hotel
	var reminders []HotelNextContactReminder
	s.db.Find(&hotels)
	s.db.Find(&reminders)

	for i := 0; i < s.count; i++ {
		if len(hotels) == 0 {
			log.Fatal("no hotels found, cannot seed hotel interactions")
		}

		var nextContactReminderID *uint
		if len(reminders) > 0 && gofakeit.Bool() {
			id := getRandomFromSlice(reminders).ID
			nextContactReminderID = &id
		}

		interaction := HotelInteraction{
			HotelID:               getRandomFromSlice(hotels).ID,
			DateUTC:               gofakeit.Date(),
			CommunicationChannel:  getRandomFromSlice(CommunicationChannels),
			Type:                  getRandomFromSlice(InteractionTypes),
			Summary:               gofakeit.Paragraph(1, 2, 5, " "),
			Agreements:            gofakeit.Paragraph(1, 2, 5, " "),
			NextContactReminderID: nextContactReminderID,
		}
		if err := s.db.Create(&interaction).Error; err != nil {
			log.Println("error seeding hotel interaction:", err)
		}
	}
}

type V5Seeder struct {
	db    *gorm.DB
	count int
}

func NewV5Seeder(db *gorm.DB, count int) *V3Seeder {
	return &V3Seeder{
		db:    db,
		count: count,
	}
}

func (s *V5Seeder) Seed() {
	s.seedHotels()
	s.seedHotelRoomCategories()
	s.seedAmenities()
	s.seedHotelNextContactReminders()
	s.seedHotelInteractions()
}
