-- +goose Up
CREATE TABLE IF NOT EXISTS teams
(
    team_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_name TEXT NOT NULL UNIQUE
);

CREATE UNIQUE INDEX idx_teams_team_name ON teams (team_name);

-- +goose Down
DROP TABLE IF EXISTS teams;

DROP INDEX IF EXISTS idx_teams_team_name;