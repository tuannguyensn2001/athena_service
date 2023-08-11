create table posts
(
    id          serial primary key,
    content     varchar,
    user_id     integer,
    workshop_id integer,
    pinned_at   integer,
    created_at  integer,
    updated_at  integer,
    deleted_at  timestamp
)