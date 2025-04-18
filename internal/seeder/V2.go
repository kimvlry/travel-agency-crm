package seeder

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
	"log"
	"time"
)

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

func (s *V2Seeder) seedCountries() error {
	for i := 0; i < s.count; i++ {
		country := Country{
			Name: gofakeit.Country(),
		}
		if err := s.db.Create(&country).Error; err != nil {
			return fmt.Errorf("error seeding country: %w", err)
		}
	}
	return nil
}

type City struct {
	ID             uint    `gorm:"primaryKey"`
	Name           string  `gorm:"not null"`
	CountryID      uint    `gorm:"not null"`
	Country        Country `gorm:"foreignKey:CountryID"`
	TimezoneOffset int     `gorm:"not null"`
}

func (s *V2Seeder) seedCities() error {
	var countries []Country
	s.db.Find(&countries)

	for i := 0; i < s.count; i++ {
		if len(countries) == 0 {
			return fmt.Errorf("no countries found, cannot seed cities")
		}

		city := City{
			Name:           gofakeit.City(),
			CountryID:      getRandomFromSlice(countries).ID,
			TimezoneOffset: gofakeit.Number(-12, 12),
		}
		if err := s.db.Create(&city).Error; err != nil {
			return fmt.Errorf("error seeding city: %w", err)
		}
	}
	return nil
}

type Client struct {
	ID            uint      `gorm:"primaryKey"`
	FullName      string    `gorm:"not null"`
	Phone         string    `gorm:"not null"`
	Email         string    `gorm:"not null"`
	BirthDate     time.Time `gorm:"not null"`
	CityID        uint      `gorm:"not null"`
	City          City      `gorm:"foreignKey:CityID"`
	IsBlacklisted bool      `gorm:"default:false"`
}

func (s *V2Seeder) seedClients() error {
	var cities []City
	s.db.Find(&cities)

	for i := 0; i < s.count; i++ {
		if len(cities) == 0 {
			return fmt.Errorf("no cities found, cannot seed clients")
		}

		client := Client{
			FullName: gofakeit.Name(),
			Phone:    gofakeit.Phone(),
			Email:    gofakeit.Email(),
			BirthDate: gofakeit.
				DateRange(
					time.Now().AddDate(-60, 0, 0),
					time.Now().AddDate(-18, 0, 0)),
			CityID: cities[gofakeit.Number(0, len(cities)-1)].ID,
		}
		if err := s.db.Create(&client).Error; err != nil {
			return fmt.Errorf("error seeding client: %w", err)
		}
	}
	return nil
}

type Ban struct {
	ID        uint
	ClientID  uint
	Client    Client
	BanReason string
}

func (s *V2Seeder) seedBans() error {
	var clients []Client
	s.db.Find(&clients)

	if len(clients) == 0 {
		return fmt.Errorf("no clients found, cannot seed bans")
	}

	for i := 0; i < s.count; i++ {
		client := clients[gofakeit.Number(0, len(clients)-1)]

		err := s.db.Transaction(func(tx *gorm.DB) error {
			ban := Ban{
				ClientID:  client.ID,
				BanReason: gofakeit.Sentence(10),
			}
			if err := tx.Create(&ban).Error; err != nil {
				return err
			}
			if err := tx.
				Model(&Client{}).
				Where("id = ?", client.ID).
				Update("is_blacklisted", true).Error; err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			log.Println("error in transaction seeding ban and updating client: ", err)
			return err
		}
	}
	return nil
}

type Passport struct {
	ID             uint
	ClientID       uint
	Client         Client
	Type           string    `gorm:"not null"`
	Number         string    `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	IssueDate      time.Time `gorm:"not null"`
}

func (s *V2Seeder) seedPassports() error {
	var clients []Client
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}

	if len(clients) == 0 {
		return fmt.Errorf("no clients found, cannot seed passports")
	}

	var passports []Passport
	for _, client := range clients {
		for _, passportType := range PassportTypes {
			passport := Passport{
				ClientID: client.ID,
				Type:     passportType,
				Number:   gofakeit.UUID(),
				ExpirationDate: gofakeit.
					DateRange(
						time.Now().AddDate(5, 0, 0),
						time.Now().AddDate(10, 0, 0)),
				IssueDate: gofakeit.
					DateRange(
						time.Now().AddDate(-10, 0, 0),
						time.Now().AddDate(-1, 0, 0)),
			}
			passports = append(passports, passport)
		}
	}

	if err := s.db.Create(&passports).Error; err != nil {
		return fmt.Errorf("failed to seed passports: %w", err)
	}
	return nil
}

type ClientNextContactReminder struct {
	ID                            uint
	ClientID                      uint
	Client                        Client
	PreferredCommunicationChannel string    `gorm:"not null"`
	Message                       string    `gorm:"not null"`
	SendTime                      time.Time `gorm:"not null"`
}

func (s *V2Seeder) seedClientNextContactReminders() error {
	var clients []Client
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}

	if len(clients) == 0 {
		return fmt.Errorf("no clients found for contact reminders")
	}

	for i := 0; i < s.count; i++ {
		reminder := ClientNextContactReminder{
			ClientID:                      getRandomFromSlice(clients).ID,
			PreferredCommunicationChannel: getRandomFromSlice(CommunicationChannels),
			Message:                       gofakeit.Sentence(6),
			SendTime: gofakeit.
				DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0)),
		}
		if err := s.db.Create(&reminder).Error; err != nil {
			log.Println("error seeding contact reminder:", err)
		}
	}
	return nil
}

type ClientInteraction struct {
	ID                   uint
	ClientID             uint
	Client               Client
	Time                 time.Time `gorm:"not null"`
	CommunicationChannel string    `gorm:"not null"`
	MeetingLocation      string
	Type                 string `gorm:"not null"`
	Summary              string
	Agreements           string
	ReminderID           *uint
	Reminder             *ClientNextContactReminder `gorm:"foreignKey:ReminderID"`
}

func (s *V2Seeder) seedClientInteractions() error {
	var clients []Client
	var reminders []ClientNextContactReminder
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}
	if err := s.db.Find(&reminders).Error; err != nil {
		return fmt.Errorf("failed to fetch reminders: %w", err)
	}

	if len(clients) == 0 {
		return fmt.Errorf("no clients found for interactions")
	}

	for i := 0; i < s.count; i++ {
		interaction := ClientInteraction{
			ClientID: getRandomFromSlice(clients).ID,
			Time: gofakeit.
				DateRange(
					time.Now().AddDate(-1, 0, 0),
					time.Now()),
			CommunicationChannel: getRandomFromSlice(CommunicationChannels),
			MeetingLocation:      gofakeit.Address().City,
			Type:                 getRandomFromSlice(InteractionTypes),
			Summary:              gofakeit.Sentence(10),
			Agreements:           gofakeit.Sentence(6),
		}

		if len(reminders) > 0 && gofakeit.Bool() {
			rem := reminders[gofakeit.Number(0, len(reminders)-1)]
			interaction.ReminderID = &rem.ID
		}

		if err := s.db.Create(&interaction).Error; err != nil {
			log.Println("error seeding interaction:", err)
		}
	}
	return nil
}

type NotificationTemplate struct {
	ID              uint
	Type            string `gorm:"not null"`
	MessageTemplate string
	PromoID         *uint
}

func (s *V2Seeder) seedNotificationTemplates() error {
	var promos []Promotion
	if err := s.db.Find(&promos).Error; err != nil {
		return fmt.Errorf("failed to fetch promotions: %w", err)
	}

	for i := 0; i < s.count; i++ {
		template := NotificationTemplate{
			Type:            getRandomFromSlice(NotificationTypes),
			MessageTemplate: gofakeit.Sentence(8),
		}

		if len(promos) > 0 && gofakeit.Bool() {
			template.PromoID = &promos[gofakeit.Number(0, len(promos)-1)].ID
		}

		if err := s.db.Create(&template).Error; err != nil {
			log.Println("error seeding notification template:", err)
		}
	}
	return nil
}

type ClientPersonalNotification struct {
	ID                            uint
	ClientID                      uint
	Client                        Client
	PreferredCommunicationChannel string `gorm:"not null"`
	TemplateID                    *uint
	SendTime                      time.Time
}

func (s *V2Seeder) seedClientPersonalNotifications() error {
	var clients []Client
	var templates []NotificationTemplate
	if err := s.db.Find(&clients).Error; err != nil {
		return fmt.Errorf("failed to fetch clients: %w", err)
	}
	if err := s.db.Find(&templates).Error; err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	if len(clients) == 0 {
		return fmt.Errorf("no clients found for personal notifications")
	}

	for i := 0; i < s.count; i++ {
		notif := ClientPersonalNotification{
			ClientID:                      getRandomFromSlice(clients).ID,
			PreferredCommunicationChannel: getRandomFromSlice(CommunicationChannels),
			SendTime: gofakeit.
				DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0)),
		}

		if len(templates) > 0 && gofakeit.Bool() {
			notif.TemplateID = &templates[gofakeit.Number(0, len(templates)-1)].ID
		}

		if err := s.db.Create(&notif).Error; err != nil {
			log.Println("error seeding personal notification:", err)
		}
	}
	return nil
}

type Promotion struct {
	ID        uint
	Title     string
	Content   string
	PromoType string `gorm:"not null"`
	CreatedAt time.Time
}

func (s *V2Seeder) seedPromotions() error {
	for i := 0; i < s.count; i++ {
		promo := Promotion{
			Title:     gofakeit.Sentence(3),
			Content:   gofakeit.Paragraph(1, 2, 5, " "),
			PromoType: getRandomFromSlice(PromotionTypes),
			CreatedAt: time.Now().Add(-time.Duration(gofakeit.Number(1, 365)) * 24 * time.Hour),
		}
		if err := s.db.Create(&promo).Error; err != nil {
			log.Println("error seeding promotion:", err)
		}
	}
	return nil
}

type V2Seeder struct {
	db    *gorm.DB
	count int
}

func NewV2Seeder(db *gorm.DB, count int) *V2Seeder {
	return &V2Seeder{
		db:    db,
		count: count,
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
	if err := s.seedClientNextContactReminders(); err != nil {
		return err
	}
	if err := s.seedClientInteractions(); err != nil {
		return err
	}
	if err := s.seedNotificationTemplates(); err != nil {
		return err
	}
	if err := s.seedClientPersonalNotifications(); err != nil {
		return err
	}
	if err := s.seedPromotions(); err != nil {
		return err
	}
	return nil
}
