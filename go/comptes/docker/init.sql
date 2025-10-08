CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tags (
    id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    label varchar(80) not null,
    description varchar(80),
    icon varchar(2)
);

CREATE TABLE budgets (
    id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    label varchar(80) not null,
    amount integer default 0.0,
    date timestamp,
    tag_id uuid references tags(id)
);

CREATE TABLE expenses (
    id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    label varchar(80) not null,
    amount integer default 0.0,
    date timestamp,
    budget_id uuid references budgets(id)
);

CREATE TABLE incomes (
    id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    label varchar(80) not null,
    amount integer default 0.0,
    date timestamp
);

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

INSERT INTO expenses (label, amount, date, budget_id) VALUES
    ('Loyer', 120000, '2025-07-02', 'a575ca9f-ddf1-4a52-a718-c018b5169757'),
    ('Leclerc', 4781, '2025-07-08', 'a853f96f-e238-49ee-97f3-1e17f0336df9');

INSERT INTO incomes (label, amount, date) VALUES
    ('Salaire 1', 200042, '2025-07-01'),
    ('Salaire 2', 210081, '2025-07-01');
