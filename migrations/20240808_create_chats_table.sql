-- +goose Up
create table chats (
    id serial primary key,
    name varchar(50) not null UNIQUE,
    created_at timestamp not null default now(),
);

-- +goose Down
drop table chats;