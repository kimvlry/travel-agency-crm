create table countries (
    id serial primary key,
    name varchar(255) not null
);

create table cities (
    id serial primary key,
    name varchar(255) not null,
    country_id integer not null references countries(id) on delete cascade,
    timezone_offset integer not null
);

create table clients (
    id serial primary key,
    full_name varchar(255) not null,
    phone varchar(20) not null,
    email varchar(255) not null,
    birth_date date not null,
    city_id integer not null references cities(id),
    is_blacklisted boolean default false
);

create table bans (
    id serial primary key,
    client_id integer references clients(id) on delete cascade,
    ban_reason text
);

create table passports (
    id serial primary key,
    client_id integer references clients(id) on delete cascade,
    type passport_type not null,
    number varchar(50) not null,
    expiration_date date not null,
    issue_date date not null,
    constraint unique_client_passport unique (client_id, type)
);

create table client_next_contact_reminders (
    id serial primary key,
    client_id integer references clients(id) on delete cascade,
    preferred_communication_channel communication_channel not null,
    message text not null,
    send_time timestamptz not null
);

create table client_interactions (
    id serial primary key,
    client_id integer references clients(id) on delete cascade,
    time timestamptz not null,
    communication_channel communication_channel not null,
    meeting_location varchar(255),
    type interaction_type not null,
    summary text,
    agreements text,
    reminder_id integer references client_next_contact_reminders(id) on delete set null
);

create table notification_templates (
    id serial primary key,
    type notification_type not null,
    message_template text,
    promo_id integer
);

create table client_personal_notification (
    id serial primary key,
    client_id integer references clients(id) on delete cascade,
    preferred_communication_channel communication_channel not null,
    template_id integer references notification_templates(id) on delete set null,
    send_time timestamptz not null
);



create table promotions (
    id serial primary key,
    title varchar(255) not null,
    content text not null,
    promo_type promotion_type not null,
    created_at timestamptz default current_timestamp
);
