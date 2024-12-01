
-- +migrate Up
CREATE TABLE user_history (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  question_id INT NOT NULL,
  question_type VARCHAR(255) NOT NULL,
  answer JSONB,
  obtain_marks FLOAT8 NOT NULL,
  total_marks FLOAT8 NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
