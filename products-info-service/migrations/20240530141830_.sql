-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  description TEXT,
  photo VARCHAR(255),
  price NUMERIC(10, 2),
  quantity INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
