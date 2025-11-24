-- +goose Up
CREATE TABLE IF NOT EXISTS pull_request_status
(
    id   INT PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO pull_request_status (id, name)
VALUES (1, 'OPEN'),
       (2, 'MERGED ');

-- +goose Down
DROP TABLE IF EXISTS pull_request_status;
