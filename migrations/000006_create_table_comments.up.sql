create table comments
(
    id         serial primary key,
    user_id    int  not null,
    post_id    int  not null,
    content    text not null,
    created_at integer,
    updated_at integer,
    deleted_at timestamp
)