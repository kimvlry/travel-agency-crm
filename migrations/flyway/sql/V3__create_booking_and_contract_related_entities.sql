create table tours (
    id serial primary key,
    title varchar(255) not null unique,
    price_eur decimal(10,2) not null,
    quota integer not null,
    meals_type meals_type,
    is_last_minute boolean default false,
    last_minute_approved boolean,
    base_duration_days integer not null
);

create table bookings (
    id serial primary key,
    tour_id integer references tours(id),
    status booking_status default 'draft',
    contract_number varchar(50) unique,
    payment_links text,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

create table booking_agreements (
    id serial primary key,
    client_id integer references clients(id),
    booking_id integer references bookings(id),
    signed_date timestamptz not null
);

create table assignees (
    id serial primary key,
    client_id integer references clients(id)
);

create table contract_templates (
    id serial primary key,
    name varchar(100) not null,
    type varchar(50) not null,
    content text not null,
    validity_period_days integer not null
);

create table contracts (
    id serial primary key,
    assignee_id integer references assignees(id),
    template_id integer references contract_templates(id),
    issue_date timestamptz not null,
    sign_date timestamptz,
    status consent_status not null,
    total_price decimal(10,2) not null
);

create table consent_templates (
    id serial primary key,
    name varchar(100) not null,
    type agreement_consent_type not null,
    content text not null
);

create table agreement_consents (
    id serial primary key,
    assignee_id integer references assignees(id),
    contract_id integer references contracts(id),
    template_id integer references consent_templates(id) on delete set null,
    date date not null,
    status consent_status not null
);

create table payment_links (
    id serial primary key,
    booking_id integer references bookings(id) on delete cascade,
    url varchar(255) not null,
    qr_code varchar(255) not null,
    status varchar(50) not null
);
