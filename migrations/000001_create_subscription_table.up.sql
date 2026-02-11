CREATE TABLE IF NOT EXISTS subscription
(
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    price INTEGER CHECK (price > 0),
    user_id UUID,
    start_date DATE,
    end_date DATE CHECK (end_date IS NULL OR end_date >= start_date)
);

CREATE INDEX IF NOT EXISTS idx_service_id ON subscription(service_name);
CREATE INDEX IF NOT EXISTS idx_user_id ON subscription(user_id);