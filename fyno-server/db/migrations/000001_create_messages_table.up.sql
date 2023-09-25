CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY NOT NULL,
    sender uuid NOT NULL,
    receiver uuid NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);