CREATE TABLE refresh_tokens (
    id           SERIAL PRIMARY KEY,
    user_id      VARCHAR(100) NOT NULL,
    expires_at   TIMESTAMP NOT NULL,
    revoked      BOOLEAN DEFAULT FALSE,
    revoked_at   TIMESTAMPTZ NULL,
    created_at   TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_refresh_tokens_user ON refresh_tokens(user_id);
