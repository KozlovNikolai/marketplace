-- +goose Up
-- +goose StatementBegin
INSERT INTO orders (id, user_id, state_id, total_amount, created_at) VALUES
(1, 1, 2, 0, '2023-02-26 22:42:04.111111'),
(2, 1, 2, 0, '2023-02-26 22:42:04.111111'),
(3, 3, 2, 0, '2023-02-26 22:42:04.111111');	

SELECT setval(pg_get_serial_sequence('orders', 'id'), max(id)) FROM orders;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM orders;
-- +goose StatementEnd