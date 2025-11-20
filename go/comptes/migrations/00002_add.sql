-- +goose Up

INSERT INTO tags (id, label, description, icon) VALUES
    ('226cb277-5208-4a0d-8b9f-37f3630e288f', 'Factures', 'Paiements récurrents, charges fixes, abonnements', '🧾'),
    ('f9383bb3-6aaf-41d7-906c-d1c580f23d49', 'Épargnes', 'On met de côté', '💰'),
    ('74b344cb-7a16-4af8-8b82-17f477a4f30e', 'Dépenses courantes', 'Dépenses usuelles', '💳'),
    ('a4f7f30c-ae34-4480-8e28-a9ab1741dfb3', 'Dépenses variables', 'Dépenses variables', '💶');

INSERT INTO budgets (id, label, amount, date, tag_id) VALUES
    ('a853f96f-e238-49ee-97f3-1e17f0336df9', 'Courses', 45000, '2025-07-01', '74b344cb-7a16-4af8-8b82-17f477a4f30e'),
    ('d253c593-440d-4bac-ac67-e4ff69355339', 'Épargne chats', 4500, '2025-07-01', 'f9383bb3-6aaf-41d7-906c-d1c580f23d49'),
    ('d3d63ae4-8680-40c6-9f00-af694d83ac6d', 'Cadeau pour jsp qui', 3900, '2025-07-01', 'a4f7f30c-ae34-4480-8e28-a9ab1741dfb3'),
    ('a575ca9f-ddf1-4a52-a718-c018b5169757', 'Loyer', 120000, '2025-07-01', '226cb277-5208-4a0d-8b9f-37f3630e288f');

INSERT INTO expenses (id, label, amount, date, budget_id) VALUES
    ('5a46c201-e9f5-4b0b-b336-0e64e5f96ac9', 'Loyer', 120000, '2025-07-02', 'a575ca9f-ddf1-4a52-a718-c018b5169757'),
    ('74bddac0-b71b-4d7c-89c4-9eef8b7e2ad3', 'Leclerc', 4781, '2025-07-08', 'a853f96f-e238-49ee-97f3-1e17f0336df9');

INSERT INTO incomes (id, label, amount, date) VALUES
    ('961a1dd1-ca6f-412a-83f0-9af6dcd85081', 'Salaire 1', 200042, '2025-07-01'),
    ('247a13b0-32bc-4dd2-8250-b9beabfc939f', 'Salaire 2', 210081, '2025-07-01');

-- +goose Down
