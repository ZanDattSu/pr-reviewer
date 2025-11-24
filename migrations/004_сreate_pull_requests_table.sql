-- +goose Up
CREATE TABLE IF NOT EXISTS pull_requests
(
    pull_request_id   TEXT PRIMARY KEY DEFAULT gen_random_uuid(),
    pull_request_name TEXT NOT NULL,
    author_id         TEXT NOT NULL REFERENCES users (user_id),
    status_id         INT  NOT NULL references pull_request_status (id),
    created_at        TIMESTAMPTZ      DEFAULT NOW(),
    merged_at         TIMESTAMPTZ
);

CREATE INDEX idx_pull_requests_author_id ON pull_requests (author_id);

CREATE INDEX idx_pull_requests_status ON pull_requests (status_id);

-- +goose Down
DROP TABLE IF EXISTS pull_requests;

DROP INDEX IF EXISTS idx_pull_requests_author_id;

DROP INDEX IF EXISTS idx_pull_requests_status;