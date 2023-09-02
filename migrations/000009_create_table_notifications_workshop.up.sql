create table notification_workshop (
    id serial primary key ,
    content varchar,
    workshop_id integer,
    user_id integer,
    created_at integer,
    updated_at integer
)