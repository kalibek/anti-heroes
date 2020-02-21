create table users (
    id bigserial,
    email varchar not null unique,
    name varchar not null,
    primary key (id)
);