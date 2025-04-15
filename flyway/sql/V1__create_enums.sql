create type communication_channel as enum (
    'phone',
    'telegram',
    'whatsapp',
    'email',
    'meeting'
);

create type interaction_type as enum (
    'discussion',
    'agreement',
    'complaint'
);

create type promotion_type as enum (
    'new_tours',
    'last_minute_tour',
    'early_booking'
);

create type booking_status as enum (
    'draft',
    'pending_signature',
    'pending_payment',
    'partially_paid',
    'fully_paid',
    'cancellation_requested',
    'canceled'
);

create type agreement_consent_type as enum (
    'personal_data',
    'contract_terms',
    'ad'
);

create type consent_status as enum (
    'granted',
    'pending',
    'revoked'
);

create type meals_type as enum (
  'breakfast',
  'half_board',
  'full_board'
);

create type notification_type as enum (
  'passport_expiry',
  'flight_reminder',
  'payment_reminder',
  'birthday_promo'
);

create type insurance_type as enum (
  'medical',
  'lost_luggage'
);

create type passport_type as enum (
  'internal',
  'foreign'
);
