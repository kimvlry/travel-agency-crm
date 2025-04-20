package impl

import (
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"time"
	"travel-agency-seeder/internal/models"
	"travel-agency-seeder/internal/seeder"
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

func (s *V2Seeder) seedCountries() error {
	for i := 0; i < s.SeedCount; i++ {
		country := models.Country{
			Name: gofakeit.Country(),
		}
		query := `insert into countries (name) values ($1)`
		if err := s.TryExecute(
			"insert countries",
			query,
			country.Name,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedCities() error {
	var countries []models.Country
	if err := s.TryGetAllEntities("countries", &countries, "select id from countries"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	for i := 0; i < s.SeedCount; i++ {
		query := `
            insert into cities
            (name, country_id, timezone_offset)
            values ($1, $2, $3)`

		city := models.City{
			Name:           gofakeit.City(),
			CountryID:      models.GetRandomFromSlice(countries).ID,
			TimezoneOffset: int32(gofakeit.Number(-12, 12)),
		}

		if err := s.TryExecute(
			"insert cities",
			query,
			city.Name,
			city.CountryID,
			city.TimezoneOffset,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedClients() error {
	var cities []models.City
	if err := s.TryGetAllEntities("cities", &cities, "select * from cities"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}
	for i := 0; i < s.SeedCount; i++ {
		query := `insert into clients (full_name, phone, email, birth_date, city_id) 
                values ($1, $2, $3, $4, $5)`
		client := models.Client{
			FullName: gofakeit.Name(),
			Phone:    gofakeit.Phone(),
			Email:    gofakeit.Email(),
			BirthDate: gofakeit.
				DateRange(
					time.Now().AddDate(-60, 0, 0),
					time.Now().AddDate(-18, 0, 0)),
			CityID: cities[gofakeit.Number(0, len(cities)-1)].ID,
		}
		if err := s.TryExecute(
			"insert clients",
			query,
			client.FullName,
			client.Phone,
			client.Email,
			client.BirthDate,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedBans() error {
	var clients []models.Client
	if err := s.TryGetAllEntities("clients", &clients, "SELECT * FROM clients"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	const (
		insertBanQuery = `
            INSERT INTO bans 
                (client_id, ban_reason) 
            VALUES 
                (:client_id, :ban_reason)`

		updateClientQuery = `
            UPDATE clients 
            SET is_blacklisted = true 
            WHERE id = :client_id`
	)

	for i := 0; i < s.SeedCount; i++ {
		client := clients[gofakeit.Number(0, len(clients)-1)]

		ban := models.Ban{
			ClientID:  client.ID,
			BanReason: gofakeit.Sentence(10),
		}

		if err := s.TryExecute(
			"insert ban",
			insertBanQuery,
			ban.ClientID,
			ban.BanReason,
		); err != nil {
			s.Logger.Printf("Error creating ban: %v", err)
			return fmt.Errorf("ban creation failed: %w", err)
		}

		if err := s.TryExecute(
			"client update",
			updateClientQuery,
			client.ID,
		); err != nil {
			s.Logger.Printf("Error updating client: %v", err)
			return fmt.Errorf("client update failed: %w", err)
		}
	}
	return nil
}

func (s *V2Seeder) seedPassports() error {
	var clients []models.Client
	query := "select * from clients"
	if err := s.TryGetAllEntities("clients", &clients, query); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	var passports []models.Passport
	for _, client := range clients {
		for _, passportType := range models.PassportTypes {
			passport := models.Passport{
				ClientID: client.ID,
				Type:     models.PassportType(passportType),
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

	query = `
        INSERT INTO passports 
            (client_id, type, number, expiration_date, issue_date) 
        VALUES 
            (:client_id, :type, :number, :expiration_date, :issue_date)`

	if _, err := s.Tx.NamedExec(query, passports); err != nil {
		s.Logger.Println("failed to seed passports: %w", err)
		return err
	}
	return nil
}

func (s *V2Seeder) seedClientNextContactReminders() error {
	var clients []models.Client
	query := "select * from clients"
	if err := s.TryGetAllEntities("clients", &clients, query); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	for i := 0; i < s.SeedCount; i++ {
		reminder := models.ClientNextContactReminder{
			ClientID:                      models.GetRandomFromSlice(clients).ID,
			PreferredCommunicationChannel: models.CommunicationChannel(models.GetRandomFromSlice(models.CommunicationChannels)),
			Message:                       gofakeit.Sentence(6),
			SendTime: gofakeit.
				DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0)),
		}

		query := `insert into client_next_contact_reminders 
    		(client_id, preferred_communication_channel, message, send_time) 
			values ($1, $2, $3, $4)`

		if err := s.TryExecute("insert client_next_contact_reminders",
			query,
			reminder.ClientID,
			reminder.PreferredCommunicationChannel,
			reminder.Message,
			reminder.SendTime,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedClientInteractions() error {
	var clients []models.Client
	query := "select * from clients"
	if err := s.TryGetAllEntities("clients", &clients, query); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	var reminders []models.ClientNextContactReminder
	query = "select * from client_next_contact_reminders"
	if err := s.TryGetAllEntities("client_next_contact_reminders", &reminders, query); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	for i := 0; i < s.SeedCount; i++ {
		interaction := models.ClientInteraction{
			ClientID: models.GetRandomFromSlice(clients).ID,
			Time: gofakeit.
				DateRange(
					time.Now().AddDate(-1, 0, 0),
					time.Now()),
			CommunicationChannel: models.CommunicationChannel(models.GetRandomFromSlice(models.CommunicationChannels)),
			MeetingLocation: sql.NullString{
				String: gofakeit.Address().City,
				Valid:  true,
			},
			Type: models.InteractionType(models.GetRandomFromSlice(models.InteractionTypes)),
			Summary: sql.NullString{
				String: gofakeit.Sentence(8),
				Valid:  true,
			},
			Agreements: sql.NullString{
				String: gofakeit.Sentence(8),
				Valid:  true,
			},
		}

		if len(reminders) > 0 && gofakeit.Bool() {
			randomIndex := gofakeit.Number(0, len(reminders)-1)
			interaction.ReminderID = sql.NullInt32{
				Int32: reminders[randomIndex].ID,
				Valid: true,
			}
		}

		query := `insert into client_interactions 
    		(client_id, time, communication_channel, meeting_location, type, summary, agreements, reminder_id)
			values ($1, $2, $3, $4, $5, $6, $7)`

		if err := s.TryExecute(
			"insert client_interactions",
			query,
			interaction.ClientID,
			interaction.Time,
			interaction.CommunicationChannel,
			interaction.MeetingLocation,
			interaction.Type,
			interaction.Summary,
			interaction.Agreements,
			interaction.ReminderID); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedNotificationTemplates() error {
	var promos []models.Promotion
	if err := s.TryGetAllEntities("promotions", &promos, "select * from promotions"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	for i := 0; i < s.SeedCount; i++ {
		template := models.NotificationTemplate{
			Type: models.NotificationType(models.GetRandomFromSlice(models.NotificationTypes)),
			MessageTemplate: sql.NullString{
				String: gofakeit.Sentence(8),
				Valid:  true,
			},
		}

		query := `insert into notification_templates (type, message_template, promo_id)  
            values ($1, $2, $3)`

		if err := s.TryExecute(
			"insert notification_templates",
			query,
			template.Type,
			template.MessageTemplate,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedClientPersonalNotifications() error {
	var clients []models.Client
	var templates []models.NotificationTemplate

	if err := s.TryGetAllEntities("clients", &clients, "select * from clients"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}
	if err := s.TryGetAllEntities("notification_templates", &templates,
		"select * from notification_templates"); err != nil {
		s.Logger.Println(err.Error())
		return err
	}

	for i := 0; i < s.SeedCount; i++ {
		notif := models.ClientPersonalNotification{
			ClientID:                      models.GetRandomFromSlice(clients).ID,
			PreferredCommunicationChannel: models.CommunicationChannel(models.GetRandomFromSlice(models.CommunicationChannels)),
			SendTime: gofakeit.
				DateRange(
					time.Now(),
					time.Now().AddDate(0, 1, 0)),
		}

		if len(templates) > 0 && gofakeit.Bool() {
			notif.TemplateID = templates[gofakeit.Number(0, len(templates)-1)].ID
		}

		query := `insert into client_personal_notifications (client_id, preferred_communication_channel, template_id, send_time)
			values ($1, $2, $3, $4)`
		if err := s.TryExecute(
			"insert client_personal_notifications",
			query,
			notif.ClientID,
			notif.PreferredCommunicationChannel,
			notif.SendTime,
			notif.TemplateID,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *V2Seeder) seedPromotions() error {
	for i := 0; i < s.SeedCount; i++ {
		promo := models.Promotion{
			Title:     gofakeit.Sentence(3),
			Content:   gofakeit.Paragraph(1, 2, 5, " "),
			PromoType: models.PromotionType(models.GetRandomFromSlice(models.PromotionTypes)),
			CreatedAt: time.Now().Add(-time.Duration(gofakeit.Number(1, 365)) * 24 * time.Hour),
		}

		query := `insert into promotions (title, content, promo_type, created_at) 
            values ($1, $2, $3, $4)`

		if err := s.TryExecute(
			"insert promotions",
			query,
			promo.Title,
			promo.Content,
			promo.PromoType,
			promo.CreatedAt,
		); err != nil {
			s.Logger.Println(err.Error())
			return err
		}
	}
	return nil
}
