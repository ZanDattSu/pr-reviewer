-- +goose Up
CREATE TABLE IF NOT EXISTS pull_request_status (
    id   INT PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS pull_request_status;
