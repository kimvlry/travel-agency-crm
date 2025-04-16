create table hotels (
    id serial primary key,
    name varchar(255) not null,
    address text not null,
    cancellation_terms text not null
);

create table hotel_room_categories (
    id serial primary key,
    hotel_id integer references hotels(id) on delete cascade,
    name varchar(100) not null,
    price_per_night decimal(10,2) not null,
    max_guests integer not null
);

create table amenities (
    id serial primary key,
    name varchar(255) not null unique,
    description text
);

create table hotel_next_contact_reminders (
    id serial primary key,
    hotel_id integer references hotels(id) on delete cascade,
    preferred_communication_channel communication_channel not null,
    message text not null,
    send_date timestamptz not null
);

create table hotel_interactions (
    id serial primary key,
    hotel_id integer references hotels(id) on delete cascade,
    date_utc timestamptz not null,
    communication_channel communication_channel not null,
    type interaction_type not null,
    summary text,
    agreements text,
    next_contact_reminder integer references hotel_next_contact_reminders(id) on delete set null
);
