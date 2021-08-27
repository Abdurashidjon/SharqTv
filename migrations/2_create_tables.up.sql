CREATE TABLE IF NOT EXISTS researcher (
    id uuid primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    profession_title varchar(255) not null,
    role_id uuid not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at bigint default 0
);

