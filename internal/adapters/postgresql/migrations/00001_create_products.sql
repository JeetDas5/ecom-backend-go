-- +goose Up
-- SELECT 'up SQL query';
CREATE TABLE
    IF NOT EXISTS products (
        id BIGSERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT,
        price_in_cents INTEGER NOT NULL CHECK (price_in_cents >= 0) ,
        created_at TIMESTAMPTZ DEFAULT NOW (),
        updated_at TIMESTAMPTZ DEFAULT NOW ()
    );

-- +goose Down
SELECT
    'down SQL query';