
-- +migrate Up
CREATE TABLE IF NOT EXISTS sst_questions (
     id INTEGER PRIMARY KEY,
     title TEXT NOT NULL,
     answer_time_limit INTEGER NOT NULL,
     audio_files JSONB NOT NULL
);
