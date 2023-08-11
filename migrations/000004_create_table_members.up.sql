create table members
(
    id          serial primary key,
    user_id     integer,
    workshop_id integer,
    role        varchar,
    created_at  integer,
    updated_at  integer
)