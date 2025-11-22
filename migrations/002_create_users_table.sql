-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    user_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username  TEXT    NOT NULL,
    team_uuid UUID    NOT NULL REFERENCES teams (team_uuid) ON DELETE RESTRICT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX idx_users_team_active ON users (team_uuid, is_active);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_users_team_active;