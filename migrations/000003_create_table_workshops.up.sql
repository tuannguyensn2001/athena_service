create table workshops
(
    id                    serial primary key,
    name                  varchar,
    thumbnail             varchar,
    private_code          varchar,
    code                  varchar,
    approve_student       boolean default false,
    prevent_student_leave boolean default false,
    approve_show_score    boolean default false,
    disable_newsfeed      boolean default false,
    limit_policy_teacher  boolean default false,
    is_show               boolean default true,
    subject               varchar,
    grade                 varchar,
    created_at            integer,
    updated_at            integer,
    deleted_at            timestamp

)