CREATE TABLE tags (
    id serial primary key,
    label varchar(80) not null,
    description varchar(80),
    icon varchar(2)
);

CREATE TABLE budgets (
    id serial primary key,
    label varchar(80) not null,
    amount decimal default 0.0,
    date timestamp,
    tag_id integer references tags(id)
);

CREATE TABLE expenses (
    id serial primary key,
    label varchar(80) not null,
    amount decimal default 0.0,
    date timestamp,
    budget_id integer references budgets(id)
);

CREATE TABLE incomes (
    id serial primary key,
    label varchar(80) not null,
    amount decimal default 0.0,
    date timestamp
);

INSERT INTO tags (label, description, icon) VALUES
    ('Factures', 'Paiements récurrents, charges fixes, abonnements', '🧾'),
    ('Épargnes', 'On met de côté', '💰'),
    ('Dépenses courantes', 'Dépenses usuelles', '💳'),
    ('Dépenses variables', 'Dépenses variables', '💶');

INSERT INTO budgets (label, amount, date, tag_id) VALUES
    ('Courses', 450.0, '2025-07-01', 3),
    ('Épargne chats', 45.0, '2025-07-01', 2),
    ('Cadeau pour jsp qui', 39., '2025-07-01', 4),
    ('Loyer', 1200.0, '2025-07-01', 1);

INSERT INTO expenses (label, amount, date, budget_id) VALUES
    ('Loyer', 1200.0, '2025-07-02', 4),
    ('Leclerc', 47.81, '2025-07-08', 1);

INSERT INTO incomes (label, amount, date) VALUES
    ('Salaire 1', 2000.42, '2025-07-01'),
    ('Salaire 2', 2100.81, '2025-07-01');
