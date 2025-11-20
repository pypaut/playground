-- +goose Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tags (
    id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    label varchar(80) not null UNIQUE,
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

-- +goose Down

DROP TABLE IF EXISTS incomes;
DROP TABLE IF EXISTS expenses;
DROP TABLE IF EXISTS budgets;
DROP TABLE IF EXISTS tags;
