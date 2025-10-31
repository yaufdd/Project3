CREATE TABLE refresh_tokens (
    Id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    jti UUID NOT NULL UNIQUE,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN NOT NULL,
    user_agent TEXT,
    ip TEXT,
    device_name TEXT
);