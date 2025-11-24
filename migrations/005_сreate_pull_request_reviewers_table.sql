-- +goose Up
CREATE TABLE IF NOT EXISTS pull_request_reviewers (
    pull_request_id TEXT NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    reviewer_id     TEXT NOT NULL REFERENCES users (user_id),
    PRIMARY KEY (pull_request_id, reviewer_id)
);

CREATE INDEX idx_pr_reviewers_reviewer_pr
    ON pull_request_reviewers (reviewer_id, pull_request_id);

CREATE INDEX idx_pr_reviewers_reviewer_id ON pull_request_reviewers (reviewer_id);

-- +goose Down
DROP TABLE IF EXISTS pull_request_reviewers;

DROP INDEX IF EXISTS idx_pr_reviewers_reviewer_id;