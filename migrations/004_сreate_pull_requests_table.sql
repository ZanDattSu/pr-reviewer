-- +goose Up
CREATE TABLE IF NOT EXISTS pull_requests
(
    pull_request_id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pull_request_name TEXT NOT NULL,
    author_id         UUID NOT NULL REFERENCES users (user_id),
    status            INT  NOT NULL references pull_request_status (id),
    created_at        TIMESTAMPTZ      DEFAULT NOW(),
    merged_at         TIMESTAMPTZ
);

-- на будущее: статистика по авторам, выборки их PR
CREATE INDEX idx_pull_requests_author_id ON pull_requests (author_id);

-- нужен для выборок только OPEN PR
CREATE INDEX idx_pull_requests_status ON pull_requests (status);

-- +goose Down
DROP TABLE IF EXISTS pull_requests;

DROP INDEX IF EXISTS idx_pull_requests_author_id;

DROP INDEX IF EXISTS idx_pull_requests_status;