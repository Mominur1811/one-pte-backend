
-- +migrate Up
CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    type VARCHAR(10) NOT NULL, -- 'SST', 'RO', 'RMMCQ'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS questions;