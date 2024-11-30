
-- +migrate Up
CREATE TABLE IF NOT EXISTS ro_questions (
  id INTEGER PRIMARY KEY,
  paragraph JSONB NOT NULL,
  correct_order INT[] NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS ro_questions;