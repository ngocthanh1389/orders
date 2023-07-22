CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    pwd TEXT NOT NULL,
    CONSTRAINT unique_name UNIQUE (name)
);
