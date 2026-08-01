CREATE TABLE IF NOT EXISTS referrals(
    id SERIAL PRIMARY KEY,
    referrer_id INT NOT NULL,
    referred_id INT UNIQUE NOT NULL,
    bonus TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)