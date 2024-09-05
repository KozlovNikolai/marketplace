-- +goose Up
-- +goose StatementBegin
INSERT INTO items (id, product_id, quantity, total_price, order_id) VALUES
(1, 1, 23, 0, 1),
(2, 2, 5, 0, 2),
(3, 3, 7, 0, 3),
(4, 4, 9, 0, 1),
(5, 5, 32, 0, 2),
(6, 6, 65, 0, 3),
(7, 7, 2, 0, 1),
(8, 8, 1, 0, 2),
(9, 9, 76, 0, 3),
(10, 10, 28, 0, 1),
(11, 11, 90, 0, 2),
(12, 12, 2000, 0, 3),
(13, 1, 23, 0, 1),
(14, 2, 6, 0, 2),
(15, 3, 8, 0, 3),
(16, 4, 234, 0, 1),
(17, 5, 654, 0, 2),
(18, 6, 186, 0, 3),
(19, 7, 908, 0, 1),
(20, 8, 34, 0, 2);

SELECT setval(pg_get_serial_sequence('items', 'id'), max(id)) FROM items;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM items;
-- +goose StatementEnd