-- +goose Up
CREATE TABLE IF NOT EXISTS teams (
    team_uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS teams;