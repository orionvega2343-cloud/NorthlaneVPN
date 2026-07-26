CREATE TABLE IF NOT EXISTS subscriptions(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    tariff_id INT NOT NULL REFERENCES tariff(id),
    server_id INT NOT NULL REFERENCES servers(id),
    status TEXT NOT NULL,
    starts_at TIMESTAMP NOT NULL DEFAULT NOW(),
    finishes_at TIMESTAMP  NOT NULL DEFAULT NOW(),
    traffic_used_gb INT NOT NULL,
    is_trial BOOLEAN NOT NULL
)