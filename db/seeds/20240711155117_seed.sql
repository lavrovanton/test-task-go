-- +goose Up
-- +goose StatementBegin
INSERT INTO services (name, type, payment_Type, price) VALUES 
('Service 1', 'VDS', 'year', 1000),
('Service 2', 'VDS', 'month', 299),
('Service 3', 'Hosting', 'year', 700),
('Service 4', 'Hosting', 'half-year', 600),
('Service 5', 'Dedicated Server', 'year', 500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM services WHERE name IN ('Service 1', 'Service 2', 'Service 3', 'Service 4', 'Service 5');
-- +goose StatementEnd
