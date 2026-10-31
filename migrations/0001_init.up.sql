CREATE TABLE IF NOT EXISTS threats (
    md5         CHAR(32) PRIMARY KEY,
    name        TEXT NOT NULL,
    severity    TEXT NOT NULL,
    type        TEXT NOT NULL,
    description TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);