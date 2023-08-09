create table profiles
(
    id         serial primary key,
    user_id    integer,
    username   varchar,
    school     varchar,
    birthday   integer,
    avatar_url varchar,
    created_at integer,
    updated_at integer
)