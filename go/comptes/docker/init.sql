CREATE TABLE tags (
    label varchar(80) not null unique primary key,
    description varchar(80),
    icon varchar(2)
);

CREATE TABLE budgets (
    label varchar(80) not null unique primary key,
    amount real default 0.0,
    date timestamp,
    tag varchar(80) references tags(label)
);

CREATE TABLE expenses (
    label varchar(80) not null unique primary key,
    amount real default 0.0,
    date timestamp,
    budget varchar(80) references budgets(label)
);

CREATE TABLE incomes (
    label varchar(80) not null unique primary key,
    amount real default 0.0,
    date timestamp
);

INSERT INTO tags (label, description, icon) VALUES
    ('Factures', 'Paiements récurrents, charges fixes, abonnements', '🧾'),
    ('Épargnes', 'On met de côté', '💰'),
    ('Dépenses courantes', 'Dépenses usuelles', '💳'),
    ('Dépenses variables', 'Dépenses variables', '💶');

INSERT INTO budgets (label, amount, date, tag) VALUES
    ('Courses', 450.0, '2025-07-01', 'Dépenses courantes'),
    ('Épargne chats', 45.0, '2025-07-01', 'Épargnes'),
    ('Cadeau pour jsp qui', 39., '2025-07-01', 'Dépenses variables'),
    ('Loyer', 1200.0, '2025-07-01', 'Factures');

INSERT INTO expenses (label, amount, date, budget) VALUES
    ('Loyer', 1200.0, '2025-07-02', 'Loyer'),
    ('Leclerc', 47.81, '2025-07-08', 'Courses');

INSERT INTO incomes (label, amount, date) VALUES
    ('Salaire 1', 2000.42, '2025-07-01'),
    ('Salaire 2', 2100.81, '2025-07-01');
