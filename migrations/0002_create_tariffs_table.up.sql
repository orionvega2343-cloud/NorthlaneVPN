CREATE TABLE IF NOT EXISTS tariffs(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    duration_days INT NOT NULL,
    taraffic_limit_gb INT NOT NULL,
    device_limit INT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    is_active BOOLEAN NOT NULL
)