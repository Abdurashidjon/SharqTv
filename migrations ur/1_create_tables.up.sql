CREATE TABLE IF NOT EXISTS company (
    id uuid primary key,
    name varchar(255) not null,
    inn varchar(255),
    owner_name varchar(255),
    email varchar(255),
    phone varchar(30),
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at bigint default 0
);

CREATE UNIQUE INDEX IF NOT EXISTS company_u1 ON company(inn);

CREATE TABLE IF NOT EXISTS respondent (
    id uuid primary key,
    name varchar(255) not null,
    email varchar(255),
    phone varchar(30),
    sber_id varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at bigint default 0
);