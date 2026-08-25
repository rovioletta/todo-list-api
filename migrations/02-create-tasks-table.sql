CREATE TYPE task_status AS ENUM (
    'backlog',
    'todo',
    'in_progress',
    'in_review',
    'done',
    'canceled'
);

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  title VARCHAR(32) NOT NULL,
  description TEXT NOT NULL,
  status task_status NOT NULL DEFAULT 'todo',
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);