create table users
(
    id                serial primary key,
    phone             varchar,
    password          varchar,
    email             varchar,
    email_verified_at integer,
    role              varchar,
    created_at        integer,
    updated_at        integer,
    deleted_at        timestamp
)