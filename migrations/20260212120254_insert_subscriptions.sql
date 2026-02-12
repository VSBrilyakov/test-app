-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO subscription (service_name, price, user_id, start_date, end_date) VALUES
    ('Web Hosting', 1200, gen_random_uuid(), '2025-05-01', '2025-12-01'),
    ('Domain Name', 500, gen_random_uuid(), '2024-02-01', '2025-02-01'),
    ('SSL Certificate', 3000, gen_random_uuid(), '2023-08-01', '2025-03-01'),
    ('24/7 Technical Support', 15000, gen_random_uuid(), '2024-01-01', NULL),
    ('Backup Service', 8000, gen_random_uuid(), '2022-12-01', NULL),
    ('Website Development', 50000, gen_random_uuid(), '2024-10-01', '2025-06-01'),
    ('SEO Optimization', 25000, gen_random_uuid(), '2026-01-01', '2026-07-01'),
    ('Contextual Advertising', 35000,gen_random_uuid(), '2024-02-01', '2024-08-01'),
    ('SMM Management', 18000, gen_random_uuid(), '2024-06-01', NULL),
    ('Security Audit', 42000, gen_random_uuid(), '2024-03-01', '2024-04-01');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
TRUNCATE TABLE subscription;