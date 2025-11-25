-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    user_id   TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
    username  TEXT    NOT NULL,
    team_id   uuid    NOT NULL REFERENCES teams (team_id) ON DELETE RESTRICT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_users_team_id ON users (team_id);

CREATE INDEX idx_users_active ON users(is_active);

CREATE INDEX idx_users_team_active ON users (team_id, is_active);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_users_team_id;

DROP INDEX IF EXISTS idx_users_team_active;

DROP INDEX IF EXISTS idx_users_team_active;