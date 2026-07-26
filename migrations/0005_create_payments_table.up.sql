CREATE TABLE IF NOT EXISTS payment(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    subscription_id INT NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    status TEXT NOT NULL,
    provider TEXT NOT NULL,
    transaction_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
)
