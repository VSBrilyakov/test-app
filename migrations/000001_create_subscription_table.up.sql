CREATE TABLE IF NOT EXISTS subscription
(
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    price INTEGER,
    user_id UUID UNIQUE,
    start_date DATE,
    end_date DATE
);

CREATE INDEX IF NOT EXISTS idx_user_id ON subscription(user_id);