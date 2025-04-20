package models

import "github.com/brianvoe/gofakeit/v7"

func GetRandomFromSlice[T any](items []T) T {
	return items[gofakeit.Number(0, len(items)-1)]
}

var CommunicationChannels = []CommunicationChannel{
	CommunicationChannelEmail,
	CommunicationChannelMeeting,
	CommunicationChannelPhone,
	CommunicationChannelTelegram,
	CommunicationChannelWhatsapp,
}

var InteractionTypes = []InteractionType{
	InteractionTypeDiscussion,
	InteractionTypeAgreement,
	InteractionTypeComplaint,
}

var PromotionTypes = []PromotionType{
	PromotionTypeNewTours,
	PromotionTypeLastMinuteTour,
	PromotionTypeEarlyBooking,
}

var BookingStatuses = []BookingStatus{
	BookingStatusDraft,
	BookingStatusPendingSignature,
	BookingStatusPendingPayment,
	BookingStatusPartiallyPaid,
	BookingStatusFullyPaid,
	BookingStatusCancellationRequested,
	BookingStatusCancelled,
}

var AgreementConsentTypes = []AgreementConsentType{
	AgreementConsentTypePersonalData,
	AgreementConsentTypeContractTerms,
	AgreementConsentTypeAd,
}

var ConsentStatuses = []ConsentStatus{
	ConsentStatusGranted,
	ConsentStatusPending,
	ConsentStatusRevoked,
}

var MealsTypes = []MealsType{
	MealsTypeBreakfast,
	MealsTypeHalfBoard,
	MealsTypeFullBoard,
}

var NotificationTypes = []NotificationType{
	NotificationTypePassportExpiry,
	NotificationTypeFlightReminder,
	NotificationTypePaymentReminder,
	NotificationTypeBirthdayPromo,
}

var InsuranceTypes = []InsuranceType{
	InsuranceTypeLostLuggage,
	InsuranceTypeMedical,
}

var PassportTypes = []PassportType{
	PassportTypeInternal,
	PassportTypeForeign,
}
