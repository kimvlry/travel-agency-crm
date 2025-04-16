package seeder

import "github.com/brianvoe/gofakeit/v7"

func getRandomFromSlice[T any](items []T) T {
	return items[gofakeit.Number(0, len(items)-1)]
}

var CommunicationChannels = []string{
	"phone",
	"telegram",
	"whatsapp",
	"email",
	"meeting",
}

var InteractionTypes = []string{
	"discussion",
	"agreement",
	"complaint",
}

var PromotionTypes = []string{
	"new_tours",
	"last_minute_tour",
	"early_booking",
}

var BookingStatuses = []string{
	"draft",
	"pending_signature",
	"pending_payment",
	"partially_paid",
	"fully_paid",
	"cancellation_requested",
	"canceled",
}

var AgreementConsentTypes = []string{
	"personal_data",
	"contract_terms",
	"ad",
}

var ConsentStatuses = []string{
	"granted",
	"pending",
	"revoked",
}

var MealsTypes = []string{
	"breakfast",
	"half_board",
	"full_board",
}

var NotificationTypes = []string{
	"passport_expiry",
	"flight_reminder",
	"payment_reminder",
	"birthday_promo",
}

var InsuranceTypes = []string{
	"medical",
	"lost_luggage",
}

var PassportTypes = []string{
	"internal",
	"foreign",
}
