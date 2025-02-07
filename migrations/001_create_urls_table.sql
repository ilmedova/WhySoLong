CREATE TABLE urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original TEXT NOT NULL,
    shortened TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
