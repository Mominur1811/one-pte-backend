
-- +migrate Up
CREATE TABLE IF NOT EXISTS mcq_questions (
   id INTEGER PRIMARY KEY,
   options JSONB NOT NULL,
   correct_option INT NOT NULL
);

