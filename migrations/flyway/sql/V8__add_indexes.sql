-- +migrate Up
CREATE INDEX idx_clients_city_id ON clients(city_id);
CREATE INDEX idx_bookings_status_created ON bookings(status, created_at);

-- +migrate Down
DROP INDEX idx_clients_city_id;
DROP INDEX idx_bookings_status_created;
