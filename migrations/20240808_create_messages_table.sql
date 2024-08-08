-- +goose Up
create table messages (
    id serial primary key,
    chat_id int references chats(id) on delete cascade,
    user_id int not null,
    text text,
    created_at timestamp not null default now(),
    primary key (id, chat_id)
);

-- +goose Down
drop table messages;