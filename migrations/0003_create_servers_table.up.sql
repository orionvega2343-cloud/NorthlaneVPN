CREATE TABLE IF NOT EXISTS servers(
    id SERIAL PRIMARY KEY,
    host TEXT NOT NULL,
    port INT NOT NULL,
    protocol TEXT NOT NULL,
    status TEXT NOT NULL,
    load_score INT NOT NULL,
    region TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)