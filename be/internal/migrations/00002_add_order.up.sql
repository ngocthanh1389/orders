CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    time TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_name ON orders("name");
