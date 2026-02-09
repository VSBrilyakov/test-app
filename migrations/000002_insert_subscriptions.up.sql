INSERT INTO subscription (service_name, price, user_id, start_date, end_date) VALUES
    ('Web Hosting', 1200, gen_random_uuid(), '2025-05-01', '2025-12-31'),
    ('Domain Name', 500, gen_random_uuid(), '2024-02-01', '2025-02-01'),
    ('SSL Certificate', 3000, gen_random_uuid(), '2023-08-01', '2025-03-01'),
    ('24/7 Technical Support', 15000, gen_random_uuid(), '2024-01-15', NULL),
    ('Backup Service', 8000, gen_random_uuid(), '2022-12-10', NULL),
    ('Website Development', 50000, gen_random_uuid(), '2024-10-05', '2025-06-05'),
    ('SEO Optimization', 25000, gen_random_uuid(), '2026-01-20', '2026-07-20'),
    ('Contextual Advertising', 35000,gen_random_uuid(), '2024-02-15', '2024-08-15'),
    ('SMM Management', 18000, gen_random_uuid(), '2024-06-15', NULL),
    ('Security Audit', 42000, gen_random_uuid(), '2024-03-10', '2024-04-10');