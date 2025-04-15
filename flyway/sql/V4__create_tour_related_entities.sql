create table tour_iterations (
    id serial primary key,
    tour_id integer references tours(id) on delete cascade,
    start_date date not null,
    end_date date not null,
    constraint unique_tour_id_start_date unique (tour_id, start_date)
);

create table tour_routes (
    id serial primary key,
    tour_id integer references tours(id) on delete cascade
);

create table route_points (
    id serial primary key,
    route_id integer references tour_routes(id) on delete cascade,
    city_id integer references cities(id) on delete set null,
    name varchar(255) not null,
    address text not null,
    duration_time interval,
    in_route_order_position integer not null
);

create table transport_services (
    id serial primary key,
    company varchar(255) not null,
    model varchar(255)
);

create table transfers (
    id serial primary key,
    tour_id integer references tours(id) on delete cascade,
    transport_id integer references transport_services(id) on delete set null,
    departure_point integer references route_points(id) on delete cascade,
    arrival_point integer references route_points(id) on delete cascade,
    departure_time timestamptz not null
);

create table organizers (
    id serial primary key,
    name varchar(255) not null,
    phone varchar(20) not null,
    email varchar(255) not null
);

create table excursions (
    id serial primary key,
    tour_id integer references tours(id) on delete cascade,
    organizer_id integer references organizers(id) on delete set null,
    name varchar(255) not null,
    meeting_location integer references route_points(id) on delete cascade,
    meeting_time timestamptz not null
);

create table insurance_companies (
    id serial primary key,
    name varchar(255) not null,
    phone varchar(20) not null,
    email varchar(255) not null
);

create table insurances (
    id serial primary key,
    tour_id integer references tours(id) on delete cascade,
    insurance_company_id integer references insurance_companies(id) on delete set null,
    coverage_type insurance_type not null
);
